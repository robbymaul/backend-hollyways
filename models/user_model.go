package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string   `json:"email" gorm:"type: varchar(255)"`
	Password string   `gorm:"type: varchar(255)"`
	FullName string   `json:"fullname" gorm:"type: varchar(255)"`
	Profile  *Profile `json:"profile"`
}

type UserResponse struct {
	Email    string `json:"email" gorm:"type: varchar(255)"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
}

func (UserResponse) TableName() string {
	return "users"
}
