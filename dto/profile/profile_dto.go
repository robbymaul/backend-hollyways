package dtoProfile

import "hollyways/models"

type ProfileUpdateRequestDTO struct {
	profile      models.Profile
	FirstName    string `json:"firstName" form:"firstName"`
	LastName     string `json:"lastName" form:"lastName"`
	ProfileImage string `json:"image" form:"image"`
	Gender       string `json:"gender" form:"gender"`
	PhoneNumber  string `json:"phoneNumber" form:"phoneNumber"`
	Address      string `json:"address" form:"address"`
}
