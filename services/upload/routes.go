package upload

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rohan3011/go-server/services/auth"
	"github.com/rohan3011/go-server/services/upload/storage"
	"github.com/rohan3011/go-server/types"
	"github.com/rohan3011/go-server/utils"
)

type Handler struct {
	store       UploadStore
	fileStorage storage.Storage
}

func NewHandler(store UploadStore, fileStorage storage.Storage) *Handler {
	return &Handler{store, fileStorage}
}

func (h *Handler) RegisterRoutes(router *chi.Mux) {
	router.Route("/uploads", func(r chi.Router) {
		r.Use(auth.AuthMiddleware)
		r.Use(FileUploadMiddleware(h.fileStorage)) // Use the file upload middleware
		r.Post("/", h.UploadFile)
		r.Get("/", h.ListFiles)
		r.Get("/{id}", h.GetFile)
		r.Delete("/{id}", h.DeleteFile)
	})
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {

	user, ok := auth.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("file upload failed"))
		return
	}

	// Retrieve the file data from the context
	fileData, ok := GetFileDataFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("file upload failed"))
		return
	}

	fileData.UserID = user.ID

	err := h.store.Create(*fileData)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, &types.Response{
		Status: "success",
		Data:   fileData,
	})
}

func (h *Handler) ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := h.store.List()
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
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	file, err := h.store.Read(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
		Data:   file,
	})
}

func (h *Handler) DeleteFile(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	file, err := h.store.Read(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid file id %v", err))
		return
	}

	// Delete record from the table
	err = h.store.Delete(id)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Delete file from storage
	err = h.fileStorage.DeleteFile(file.Filename)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("failed to delete file from storage %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, &types.Response{
		Status: "success",
		Data:   file,
	})
}
