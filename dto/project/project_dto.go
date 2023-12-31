package projectdto

import (
	"time"
)

// data transfer object request create if admin will be create project donation
type ProjectRequestDTO struct {
	ProjectName        string    `json:"projectName" form:"projectName" validate:"required"`
	ProjectDescription string    `json:"projectDescription" form:"projectDescription" validate:"required"`
	ProjectImage       string    `json:"image" form:"image" validate:"required"`
	TargetDonation     int       `json:"target" form:"target" validate:"required"`
	StartDate          time.Time `json:"startDate" form:"startDate" validate:"required"`
	DueDate            time.Time `json:"dueDate" form:"dueDate" validate:"required"`
}

// data transfer object response if success get data project
type ProjectResponseDTO struct {
	ID                 int     `json:"id"`
	ProjectName        string  `json:"projectName"`
	ProjectDescription string  `json:"projectDescription"`
	ProjectImage       string  `json:"image"`
	Donation           int     `json:"donation"`
	TargetDonation     int     `json:"target"`
	StartDate          string  `json:"startDate"`
	DueDate            string  `json:"dueDate"`
	Progress           float64 `json:"progress"`
}

// data transfer object request update if admin will be update project donation
type ProjectUpdateRequestDTO struct {
	ProjectName        string    `json:"projectName" form:"projectName"`
	ProjectDescription string    `json:"projectDescription" form:"projectDescritpion"`
	ProjectImage       string    `json:"image" form:"image"`
	DueDate            time.Time `json:"dueDate" form:"dueDate"`
}
