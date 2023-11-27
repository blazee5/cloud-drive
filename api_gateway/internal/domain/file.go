package domain

type UpdateFileInput struct {
	FileName string `json:"filename" binding:"required"`
}
