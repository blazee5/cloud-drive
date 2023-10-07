package models

import "time"

type File struct {
	Id            int       `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	DownloadCount int       `json:"download_count" db:"download_count"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
