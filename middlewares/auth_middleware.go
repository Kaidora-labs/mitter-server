package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaidora-labs/mitter-server/handlers"
	"github.com/kaidora-labs/mitter-server/services"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Result{
				Message: "Authorization header required",
				Error:   "no authorization header provided",
			})

			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Result{
				Message: "Invalid authorization header format",
				Error:   "authorization header must be in the format 'Bearer <token>'",
			})

			return
		}
		token := parts[1]

		claims, err := services.ValidateJWT(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, handlers.Result{
				Message: "Invalid or expired token",
				Error:   err.Error(),
			})

			return
		}

		c.Set(services.ClaimsKey{}, claims)
		c.Next()
	}
}
