package utils

import (
	"fmt"
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
