package repositories

import "hollyways/models"

type AuthRepository interface {
	Register(user models.User) error
	GetEmailRegistrasi(email string) ([]models.User, error)
}

func (r *repository) Register(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *repository) GetEmailRegistrasi(email string) ([]models.User, error) {
	var user []models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}
