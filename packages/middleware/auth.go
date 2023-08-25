package middleware

import (
	dtoResult "hollyways/dto/result"
	jwtauth "hollyways/packages/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// function authectication user login by JWT
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.Set("userLogin", "")
			c.Next()
			return
		}

		token = strings.Split(token, " ")[1]

		claims, err := jwtauth.DecodeToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
				Status:  http.StatusUnauthorized,
				Message: "Unauthorized",
			})
			return
		}

		c.Set("userLogin", claims)
		c.Next()
	}
}
