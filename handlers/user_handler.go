package handlers

import (
	dtoResult "hollyways/dto/result"
	dtoUser "hollyways/dto/user"
	"hollyways/models"
	"hollyways/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type userHandler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *userHandler {
	return &userHandler{UserRepository}
}

func (h *userHandler) FindUser(c *gin.Context) {
	var userDTO []dtoUser.UserResponseDTO

	userLogin, _ := c.Get("userLogin")

	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))

	if userRole == 3 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	user, err := h.UserRepository.FindUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	for _, user := range user {
		userDTO = append(userDTO, ConvertUserResponse(user))
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   userDTO,
	})
}

func (h *userHandler) GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   ConvertUserResponse(user),
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")

	userID := int(userLogin.(jwt.MapClaims)["id"].(float64))
	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))

	if userRole != 3 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	request := new(dtoUser.UserUpdateRequestDTO)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "the data entered is incorrect",
		})
		return
	}

	user, err := h.UserRepository.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	err = h.UserRepository.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Updated Successfully",
	})
}

func (h *userHandler) UpdateUserByAdmin(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")
	userID, _ := strconv.Atoi(c.Param("id"))

	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))

	if userRole == 3 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	request := new(dtoUser.UserUpdateByAdminRequestDTO)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "the data entered is incorrect",
		})
	}

	user, err := h.UserRepository.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "Data Not Found",
		})
		return
	}

	if request.StatusID == "" {
		parseStatusId, _ := strconv.Atoi(request.StatusID)
		user.StatusID = parseStatusId
	}

	err = h.UserRepository.UpdateUserByAdmin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to update",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "updated successfully",
	})
}

func (h *userHandler) DeleteUserByAdmin(c *gin.Context) {
	userLogin, _ := c.Get("userLogin")
	userId, _ := strconv.Atoi(c.Param("id"))

	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))

	if userRole == 3 {
		c.JSON(http.StatusUnauthorized, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	user, err := h.UserRepository.GetUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "data not found",
		})
		return
	}

	err = h.UserRepository.DeleteUserByAdmin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to deleted",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "deleted successfully",
	})
}

func ConvertUserResponse(u models.User) dtoUser.UserResponseDTO {
	return dtoUser.UserResponseDTO{
		Email:    u.Email,
		FullName: u.FullName,
		Status:   models.StatusReponse(u.Status),
		Role:     models.RoleResponse(u.Role),
		Profile:  u.Profile,
	}
}
