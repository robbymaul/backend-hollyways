package repositories

import "hollyways/models"

// contract interface function query database model
type AuthRepository interface {
	Register(user models.User) error
	Login(email, password string) (models.User, error)
	CheckAuth(email string, role int) (models.User, error)
}

// function register user table users with ORM
func (r *repository) Register(user models.User) error {
	err := r.db.Create(&user).Error

	return err
}

// function login user table users with ORM, checking by email and password user
func (r *repository) Login(email, password string) (models.User, error) {
	var user models.User
	err := r.db.Preload("Status").Preload("Profile").Preload("Role").Where("email = ? AND password = ?", email, password).First(&user).Error

	return user, err
}

// function checkauth if user has been login with ORM, checking by email and role user
func (r *repository) CheckAuth(email string, role int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Role").Where("email = ? AND role_id = ?", email, role).First(&user).Error

	return user, err
}
