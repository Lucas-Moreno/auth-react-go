package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"api"
)

func NewHandler(store api.Store) *Handler {
	h := &Handler{
		Mux:	chi.NewMux(),
		store: 	store,
	}
	h.Use(middleware.Logger)
	h.Route("/thread", func(r chi.Router){
		r.Get("/", h.ThreadList())
	})
	h.Route("/thread/{id}", func(r chi.Router){
		r.Get("/", h.ThreadById())
	})
	
	return h
}

type Handler struct {
	*chi.Mux
	store api.Store
}
