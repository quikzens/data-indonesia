package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *Handler) chi.Router {
	r := chi.NewRouter()

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
