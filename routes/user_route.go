package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoute(g *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerUser(repo)

	g.GET("/users", handler.FindUser)
}
