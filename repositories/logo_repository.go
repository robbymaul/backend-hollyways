package repositories

import "hollyways/models"

type LogoRepository interface {
	CreateLogo(logo models.Logo) error
	FindLogo() ([]models.Logo, error)
}

func (r *repository) CreateLogo(logo models.Logo) error {
	err := r.db.Create(&logo).Error

	return err
}

func (r *repository) FindLogo() ([]models.Logo, error) {
	var logo []models.Logo
	err := r.db.Find(&logo).Error

	return logo, err
}
