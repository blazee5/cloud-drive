package domain

type File struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	UserID      string `json:"user_id" db:"user_id"`
	ContentType string `json:"content_type" db:"content_type"`
	Chunk       []byte `json:"chunk" db:"chunk"`
}
