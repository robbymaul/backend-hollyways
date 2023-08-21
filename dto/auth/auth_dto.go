package dtoAuth

import "hollyways/models"

type RegisterRequestDTO struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullName" form:"fullName" validate:"required"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type LoginResponseDTO struct {
	Email    string         `json:"email"`
	FullName string         `json:"fullName"`
	Status   models.Status  `json:"status"`
	Role     models.Role    `json:"role"`
	Profile  models.Profile `json:"profile"`
	Token    string         `json:"token"`
}

type CheckAuthResponseDTO struct {
	Email    string              `json:"email"`
	FullName string              `json:"fullName"`
	Role     models.RoleResponse `json:"role"`
	Profile  models.Profile      `json:"profile"`
}
