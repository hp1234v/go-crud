package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hp1234v/go-crud/services"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	result := <-services.CreatePostService(body.Title, body.Body) // 🚀 waiting for channel

	if result.Err != nil {
		c.JSON(500, gin.H{"error": "Could not create post"})
		return
	}

	c.JSON(200, gin.H{
		"post": result.Post,
	})
}

func UpdatePost(c *gin.Context) {
	idParam := c.Param("id")

	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	result := <-services.UpdatePostService(idParam, body.Title, body.Body) // 🚀

	if result.Err != nil {
		if result.Err.Error() == "WHERE conditions required" {
			// Handle the specific error case
			c.JSON(404, gin.H{"error": "Post is not present"})
		} else {
			// Handle other errors
			c.JSON(500, gin.H{"error": "Something went wrong"})
		}
		return
	}

	c.JSON(200, gin.H{
		"post": result.Post,
	})
}

func GetAllPosts(c *gin.Context) {
	result := <-services.GetAllPostsService() // 🚀

	if result.Err != nil {
		c.JSON(500, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, gin.H{
		"posts": result.Posts,
	})
}

func GetPost(c *gin.Context) {
	idParam := c.Param("id")

	result := <-services.GetPostService(idParam) // 🚀

	if result.Err != nil {
		c.JSON(500, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, gin.H{
		"post": result.Post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")

	result := <-services.DeletePostService(id) // 🚀

	if result.Err != nil {
		c.JSON(500, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleted Successfully",
	})
}