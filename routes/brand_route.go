package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func branRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerBrand(repo)

	r.GET("/brand", handler.GetBrand)
}
