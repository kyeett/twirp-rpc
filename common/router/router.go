package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewDefault() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	return r
}
