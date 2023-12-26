package domain

import (
	"time"
)

type FileInfo struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	UserID        string    `json:"user_id"`
	DownloadCount int       `json:"download_count"`
	CreatedAt     time.Time `json:"created_at"`
}

type FileList struct {
	Total      int        `json:"total"`
	TotalPages int        `json:"total_pages"`
	Page       int        `json:"page"`
	Size       int        `json:"size"`
	Files      []FileInfo `json:"files"`
}

type UpdateFileInput struct {
	FileName string `json:"filename" binding:"required"`
}
