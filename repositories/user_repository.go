package repositories

import "hollyways/models"

type UserRepository interface {
	FindUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) error
	UpdateUserByAdmin(user models.User) error
	DeleteUserByAdmin(user models.User) error
}

func (r *repository) FindUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Preload("Profile").Preload("Role").Find(&user).Error

	return user, err
}

func (r *repository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Role").First(&user, "id = ?", id).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *repository) UpdateUserByAdmin(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

func (r *repository) DeleteUserByAdmin(user models.User) error {
	err := r.db.Delete(&user).Error

	return err
}
