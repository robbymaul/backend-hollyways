package models

import "gorm.io/gorm"

// model structure database table brands
type Brand struct {
	gorm.Model
	LogoID int  `json:"logo_id" gorm:"type: int;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Logo   Logo `json:"logo" gorm:"foreignKey:LogoID"`
}

// model response if table joining relation schema
type BrandResponse struct {
	Logo LogoResponse `json:"logo"`
}

// function for handle cant create new table brands
func (BrandResponse) TableName() string {
	return "brands"
}
