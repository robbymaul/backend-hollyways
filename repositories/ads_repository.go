package repositories

import "hollyways/models"

type AdsRepository interface {
	CreateAds(ads models.Ads) error
	FindAds() ([]models.Ads, error)
	GetAdsById(id int) (models.Ads, error)
	UpdateAds(ads models.Ads) error
	DeleteAds(ads models.Ads) error
}

func (r *repository) CreateAds(ads models.Ads) error {
	err := r.db.Create(&ads).Error

	return err
}

func (r *repository) FindAds() ([]models.Ads, error) {
	var ads []models.Ads
	err := r.db.Find(&ads).Error

	return ads, err
}

func (r *repository) GetAdsById(id int) (models.Ads, error) {
	var ads models.Ads
	err := r.db.First(&ads, "id = ?", id).Error

	return ads, err
}

func (r *repository) UpdateAds(ads models.Ads) error {
	err := r.db.Save(&ads).Error

	return err
}

func (r *repository) DeleteAds(ads models.Ads) error {
	err := r.db.Delete(&ads).Error

	return err
}
