package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yogaprasetya22/api-gotokopedia/internal/store"
	"go.uber.org/zap"
)

type application struct {
	config config
	store  store.Storage
	logger *zap.SugaredLogger
}

type config struct {
	addr string
	db   dbConfig
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

		r.Route("/category", func(r chi.Router) {
			r.Post("/", app.createCategoryHandler)
		})

		r.Route("/toko", func(r chi.Router) {
			r.Post("/", app.createTokoHandler)
		})

		r.Route("/product", func(r chi.Router) {
			r.Post("/", app.createProductHandler)

			r.Route("/{productID}", func(r chi.Router) {
				r.Use(app.productContextMiddleware)
				
				r.Get("/", app.getProductHandler)
				// r.Put("/", app.updateProductHandler)
				// r.Delete("/", app.deleteProductHandler)
			})
		})

	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	app.logger.Infof("Server running on %s", app.config.addr)

	return srv.ListenAndServe()
}
