package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hp1234v/go-crud/initializers"
	"github.com/hp1234v/go-crud/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r :=  gin.Default()
	routes.PostRoutes(r)
	r.Run()
}