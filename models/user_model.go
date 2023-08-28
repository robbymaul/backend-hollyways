package models

import (
	"gorm.io/gorm"
)

// models stucture databse table users
type User struct {
	gorm.Model
	Email    string  `json:"email" gorm:"type: varchar(255);uniqueIndex"`
	Password string  `gorm:"type: varchar(255)"`
	FullName string  `json:"fullname" gorm:"type: varchar(255)"`
	StatusID int     `json:"statusId" gorm:"type: int"`
	Status   Status  `json:"status" gorm:"foreignKey:StatusID"`
	RoleID   uint    `json:"roleId" gorm:"type: int"`
	Role     Role    `json:"role" gorm:"foreignKey:RoleID"`
	Profile  Profile `json:"profile"`
}

// models response if table joining relation schema
type UserResponse struct {
	gorm.Model `json:"-"`
	Email      string  `json:"email" gorm:"type: varchar(255)"`
	Password   string  `json:"-" gorm:"type: varchar(255)"`
	StatusID   int     `json:"-" gorm:"type: int"`
	FullName   string  `json:"fullname" gorm:"type: varchar(255)"`
	Status     Status  `json:"status" gorm:"foreignKey:StatusID"`
	RoleID     uint    `json:"-" gorm:"type: int"`
	Role       Role    `json:"role" gorm:"foreignKey:RoleID"`
	Profile    Profile `json:"profile"`
}

// function for handle not create new table users
func (UserResponse) TableName() string {
	return "users"
}
