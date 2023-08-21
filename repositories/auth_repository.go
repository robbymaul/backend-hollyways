package repositories

import "hollyways/models"

type AuthRepository interface {
	Register(user models.User) error
	GetEmailRegistrasi() ([]models.User, error)
	Login(email, password string) (models.User, error)
	CheckAuth(email string, role int) (models.User, error)
}

func (r *repository) Register(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *repository) GetEmailRegistrasi() ([]models.User, error) {
	var user []models.User
	err := r.db.Select("email").Find(&user).Error

	return user, err
}

func (r *repository) Login(email, password string) (models.User, error) {
	var user models.User
	err := r.db.Preload("Status").Preload("Profile").Preload("Role").Where("email = ? AND password = ?", email, password).First(&user).Error

	return user, err
}

func (r *repository) CheckAuth(email string, role int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Role").Where("email = ? AND role_id = ?", email, role).First(&user).Error

	return user, err
}
