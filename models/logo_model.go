package models

import "gorm.io/gorm"

// model structure database table logos
type Logo struct {
	gorm.Model `json:"-"`
	Image      string `json:"image" gorm:"type: varchar(255)"`
}

// models response if table joining relation schema
type LogoResponse struct {
	Image string `json:"image"`
}

// function for handle not create new table logos
func (LogoResponse) TableName() string {
	return "logos"
}
