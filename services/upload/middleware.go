package upload

import (
	"context"
	"net/http"
	"strings"

	"github.com/rohan3011/go-server/services/upload/storage"
	"github.com/rohan3011/go-server/utils"
)

type FileURLKey string

const fileURLKey = FileURLKey("fileURL")

// FileUploadMiddleware handles file upload and stores the file URL or path in the context
func FileUploadMiddleware(storage storage.Storage) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
				file, header, err := r.FormFile("file")
				if err != nil {
					http.Error(w, "Error retrieving file", http.StatusBadRequest)
					return
				}
				defer file.Close()

				// Generate file name
				filename := utils.GenerateFilename(header.Filename)

				// Upload file and get its URL or file path
				filePath, err := storage.UploadFile(file, filename)
				if err != nil {
					http.Error(w, "Error uploading file", http.StatusInternalServerError)
					return
				}

				// Attach file URL or path to request context
				ctx := context.WithValue(r.Context(), fileURLKey, filePath)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

// GetFileURLFromContext retrieves the file URL or path from the request context
func GetFileURLFromContext(ctx context.Context) (string, bool) {
	url, ok := ctx.Value(fileURLKey).(string)
	return url, ok
}
