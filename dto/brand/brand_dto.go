package dtoBrand

import "hollyways/models"

type BrandRequestUpdateDTO struct {
	LogoID string `json:"logo_id" form:"logo_id"`
}

type BrandResponseDTO struct {
	Logo models.LogoResponse `json:"brand"`
}
