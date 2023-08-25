package models

import "gorm.io/gorm"

// models structure database table roles
type Role struct {
	gorm.Model `json:"-"`
	Role       string `json:"role" gorm:"type:varchar (255)"`
}

// models response if table joining relation schema
type RoleResponse struct {
	gorm.Model `json:"-"`
	Role       string `json:"role"`
}

// function for handle not create new table roles
func (RoleResponse) TableName() string {
	return "roles"
}
