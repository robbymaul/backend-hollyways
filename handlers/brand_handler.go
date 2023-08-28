package handlers

import (
	dtoBrand "hollyways/dto/brand"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	"hollyways/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type brandHandler struct {
	BrandRepository repositories.BrandRepository
}

func HandlerBrand(BrandRepository repositories.BrandRepository) *brandHandler {
	return &brandHandler{BrandRepository}
}

func (h *brandHandler) GetBrand(c *gin.Context) {
	var err error

	brand, err := h.BrandRepository.GetBrand()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   convertResponseBrand(brand),
	})
}

func convertResponseBrand(b models.Brand) dtoBrand.BrandResponseDTO {
	return dtoBrand.BrandResponseDTO{
		Logo: models.LogoResponse{Image: b.Logo.Image},
	}
}
