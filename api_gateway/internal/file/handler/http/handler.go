package http

import (
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"github.com/blazee5/cloud-drive/api_gateway/internal/file"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
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

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid page query",
		})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "0"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "invalid size query",
		})
		return
	}

	if ok != true {
		return
	}

	files, err := h.fileService.GetFiles(c, userID.(string), page, size)

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
		h.log.Infof("error while open file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error while open file",
		})
		return
	}

	id, err := h.fileService.UploadFile(c, userID.(string), file)

	if err != nil {
		h.log.Infof("error while upload file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	userID, ok := c.Get("userID")

	if ok != true {
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid file id",
		})
	}

	file, err := h.fileService.DownloadFile(c, ID, userID.(string))

	st, _ := status.FromError(err)

	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "file not found",
		})
		return
	}

	if err != nil {
		h.log.Infof("error while download file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.Data(http.StatusOK, file.GetContentType(), file.GetChunk())
}

func (h *Handler) UpdateFile(c *gin.Context) {
	var input domain.UpdateFileInput

	userID, ok := c.Get("userID")

	if ok != true {
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid file id",
		})
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	err = h.fileService.UpdateFile(c, ID, userID.(string), input)

	st, _ := status.FromError(err)

	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "file not found",
		})
		return
	}

	if err != nil {
		h.log.Infof("error while update file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}

func (h *Handler) DeleteFile(c *gin.Context) {
	userID, ok := c.Get("userID")

	if ok != true {
		return
	}

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid file id",
		})
	}

	err = h.fileService.DeleteFile(c, ID, userID.(string))

	st, _ := status.FromError(err)

	if st.Code() == codes.NotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "file not found",
		})
		return
	}

	if err != nil {
		h.log.Infof("error while delete file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.String(http.StatusOK, "OK")
}
