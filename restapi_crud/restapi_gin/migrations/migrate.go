package main

import (
	"yashprakash13/Go-Basics/initializers"
	"yashprakash13/Go-Basics/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
