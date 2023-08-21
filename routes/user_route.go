package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerUser(repo)

	r.GET("/users", middleware.Auth(), handler.FindUser)
	r.GET("/user/:id", middleware.Auth(), handler.GetUserById)
	r.PATCH("user", middleware.Auth(), handler.UpdateUser)
}
