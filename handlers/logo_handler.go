package handlers

import (
	dtoLogo "hollyways/dto/logo"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	"hollyways/packages/middleware"
	"hollyways/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type logoHandler struct {
	LogoRepository repositories.LogoRepository
}

func HandlerLogo(LogoRepository repositories.LogoRepository) *logoHandler {
	return &logoHandler{LogoRepository}
}

func (h *logoHandler) CreateLogo(c *gin.Context) {
	var err error

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	imageFile, _ := c.Get("file")
	if imageFile == "" {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "required input data file",
		})
		return
	}

	request := dtoLogo.LogoRequestDTO{
		Image: imageFile.(string),
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "Validation error",
		})
		return
	}

	logo := models.Logo{
		Image: request.Image,
	}

	err = h.LogoRepository.CreateLogo(logo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Failed to create data logo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Create data logo successfully",
	})
}

func (h *logoHandler) FindLogo(c *gin.Context) {
	var err error
	var dtoLogo []dtoLogo.LogoResponseDTO

	logo, err := h.LogoRepository.FindLogo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(logo) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Data logo not found",
			"data":    logo,
		})
		return
	}

	for _, logo := range logo {
		dtoLogo = append(dtoLogo, convertResponseLogo(logo))
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   dtoLogo,
	})

}

func convertResponseLogo(l models.Logo) dtoLogo.LogoResponseDTO {
	return dtoLogo.LogoResponseDTO{
		ID:    int(l.ID),
		Image: l.Image,
	}
}
