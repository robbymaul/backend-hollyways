package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID       uint   `json:"user_id" gorm:"type: int(11);constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User         User   `json:"user" gorm:"foreignKey:UserID"`
	FirstName    string `json:"firstName" gorm:"type: varchar (255)"`
	LastName     string `json:"lastName" gorm:"type: varchar (255)"`
	ProfileImage string `json:"image" gorm:"type: varchar (255)"`
	Gender       string `json:"gender" gorm:"type: varchar (255)"`
	PhoneNumber  string `json:"phoneNumber" gorm:"type: varchar(255)"`
	Address      string `json:"address" gorm:"type: varchar(255)"`
}

type ProfileResponse struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"image"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
