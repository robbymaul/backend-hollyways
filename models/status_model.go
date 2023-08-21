package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model `json:"-"`
	Status     string `json:"status" gorm:"type:varchar (255)"`
}

type StatusReponse struct {
	gorm.Model `json:"-"`
	Status     string `json:"status" gorm:"type:varchar (255)"`
}

func (StatusReponse) TableName() string {
	return "statuses"
}
