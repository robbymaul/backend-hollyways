package handlers

import (
	dtoAuth "hollyways/dto/auth"
	dtoResult "hollyways/dto/result"
	"hollyways/models"
	jwtauth "hollyways/packages/jwt"
	"hollyways/repositories"
	"hollyways/utilities"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c *gin.Context) {
	request := new(dtoAuth.RegisterRequestDTO)
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	emailValid, err := utilities.ValidateAndSanitazeEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "invalid email",
		})
		return
	}

	fullNameValid, err := utilities.ValidateInput(request.FullName)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user := models.User{
		Email:    emailValid,
		Password: request.Password,
		FullName: fullNameValid,
		RoleID:   3,
		StatusID: 1,
	}

	err = h.AuthRepository.Register(user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
				Status:  http.StatusBadRequest,
				Message: "Email already exists",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Successful registration",
	})
}

func (h *handlerAuth) Login(c *gin.Context) {
	request := new(dtoAuth.LoginRequestDTO)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.AuthRepository.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtoResult.ErrorResult{
			Status:  http.StatusBadRequest,
			Message: "wrong email or password",
		})
		return
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["fullName"] = user.FullName
	claims["role"] = user.Role.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, err := jwtauth.GenerateToken(&claims)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: "failed to generate token",
		})
		return
	}

	response := dtoAuth.LoginResponseDTO{
		Email:    user.Email,
		FullName: user.FullName,
		Profile:  user.Profile,
		Role:     user.Role,
		Status:   user.Status,
		Token:    token,
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (h *handlerAuth) CheckAuth(c *gin.Context) {
	userLogin, exists := c.Get("userLogin")
	if !exists {
		c.JSON(http.StatusUnsupportedMediaType, dtoResult.ErrorResult{
			Status:  http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	userEmail := userLogin.(jwt.MapClaims)["email"].(string)
	userRole := int(userLogin.(jwt.MapClaims)["role"].(float64))

	user, err := h.AuthRepository.CheckAuth(userEmail, userRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtoResult.ErrorResult{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dtoResult.SuccessResult{
		Status: http.StatusOK,
		Data:   ConvertCheckAuthResp(user),
	})
}

func ConvertCheckAuthResp(u models.User) dtoAuth.CheckAuthResponseDTO {
	return dtoAuth.CheckAuthResponseDTO{
		Email:    u.Email,
		FullName: u.FullName,
		Role:     models.RoleResponse(u.Role),
		Profile:  u.Profile,
	}
}
