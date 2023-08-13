package handlers

import (
	dtoResult "hollyways/dto/result"
	"hollyways/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *userHandler {
	return &userHandler{UserRepository}
}

func (h *userHandler) FindUser(g *gin.Context) {
	user, err := h.UserRepository.FindUser()
	if err != nil {
		g.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   user,
	})
}
