package models

import "gorm.io/gorm"

type Logo struct {
	gorm.Model
	Image string `json:"image" gorm:"type: varchar(255)"`
}

type LogoResponse struct {
	Image string `json:"image"`
}

func (LogoResponse) TableName() string {
	return "logos"
}
