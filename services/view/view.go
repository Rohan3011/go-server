package view

import (
	"fmt"
	"net/http"

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
	router.Route("/view/actions", func(r chi.Router) {
		r.Post("/global-increment", actionGlobalIncrement)
	})
}

func actionGlobalIncrement(w http.ResponseWriter, r *http.Request) {
	// Increment the global counter
	global.Count += 1

	// Write the new count as a response
	response := fmt.Sprintf("New count: %d", global.Count)
	w.Write([]byte(response))
}
