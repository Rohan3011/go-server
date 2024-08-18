package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type LocalStorage struct {
	Directory string
}

func NewLocalStorage(directory string) *LocalStorage {
	return &LocalStorage{Directory: directory}
}

func (s *LocalStorage) UploadFile(file multipart.File, filename string) (string, error) {
	path := filepath.Join(s.Directory, filename)
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// Return the file path instead of the file itself
func (s *LocalStorage) GetFile(filename string) (string, error) {
	path := filepath.Join(s.Directory, filename)
	return path, nil
}

func (s *LocalStorage) DeleteFile(filename string) error {
	path := filepath.Join(s.Directory, filename)
	return os.Remove(path)
}

func (s *LocalStorage) ListFiles() ([]string, error) {
	var files []string
	entries, err := os.ReadDir(s.Directory)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}
