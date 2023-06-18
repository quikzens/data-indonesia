package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(h *Handler) chi.Router {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // use this to allow specific origin hosts
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
	}))

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	r.Route("/provinces", func(r chi.Router) {
		r.Get("/", h.GetProvinces)
		r.Get("/paginate", h.PaginateProvinces)
	})

	r.Route("/cities", func(r chi.Router) {
		r.Get("/", h.GetCities)
		r.Get("/paginate", h.PaginateCities)
	})

	r.Route("/subdistricts", func(r chi.Router) {
		r.Get("/", h.GetSubdistricts)
		r.Get("/paginate", h.PaginateSubdistricts)
	})

	r.Route("/villages", func(r chi.Router) {
		r.Get("/", h.GetVillages)
		r.Get("/paginate", h.PaginateVillages)
	})

	return r
}
