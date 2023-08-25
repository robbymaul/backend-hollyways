package models

import "gorm.io/gorm"

// models structure database table profiles
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

// models response if table joining relation schema
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

// function for handle if user has been register auto create profile on table profiles
func (user *User) AfterCreate(tx *gorm.DB) error {
	profile := &Profile{
		UserID: user.ID,
	}

	err := tx.Create(profile).Error

	return err
}

// funtion for hande not create new table profiles
func (ProfileResponse) TableName() string {
	return "profiles"
}
