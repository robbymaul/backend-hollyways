package repositories

import "hollyways/models"

// contract interface funtion profile repository for model profile
type ProfileRepository interface {
	GetProfileByUserId(id int) (models.Profile, error)
	UpdateProfileByUser(profile models.Profile) error
}

// function select specific profile data in table profiles checking by user_id
func (r *repository) GetProfileByUserId(id int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.First(&profile, "user_id = ?", id).Error

	return profile, err
}

// function update profile by user in tbale profiles (with ORM)
func (r *repository) UpdateProfileByUser(profile models.Profile) error {
	err := r.db.Save(&profile).Error

	return err
}
