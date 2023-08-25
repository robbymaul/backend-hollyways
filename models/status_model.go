package models

import "gorm.io/gorm"

// models structure database table statuses
type Status struct {
	gorm.Model `json:"-"`
	Status     string `json:"status" gorm:"type:varchar (255)"`
}

// models response if table joining relation schema
type StatusReponse struct {
	gorm.Model `json:"-"`
	Status     string `json:"status" gorm:"type:varchar (255)"`
}

// function for handle not create new table statuses
func (StatusReponse) TableName() string {
	return "statuses"
}
