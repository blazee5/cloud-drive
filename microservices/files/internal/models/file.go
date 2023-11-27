package models

type File struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserID string `json:"user_id" db:"user_id"`
	Chunk  []byte `json:"chunk" db:"chunk"`
}

type FileInfo struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserID string `json:"user_id" db:"user_id"`
}
