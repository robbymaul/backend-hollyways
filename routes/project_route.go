package routes

import (
	"hollyways/handlers"
	"hollyways/packages/connection"
	"hollyways/packages/middleware"
	"hollyways/repositories"

	"github.com/gin-gonic/gin"
)

func projectRoute(r *gin.RouterGroup) {
	repo := repositories.MakeRepository(connection.DB)
	handler := handlers.ProjectHandler(repo)

	r.POST("/project", middleware.Auth(), middleware.UploadFile(), handler.CreateProject)
	r.GET("/projects", handler.FindProject)
	r.GET("/project/:id", handler.GetProject)
	r.PATCH("/project/:id", middleware.Auth(), middleware.UploadFile(), handler.UpdateProjectByAdmin)
	r.DELETE("project/:id", middleware.Auth(), handler.DeleteProjectByAdmin)
}
