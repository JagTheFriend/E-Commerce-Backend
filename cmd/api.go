package main

import (
	"log"
	"net/http"
	"time"

	"github.com/JagTheFriend/ecommerce/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dbUrl string
}

type application struct {
	config config
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID) // used for ratelimiting
	r.Use(middleware.RealIP)    // used for ratelimiting and analytics
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Time out requests after 30 seconds, preventing further processing of the request
	r.Use(middleware.Timeout(30 * time.Second))

	productsService := products.NewService()
	productsHandler := products.NewHandler(productsService)
	r.Route("/products", func(r chi.Router) {
		r.Get("/", productsHandler.ListProducts)
	})

	return r
}

func (a *application) Start(h http.Handler) error {
	srv := &http.Server{
		Addr:         a.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Minute,
	}
	log.Printf("Server has started on %v\n", a.config.addr)
	return srv.ListenAndServe()

}
