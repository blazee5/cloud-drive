package http

import (
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/file"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	log         *zap.SugaredLogger
	fileService file.Service
}

func NewHandler(log *zap.SugaredLogger, fileService file.Service) *Handler {
	return &Handler{log: log, fileService: fileService}
}

func (h *Handler) GetUserFiles(c *gin.Context) {
	userID, ok := c.Get("userID")

	if ok != true {
		return
	}

	files, err := h.fileService.GetFiles(c, userID.(string))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) UploadFile(c *gin.Context) {
	userID, ok := c.Get("userID")

	if ok != true {
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while open file",
		})
		return
	}

	id, err := h.fileService.UploadFile(c, userID.(string), file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
