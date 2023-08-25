package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func profileRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.HandlerProfile(repo)

	r.PATCH("/profile", middleware.Auth(), middleware.UploadFile(), handler.UpdateProfileByUser)
}
