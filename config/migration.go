package config

import (
	"fmt"
	"hollyways/models"
	"hollyways/packages/connection"
)

// function auto migrate model to database
func Migration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Role{},
		&models.Status{},
		&models.Transaction{},
		&models.Project{},
		&models.Brand{},
		&models.Logo{},
		&models.Ads{},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Migration success")
}
