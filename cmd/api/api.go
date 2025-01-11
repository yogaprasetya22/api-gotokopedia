package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yogaprasetya22/api-gotokopedia/docs"
	"github.com/yogaprasetya22/api-gotokopedia/internal/auth"
	"github.com/yogaprasetya22/api-gotokopedia/internal/env"
	"github.com/yogaprasetya22/api-gotokopedia/internal/mailer"
	"github.com/yogaprasetya22/api-gotokopedia/internal/ratelimiter"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type application struct {
	config            config
	store             store.Storage
	cacheStorage      cache.Storage
	logger            *zap.SugaredLogger
	mailer            mailer.Client
	authenticator     auth.Authenticator
	rateLimiter       ratelimiter.Limiter
	googleOauthConfig *oauth2.Config
	session           sessions.Store
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	mail        mailConfig
	frontendURL string
	auth        authConfig
	redisCfg    redisConfig
	rateLimiter ratelimiter.Config
	google      googleConfig
}

type googleConfig struct {
	clientID     string
	clientSecret string
	redirectURL  string
	scopes       []string
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type basicConfig struct {
	usrname string
	pass    string
}

type mailConfig struct {
	host      string
	port      int
	username  string
	password  string
	timeout   time.Duration
	sender    string
	fromEmail string
	exp       time.Duration
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(app.RateLimiterMiddleware)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{env.GetString("CORS_ALLOWED_ORIGIN", "http://localhost:3000")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/auth/google", app.googleLoginHandler)
	r.Get("/auth/google/callback", app.googleCallbackHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		// r.With(app.BasicAuthMiddleware()).Get("/debug/vars", expvar.Handler().ServeHTTP)

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		/// category
		r.Route("/category", func(r chi.Router) {
			r.Get("/", app.getAllCategoryHandler)
			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Post("/", app.createCategoryHandler)
			})
		})

		/// toko
		r.Route("/toko", func(r chi.Router) {
			r.Post("/", app.createTokoHandler)
			r.Route("/{slug_toko}", func(r chi.Router) {
				r.Get("/", app.getProductTokoHandler)
			})
		})

		/// comment feed
		r.Route("/comment", func(r chi.Router) {
			r.Route("/{slug}", func(r chi.Router) {
				r.Get("/", app.getCommentsHandler)
			})
		})

		/// product
		r.Route("/product", func(r chi.Router) {
			r.Post("/", app.createProductHandler)

			r.Route("/{productID}", func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Use(app.productContextMiddleware)

				r.Patch("/", app.checkProductOwnership(app.updateProductHandler))
				r.Delete("/", app.checkProductOwnership(app.deleteProductHandler))

				/// comment
				r.Route("/comment", func(r chi.Router) {
					r.Post("/", app.createCommentHandler)
					r.Get("/", app.getCommentsHandler)

					r.Route("/{commentID}", func(r chi.Router) {
						r.Use(app.commentContextMiddleware)

						r.Patch("/", app.checkCommentOwnership(app.updateCommentHandler))
						r.Delete("/", app.checkCommentOwnership(app.deleteCommentHandler))
					})
				})
			})
		})

		r.Route("/catalogue", func(r chi.Router) {
			r.Get("/", app.getProductCategoryFeed)

			r.Route("/{slug_toko}", func(r chi.Router) {
				r.Get("/", app.getTokoHandler)

				r.Route("/{slug_product}", func(r chi.Router) {
					r.Get("/", app.getProductHandler)
				})
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getProductFeedHandler)
			})
		})

		/// user
		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Get("/current", app.getCurrentUserHandler)
			})

			r.Route("/{userID}", func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)

				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})
		})

		/// Auth routes
		r.Route("/authentication", func(r chi.Router) {
			r.Post("/user", app.registerUserHandler)
			r.Get("/token", app.createTokenHandler)
			r.Get("/logout", app.logoutHandler)
		})

	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	///docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/api/v1"

	/// Membuat server HTTP baru
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
