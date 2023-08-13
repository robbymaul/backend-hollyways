package repositories

import "hollyways/models"

type UserRepository interface {
	FindUser() ([]models.User, error)
}

func (r *repository) FindUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Preload("Profile").Find(&user).Error

	return user, err
}
