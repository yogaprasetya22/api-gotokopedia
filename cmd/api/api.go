package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/yogaprasetya22/api-gotokopedia/docs"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr   string
	db     dbConfig
	apiURL string
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

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

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
			r.Post("/", app.createProductHandler)
			r.Get("/", app.getProductCategoryFeed)

			r.Route("/{productID}", func(r chi.Router) {
				r.Use(app.productContextMiddleware)

				r.Get("/", app.getProductHandler)
				r.Patch("/", app.updateProductHandler)
				r.Delete("/", app.deleteProductHandler)
			})

			r.Group(func(r chi.Router) {
				r.Get("/feed", app.getProductFeedHandler)
			})
		})

		// user
		r.Route("/users", func(r chi.Router) {

			r.Route("/{userID}", func(r chi.Router) {

				r.Get("/", app.getUserHandler)
				r.Put("/follow", app.followUserHandler)
				r.Put("/unfollow", app.unfollowUserHandler)
			})
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
