package dtoUser

import "hollyways/models"

// data transfer object for respone json if get data success
type UserResponseDTO struct {
	Email    string               `json:"email"`
	FullName string               `json:"fullName"`
	Status   models.StatusReponse `json:"status"`
	Role     models.RoleResponse  `json:"role"`
	Profile  models.Profile       `json:"profile"`
}

// data transfer object for update request if user will be update data
type UserUpdateRequestDTO struct {
	Password string `json:"password" form:"password"`
	FullName string `json:"fullName" form:"fullName"`
}

// data transfer object for update by admin will be controller user status
type UserUpdateByAdminRequestDTO struct {
	StatusID string `json:"statusId" form:"statusId"`
}
