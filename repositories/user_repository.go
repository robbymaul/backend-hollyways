package repositories

import "hollyways/models"

type UserRepository interface {
	FindUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(user models.User) error
	UpdateUserByAdmin(user models.User) error
	DeleteUserByAdmin(user models.User) error
}

// function get all users data in table users with ORM
func (r *repository) FindUser() ([]models.User, error) {
	var user []models.User
	err := r.db.Preload("Profile").Preload("Role").Find(&user).Error

	return user, err
}

// function select specific data user in table user with ORM, checking by field id
func (r *repository) GetUserById(id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Role").First(&user, "id = ?", id).Error

	return user, err
}

// function update user by user will be update data table users with ORM
func (r *repository) UpdateUser(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

// function update user by admin will be update data table users with ORM
func (r *repository) UpdateUserByAdmin(user models.User) error {
	err := r.db.Save(&user).Error

	return err
}

// function delete user by admin will be delete data users in table user with ORM
func (r *repository) DeleteUserByAdmin(user models.User) error {
	err := r.db.Delete(&user).Error

	return err
}
