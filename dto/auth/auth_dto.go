package dtoAuth

import "hollyways/models"

// data transfer object request user will be register
type RegisterRequestDTO struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
}

// data transfer object request if user will login
type LoginRequestDTO struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// data transfer object response if user success login and payload jwt (token)
type LoginResponseDTO struct {
	Email    string         `json:"email"`
	FullName string         `json:"fullName"`
	Status   models.Status  `json:"status"`
	Role     models.Role    `json:"role"`
	Profile  models.Profile `json:"profile"`
	Token    string         `json:"token"`
}

// data transfer object for checking token if user has been login
type CheckAuthResponseDTO struct {
	Email    string              `json:"email"`
	FullName string              `json:"fullName"`
	Role     models.RoleResponse `json:"role"`
	Profile  models.Profile      `json:"profile"`
}
