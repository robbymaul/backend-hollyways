package handlers

import (
	dtoProfile "hollyways/dto/profile"
	dtoResult "hollyways/dto/result"
	"hollyways/repositories"
	"hollyways/utilities"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type profileHandler struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *profileHandler {
	return &profileHandler{ProfileRepository}
}

// function logic handler update profile user
func (h *profileHandler) UpdateProfileByUser(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")
	userId := int(userLogin.(jwt.MapClaims)["id"].(float64))
	profileImage, _ := c.Get("file")

	request := dtoProfile.ProfileUpdateRequestDTO{
		FirstName:    c.PostForm("firstName"),
		LastName:     c.PostForm("lastName"),
		ProfileImage: profileImage.(string),
		Gender:       c.PostForm("gender"),
		PhoneNumber:  c.PostForm("phoneNumber"),
		Address:      c.PostForm("address"),
	}

	profile, err := h.ProfileRepository.GetProfileByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error() + "failed update profile",
		})
		return
	}

	if request.FirstName != "" {
		validateFirstName, err := utilities.ValidateInput(request.FirstName)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		profile.FirstName = html.EscapeString(validateFirstName)
	}

	if request.LastName != "" {
		validateLastName, err := utilities.ValidateInput(request.LastName)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
		profile.LastName = html.EscapeString(validateLastName)
	}

	if request.ProfileImage != "" {
		profile.ProfileImage = html.EscapeString(request.ProfileImage)
	}

	if request.Gender != "" {
		profile.Gender = html.EscapeString(request.Gender)
	}

	if request.PhoneNumber != "" {
		profile.PhoneNumber = html.EscapeString(request.PhoneNumber)
	}

	if request.Address != "" {
		profile.Address = html.EscapeString(request.Address)
	}

	err = h.ProfileRepository.UpdateProfileByUser(profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Failed update profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Update profile successfully",
	})
}
