package main

import (
	"github.com/hp1234v/go-crud/initializers"
	"github.com/hp1234v/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}