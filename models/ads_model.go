package models

import "gorm.io/gorm"

type Ads struct {
	gorm.Model
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Image       string `json:"image" gorm:"type: varchar(255)"`
}
