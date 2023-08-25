package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	middlewareauth "hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func authRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerAuth(repo)

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/check-auth", middlewareauth.Auth(), handler.CheckAuth)
}
