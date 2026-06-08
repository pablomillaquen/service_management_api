package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pablomillaquen/speckit_golang_api/pkg/response"
)

func RequiredRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			response.Unauthorized(c, "Authentication required")
			return
		}
		role, ok := userRole.(string)
		if !ok {
			response.Unauthorized(c, "Invalid user role")
			return
		}
		for _, r := range roles {
			if role == r {
				c.Next()
				return
			}
		}
		response.Forbidden(c, "Insufficient permissions")
	}
}
