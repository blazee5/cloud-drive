package models

type File struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	UserId string `json:"user_id" db:"user_id"`
	Chunk  []byte `json:"chunk" db:"chunk"`
}
