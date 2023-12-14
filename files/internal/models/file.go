package models

import (
	pb "github.com/blazee5/cloud-drive-protos/files"
	"time"
)

type File struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	UserID        string    `json:"user_id" db:"user_id"`
	ContentType   string    `json:"content_type" db:"content_type"`
	DownloadCount int       `json:"download_count" db:"download_count"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	Chunk         []byte
}

type FileInfo struct {
	ID            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	UserID        string    `json:"user_id" db:"user_id"`
	DownloadCount int       `json:"download_count" db:"download_count"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type FileList struct {
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
	Page       int            `json:"page"`
	Size       int            `json:"size"`
	Files      []*pb.FileInfo `json:"files"`
}
