package utils

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

// GenerateFilename creates a unique filename using the base name and a timestamp
func GenerateFilename(originalFilename string) string {
	ext := filepath.Ext(originalFilename)                     // Get the file extension
	base := originalFilename[:len(originalFilename)-len(ext)] // Get the base name without extension
	timestamp := time.Now().Format("20060102150405")          // Current timestamp in the format YYYYMMDDHHMMSS
	filename := fmt.Sprintf("%s-%s%s", base, timestamp, ext)  // Combine them into a unique filename
	return filename
}

func GetFileInfo(file multipart.File, fileHeader *multipart.FileHeader) (int64, string, error) {
	// Get the file size
	fileSize := fileHeader.Size

	// Read the first 512 bytes of the file to detect the content type
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		return 0, "", err
	}
	// Reset the file pointer to the beginning for further use
	_, err = file.Seek(0, 0)
	if err != nil {
		return 0, "", err
	}

	// Detect the MIME type
	mimeType := http.DetectContentType(buf)

	return fileSize, mimeType, nil
}
