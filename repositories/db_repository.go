package repositories

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

// function return value repository db connection for use *gorm.DB
func MakeRepository(db *gorm.DB) *repository {
	return &repository{db}
}
