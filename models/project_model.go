package models

import (
	"time"

	"gorm.io/gorm"
)

// models structure database table projects
type Project struct {
	gorm.Model
	ProjectName        string    `json:"projectName" gorm:"type: varchar(255)"`
	ProjectDescription string    `json:"projectDescription" gorm:"type: varchar(255)"`
	ProjectImage       string    `json:"image" gorm:"type: varchar(255)"`
	Donation           int       `json:"donation" gorm:"type: int"`
	TargetDonation     int       `json:"target" gorm:"type: int"`
	StartDate          time.Time `json:"startDate" gorm:"type: date"`
	DueDate            time.Time `json:"dueDate" gorm:"type: date"`
	Progress           float64   `json:"progress" gorm:"type: float"`
}

// modesl response if table joining relation schema
type ProjectResponse struct {
	gorm.Model         `json:"-"`
	ProjectName        string    `json:"projectName"`
	ProjectDescription string    `json:"projectDescription"`
	ProjectImage       string    `json:"image"`
	Donation           int       `json:"donation"`
	TargetDonation     int       `json:"target"`
	StartDate          time.Time `json:"startDate"`
	DueDate            time.Time `json:"dueDate"`
	Progress           float64   `json:"-" gorm:"type: float"`
}

// function for handle not create new table
func (ProjectResponse) TableName() string {
	return "projects"
}
