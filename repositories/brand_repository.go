package repositories

import "hollyways/models"

type BrandRepository interface {
	GetBrand() (models.Brand, error)
}

func (r *repository) GetBrand() (models.Brand, error) {
	var brand models.Brand
	err := r.db.Preload("Logo").First(&brand).Error

	return brand, err
}
