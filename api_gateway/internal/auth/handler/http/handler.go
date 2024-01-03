package http

import (
	"github.com/blazee5/cloud-drive/api_gateway/internal/auth"
	"github.com/blazee5/cloud-drive/api_gateway/internal/domain"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type Handler struct {
	log         *zap.SugaredLogger
	authService auth.Service
	tracer      trace.Tracer
}

func NewHandler(log *zap.SugaredLogger, authService auth.Service, tracer trace.Tracer) *Handler {
	return &Handler{log: log, authService: authService, tracer: tracer}
}

func (h *Handler) SignUp(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "authHandler.SignUp")
	defer span.End()

	var input domain.SignUpRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	id, err := h.authService.SignUp(ctx, input)

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
	ctx, span := h.tracer.Start(c.Request.Context(), "authHandler.SignIn")
	defer span.End()

	var input domain.SignInRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request",
		})
		return
	}

	token, err := h.authService.SignIn(ctx, input)

	st, ok := status.FromError(err)

	if ok {
		if st.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid credentials",
			})
			return
		}
	}

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

func (h *Handler) ActivateAccount(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "authHandler.ActivateAccount")
	defer span.End()

	code := c.Query("code")

	res, err := h.authService.ActivateAccount(ctx, code)

	st, ok := status.FromError(err)

	if ok {
		if st.Code() == codes.NotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "code not found",
			})
			return
		}
	}

	if err != nil {
		h.log.Infof("error while activate account: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": res,
	})
}
