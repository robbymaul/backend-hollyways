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

func (logo *Logo) AfterCreate(tx *gorm.DB) error {
	var brand Brand
	err := tx.Raw("UPDATE brands SET logo_id =?, created_at = ?, updated_at = ? WHERE id = 1", logo.ID, logo.CreatedAt, logo.UpdatedAt).Scan(&brand).Error

	return err
}

// function for handle cant create new table brands
func (BrandResponse) TableName() string {
	return "brands"
}
