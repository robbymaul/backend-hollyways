package dtoUser

import "hollyways/models"

type UserResponseDTO struct {
	Email    string               `json:"email"`
	FullName string               `json:"fullName"`
	Status   models.StatusReponse `json:"status"`
	Role     models.RoleResponse  `json:"role"`
	Profile  models.Profile       `json:"profile"`
}

type UserUpdateRequestDTO struct {
	Password string `json:"password" form:"password"`
	FullName string `json:"fullName" form:"fullName"`
}

type UserUpdateByAdminRequestDTO struct {
	StatusID string `json:"statusId" form:"statusId"`
}
