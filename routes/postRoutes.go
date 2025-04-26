package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hp1234v/go-crud/controllers"
)

func PostRoutes(r *gin.Engine) {
	r.POST("/createPost", controllers.CreatePost )
	r.GET("/getAllPost", controllers.GetAllPosts )
	r.GET("/getPost/:id", controllers.GetPost )
	r.PUT("/updatePost/:id", controllers.UpdatePost )
	r.DELETE("/deletePost/:id", controllers.DeletePost )
}