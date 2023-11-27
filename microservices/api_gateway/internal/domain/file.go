package domain

type UpdateFileInput struct {
	FileName string `json:"file_name" binding:"required"`
}
