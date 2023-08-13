package models

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	LogoID int  `json:"logo_id" gorm:"type: int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logo   Logo `json:"logo" gorm:"foreignKey:LogoID"`
}

type BrandResponse struct {
	Logo LogoResponse `json:"logo"`
}

func (BrandResponse) TableName() string {
	return "brands"
}
