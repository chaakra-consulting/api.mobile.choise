package middlewares

import (
	"errors"
	"strings"

	"api.choise/internals/helpers"
	"github.com/gin-gonic/gin"
)

// Middleware function to check user token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			helpers.ErrorResponse(c, 401, c.Error(errors.New(("Unauthorized: Missing or invalid token"))))
			c.Abort()
			return
		}
		c.Next()
	}
}
