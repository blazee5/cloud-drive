package http

import (
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/auth"
	"github.com/blazee5/cloud-drive/microservices/api_gateway/internal/domain"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	log         *zap.SugaredLogger
	authService auth.Service
}

func NewHandler(log *zap.SugaredLogger, authService auth.Service) *Handler {
	return &Handler{log: log, authService: authService}
}

func (h *Handler) SignUp(c *gin.Context) {
	var input domain.SignUpRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	id, err := h.authService.SignUp(c, input)

	if err != nil {
		h.log.Infof("error while sign up: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input domain.SignInRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	token, err := h.authService.SignIn(c, input)

	if err != nil {
		h.log.Infof("error while sign in: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}