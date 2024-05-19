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

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	component := templates.HelloWorld("title")
	router.Handle("/", templ.Handler(component))
}
