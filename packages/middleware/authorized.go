package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authorized(c *gin.Context) string {
	userLogin, exists := c.Get("userLogin")
	if !exists {
		return "unauthorized"
	}

	if userLogin == "" {
		return "unauthorized"
	}

	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))
	if userRole != 1 && userRole != 2 {
		return "unauthorized"
	}

	return "authorize"
}
