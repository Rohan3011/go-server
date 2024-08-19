package upload

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type FileStore struct {
	db *sql.DB
}

type FileMetadata struct {
	ID         int       `json:"id"`
	Filename   string    `json:"filename"`
	UserID     int       `json:"user_id"`
	UploadTime time.Time `json:"upload_time"`
	FileURL    string    `json:"file_url"`
	FileSize   int64     `json:"file_size"`
	MimeType   string    `json:"mime_type"`
}

type FileInsert struct {
	Filename string `json:"filename"`
	UserID   int    `json:"user_id"`
	FileURL  string `json:"file_url"`
	FileSize int64  `json:"file_size"`
	MimeType string `json:"mime_type"`
}

func NewFileStore(db *sql.DB) *FileStore {
	return &FileStore{db: db}
}

func (s *FileStore) Create(metadata FileInsert) error {
	query := `INSERT INTO file_metadata (filename, user_id, upload_time, file_url, file_size, mime_type) VALUES (?, ?, ?, ?, ?, ?);`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, metadata.Filename, metadata.UserID, time.Now(), metadata.FileURL, metadata.FileSize, metadata.MimeType)
	if err != nil {
		return fmt.Errorf("could not create file metadata: %v", err)
	}
	return nil
}

func (s *FileStore) List() ([]FileMetadata, error) {
	query := `SELECT id, filename, user_id, upload_time, file_url, file_size, mime_type FROM file_metadata`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying database: %v", err)
	}
	defer rows.Close()

	var files []FileMetadata
	for rows.Next() {
		var file FileMetadata
		if err := rows.Scan(&file.ID, &file.Filename, &file.UserID, &file.UploadTime, &file.FileURL, &file.FileSize, &file.MimeType); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		files = append(files, file)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return files, nil
}

func (s *FileStore) Read(id int) (FileMetadata, error) {
	query := `SELECT id, filename, user_id, upload_time, file_url, file_size, mime_type FROM file_metadata WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row := s.db.QueryRowContext(ctx, query, id)
	var file FileMetadata
	err := row.Scan(&file.ID, &file.Filename, &file.UserID, &file.UploadTime, &file.FileURL, &file.FileSize, &file.MimeType)
	if err != nil {
		if err == sql.ErrNoRows {
			return FileMetadata{}, fmt.Errorf("file metadata not found")
		}
		return FileMetadata{}, fmt.Errorf("could not read file metadata: %v", err)
	}
	return file, nil
}

func (s *FileStore) Delete(id int) error {
	query := `DELETE FROM file_metadata WHERE id = ?`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("could not delete file metadata: %v", err)
	}
	return nil
}
