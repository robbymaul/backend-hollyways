package dtoTransaction

import (
	projectdto "hollyways/dto/project"
	dtoUser "hollyways/dto/user"
)

type TransactionRequestDTO struct {
	ProjectID int `json:"projectId" form:"projectId" validate:"required"`
	Donation  int `json:"donation" form:"donation" validate:"required"`
}

type TransactionResponseDTO struct {
	ID       int                           `json:"id"`
	User     dtoUser.UserResponseDTO       `json:"user"`
	Project  projectdto.ProjectResponseDTO `json:"project"`
	Donation int                           `json:"donation"`
	Status   string                        `json:"status"`
}
