package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yogaprasetya22/api-gotokopedia/docs"
	"github.com/yogaprasetya22/api-gotokopedia/internal/auth"
	"github.com/yogaprasetya22/api-gotokopedia/internal/mailer"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store/cache"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	cacheStorage  cache.Storage
	logger        *zap.SugaredLogger
	mailer        mailer.Client
	authenticator auth.Authenticator
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

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/auth/google", app.googleLoginHandler)
	r.Get("/auth/google/callback", app.googleCallbackHandler)

	r.Route("/api/v1", func(r chi.Router) {
		r.With(app.BasicAuthMiddleware()).Get("/health", app.healthCheckHandler)

		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		// category
		r.Route("/category", func(r chi.Router) {
			r.Post("/", app.createCategoryHandler)
		})

		// toko
		r.Route("/toko", func(r chi.Router) {
			r.Post("/", app.createTokoHandler)
		})

		// product
		r.Route("/catalogue", func(r chi.Router) {
			r.Use(app.AuthTokenMiddleware)
			r.Post("/", app.createProductHandler)
			r.Get("/", app.getProductCategoryFeed)

			r.Route("/{productID}", func(r chi.Router) {
				r.Use(app.productContextMiddleware)

				r.Get("/", app.getProductHandler)
				r.Patch("/", app.updateProductHandler)
				r.Delete("/", app.deleteProductHandler)

				// comment
				r.Route("/comment", func(r chi.Router) {
					r.Post("/", app.createCommentHandler)

					r.Route("/{commentID}", func(r chi.Router) {
						r.Use(app.commentContextMiddleware)

						r.Patch("/", app.updateCommentHandler)
						r.Delete("/", app.deleteCommentHandler)
					})
				})
			})

			r.Group(func(r chi.Router) {
				r.Use(app.AuthTokenMiddleware)
				r.Get("/feed", app.getProductFeedHandler)
			})
		})

		// user
		r.Route("/users", func(r chi.Router) {
			r.Put("/activate/{token}", app.activateUserHandler)

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
			r.Post("/token", app.createTokenHandler)
		})

	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/api/v1"

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("Server running on %s", app.config.addr)

	return srv.ListenAndServe()
}
