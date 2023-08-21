package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model   `json:"-"`
	UserID       uint   `json:"user_id" gorm:"type: int(11);constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	FirstName    string `json:"firstName" gorm:"type: varchar (255)"`
	LastName     string `json:"lastName" gorm:"type: varchar (255)"`
	ProfileImage string `json:"image" gorm:"type: varchar (255)"`
	Gender       string `json:"gender" gorm:"type: varchar (255)"`
	PhoneNumber  string `json:"phoneNumber" gorm:"type: varchar(255)"`
	Address      string `json:"address" gorm:"type: varchar(255)"`
}

type ProfileResponse struct {
	gorm.Model   `json:"-"`
	UserID       int    `json:"-"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"image"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phoneNumber"`
	Address      string `json:"address"`
}

func (user *User) AfterCreate(tx *gorm.DB) error {
	profile := &Profile{
		UserID: user.ID,
	}

	err := tx.Create(profile).Error

	return err
}

func (ProfileResponse) TableName() string {
	return "profiles"
}
