package upload

import (
	"mime/multipart"
)

type Storage interface {
	UploadFile(file multipart.File, filename string) (string, error)
	GetFile(filename string) (string, error)
	DeleteFile(filename string) error
	ListFiles() ([]string, error)
}
