package upload

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/rohan3011/go-server/services/upload/storage"
	"github.com/rohan3011/go-server/utils"
)

type FileDataKey string

const fileDataKey = FileDataKey("fileData")

// FileUploadMiddleware handles file upload and stores the file URL or path in the context
func FileUploadMiddleware(storage storage.Storage) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
				file, header, err := r.FormFile("file")
				if err != nil {
					utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error retrieving file"))
					return
				}
				defer file.Close()

				// Generate file name
				filename := utils.GenerateFilename(header.Filename)

				// Upload file and get its URL or file path
				filePath, err := storage.UploadFile(file, filename)
				if err != nil {
					utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error uploading file: %s", err))

					return
				}

				// Attach file URL or path to request context
				ctx := context.WithValue(r.Context(), fileDataKey, &FileInsert{
					Filename: filename,
					FileURL:  filePath,
					FileSize: header.Size,
				})
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

// retrieves the file data from the request context
func GetFileDataFromContext(ctx context.Context) (*FileInsert, bool) {
	fileData, ok := ctx.Value(fileDataKey).(*FileInsert)
	return fileData, ok
}
