package todo

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rohan3011/go-server/services/auth"
	"github.com/rohan3011/go-server/types"
	"github.com/rohan3011/go-server/utils"
)

type Handler struct {
	store TodoStore
}

func NewHandler(store TodoStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {

	router.Route("/todos", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Get("/", h.ListTodos)
		r.Get("/{id}", h.GetTodo)
		r.Post("/", h.CreateTodo)
		r.Put("/{id}", h.UpdateTodo)
		r.Delete("/{id}", h.DeleteTodo)
	})
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(auth.UserKey).(*auth.Claims)
	if !ok {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	todos, err := h.store.List(user.ID, 10, 0, nil)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
		Data:   todos,
	})
}

func (h *Handler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	todo, err := h.store.Read(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
		Data:   todo,
	})
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	//get JSON payload
	var payload TodoInsert
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err := h.store.Create(payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, &types.Response{
		Status: "success",
	})
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	//get JSON payload
	var payload TodoUpdate
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	err = h.store.Update(id, payload)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, &types.Response{
		Status: "success",
	})
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid ID"))
		return
	}

	err = h.store.Delete(id)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
	})
}
