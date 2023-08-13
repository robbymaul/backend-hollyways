package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func AuthRoute(g *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerAuth(repo)

	g.POST("/register", handler.Register)
}
