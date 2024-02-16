package handlers

import (
	"github.com/allrivenjs/goapi/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddleware.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		/// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
