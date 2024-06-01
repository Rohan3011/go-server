package view

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/rohan3011/go-server/templates"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type GlobalState struct {
	Count int
}

var global GlobalState

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	component := templates.IndexPage(global.Count, 0)
	router.Handle("/", templ.Handler(component))
}
