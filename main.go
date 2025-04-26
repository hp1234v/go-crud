package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/hp1234v/go-crud/initializers"
	"github.com/hp1234v/go-crud/models"
	"github.com/hp1234v/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	mode := os.Getenv("GIN_MODE")
	if mode == "" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	initializers.DB.AutoMigrate(&models.Post{})

	r := gin.Default()
	routes.PostRoutes(r)
	r.Run()
}