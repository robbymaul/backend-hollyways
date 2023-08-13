package handlers

import (
	dtoAuth "hollyways/dto/auth"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	"hollyways/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(g *gin.Context) {
	request := new(dtoAuth.RegisterRequestDTO)
	if err := g.ShouldBind(&request); err != nil {
		g.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	email, _ := h.AuthRepository.GetEmailRegistrasi(request.Email)
	if len(email) > 0 {
		g.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "email already exists",
		})
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
		FullName: request.FullName,
	}
	err = h.AuthRepository.Register(user)
	if err != nil {
		g.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successful registration",
	})

}
