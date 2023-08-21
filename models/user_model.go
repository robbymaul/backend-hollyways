package models

import (
	"gorm.io/gorm"
)

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

type UserResponse struct {
	Email    string `json:"email" gorm:"type: varchar(255)"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
}

func (UserResponse) TableName() string {
	return "users"
}
