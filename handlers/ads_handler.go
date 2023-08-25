package handlers

import (
	dtoAds "hollyways/dto/ads"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	"hollyways/packages/middleware"
	"hollyways/repositories"
	"hollyways/utility"
	"html"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type adsHandler struct {
	AdsRepository repositories.AdsRepository
}

func HandlerAds(AdsRepository repositories.AdsRepository) *adsHandler {
	return &adsHandler{AdsRepository}
}

func (h *adsHandler) CreateAds(c *gin.Context) {
	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	imageFile, _ := c.Get("file")
	title, err := utility.ValidateInput(c.PostForm("title"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	request := dtoAds.AdsRequestDTO{
		Title:       title,
		Description: c.PostForm("description"),
		Image:       imageFile.(string),
	}

	validation := validator.New()
	if err = validation.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	ads := models.Ads{
		Title:       html.EscapeString(request.Title),
		Description: html.EscapeString(request.Description),
		Image:       request.Image,
	}

	err = h.AdsRepository.CreateAds(ads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to create ads",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "create ads successfully",
	})
}

func (h *adsHandler) FindAds(c *gin.Context) {
	var err error
	var adsDTO []dtoAds.AdsResponseDTO

	ads, err := h.AdsRepository.FindAds()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, ads := range ads {
		adsDTO = append(adsDTO, ConverAdsResponse(ads))
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   adsDTO,
	})
}

func (h *adsHandler) GetUserById(c *gin.Context) {
	var err error

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	adsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		})
		return
	}

	ads, err := h.AdsRepository.GetAdsById(adsId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   ConverAdsResponse(ads),
	})
}

func (h *adsHandler) UpdateAds(c *gin.Context) {
	var err error

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	adsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "invalid parameter",
		})
	}

	imageAds, _ := c.Get("file")
	request := dtoAds.AdsUpdateRequestDTO{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		Image:       imageAds.(string),
	}

	ads, err := h.AdsRepository.GetAdsById(adsId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "data not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if request.Title != "" {
		titleValid, err := utility.ValidateInput(request.Title)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		ads.Title = html.EscapeString(titleValid)
	}

	if request.Description != "" {
		ads.Description = html.EscapeString(request.Description)
	}

	if request.Image != "" {
		ads.Image = request.Image
	}

	err = h.AdsRepository.UpdateAds(ads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to update ads",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "update  ads successfully",
	})
}

func (h *adsHandler) DeleteAds(c *gin.Context) {
	var err error

	authorized := middleware.Authorized(c)
	if authorized != "authorize" {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "unauthorized",
		})
		return
	}

	adsId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "invalid parameter",
		})
		return
	}

	ads, err := h.AdsRepository.GetAdsById(adsId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "data not found, try again",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	err = h.AdsRepository.DeleteAds(ads)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to delete ads",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "delete ads successfully",
	})
}

func ConverAdsResponse(a models.Ads) dtoAds.AdsResponseDTO {
	return dtoAds.AdsResponseDTO{
		ID:          int(a.ID),
		Title:       a.Title,
		Description: a.Description,
		Image:       a.Image,
	}
}
