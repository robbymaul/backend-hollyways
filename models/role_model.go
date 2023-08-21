package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model `json:"-"`
	Role       string `json:"role" gorm:"type:varchar (255)"`
}

type RoleResponse struct {
	gorm.Model `json:"-"`
	Role       string `json:"role"`
}

func (RoleResponse) TableName() string {
	return "roles"
}
