package main

import (
	"github.com/wigwamwam/CRUD_app/initializers"
	"github.com/wigwamwam/CRUD_app/models"
)

func init() {
	initializers.LoadEnvVairables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Bank{})
	// modle struct to create - why & in front

}
