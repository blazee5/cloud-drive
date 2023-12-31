package http

import (
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"github.com/blazee5/cloud-drive/api_gateway/internal/file"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

type Handler struct {
	log         *zap.SugaredLogger
	fileService file.Service
	tracer      trace.Tracer
}

func NewHandler(log *zap.SugaredLogger, fileService file.Service, trace trace.Tracer) *Handler {
	return &Handler{log: log, fileService: fileService, tracer: trace}
}

func (h *Handler) GetUserFiles(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "filesHandler.GetUserFiles")
	defer span.End()

	userID, ok := c.Get("userID")

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))

	orderBy := c.DefaultQuery("sortBy", "name")
	orderDir := c.DefaultQuery("sortDir", "asc")

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

	files, err := h.fileService.GetFiles(ctx, userID.(string), orderBy, orderDir, page, size)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, files)
}

func (h *Handler) UploadFile(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "filesHandler.UploadFile")
	defer span.End()

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

	id, err := h.fileService.UploadFile(ctx, userID.(string), file)

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
	ctx, span := h.tracer.Start(c.Request.Context(), "filesHandler.DownloadFile")
	defer span.End()

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

	file, err := h.fileService.DownloadFile(ctx, ID, userID.(string))

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
	ctx, span := h.tracer.Start(c.Request.Context(), "filesHandler.UpdateFile")
	defer span.End()

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

	err = h.fileService.UpdateFile(ctx, ID, userID.(string), input)

	st, ok := status.FromError(err)

	if ok {
		if st.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "file not found",
			})
			return
		}
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
	ctx, span := h.tracer.Start(c.Request.Context(), "filesHandler.DeleteFile")
	defer span.End()

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

	err = h.fileService.DeleteFile(ctx, ID, userID.(string))

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
