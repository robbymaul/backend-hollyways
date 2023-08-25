package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func adsRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerAds(repo)

	r.POST("/ads", middleware.Auth(), middleware.UploadFile(), handler.CreateAds)
	r.GET("/adses", handler.FindAds)
	r.GET("/ads/:id", middleware.Auth(), handler.GetUserById)
	r.PATCH("/ads/:id", middleware.Auth(), middleware.UploadFile(), handler.UpdateAds)
	r.DELETE("/ads/:id", middleware.Auth(), handler.DeleteAds)
}
