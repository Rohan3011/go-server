package upload

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rohan3011/go-server/services/auth"
	"github.com/rohan3011/go-server/types"
	"github.com/rohan3011/go-server/utils"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Route("/uploads", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Use(FileUploadMiddleware(h.storage)) // Use the file upload middleware
		r.Post("/", h.UploadFile)
		r.Get("/", h.ListFiles)
		r.Get("/{filename}", h.GetFile)
		r.Delete("/{filename}", h.DeleteFile)
	})
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the file URL or path from the context
	fileURL, ok := GetFileURLFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("file upload failed"))
		return
	}

	utils.WriteJSON(w, http.StatusCreated, &types.Response{
		Status: "success",
		Data:   map[string]string{"fileURL": fileURL},
	})
}

func (h *Handler) ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := h.storage.ListFiles()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
		Data:   files,
	})
}

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	filePath, err := h.storage.GetFile(filename)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	http.ServeFile(w, r, filePath)
}

func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")
	err := h.storage.DeleteFile(filename)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
	})
}
