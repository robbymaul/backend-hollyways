package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func logoRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerLogo(repo)

	r.POST("/logo", middleware.Auth(), middleware.UploadFile(), handler.CreateLogo)
	r.GET("/logos", middleware.Auth(), handler.FindLogo)
}
