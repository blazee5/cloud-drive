package server

import (
	"github.com/blazee5/cloud-drive/api_gateway/lib/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

func (s *Server) UserMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "empty auth header",
		})
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid auth header",
		})
		return
	}

	userID, err := auth.ParseToken(headerParts[1])

	st, _ := status.FromError(err)

	if st.Code() == codes.NotFound {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	if st.Code() == codes.Unauthenticated {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token",
		})
		return
	}

	if err != nil {
		s.log.Error("error while validate token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
		return
	}

	c.Set("userID", userID)
	c.Next()
}
