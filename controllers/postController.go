package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hp1234v/go-crud/services"
)

func CreatePost(c *gin.Context) {
	// Get data off req.body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	// Call the service
	post, err := services.CreatePostService(body.Title, body.Body)
	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	// Send response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Getting the id from url
	idParam := c.Param("id")

	// Get data off req.body
	var body struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, "Invalid request body")
		return
	}

	// Call the service
	post, err := services.UpdatePostService(idParam, body.Title, body.Body)
	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	// Send response
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	// Get the posts
	posts, err := services.GetAllPostsService()

	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Getting the id from url
	idParam := c.Param("id")

	// Get the posts
	post, err := services.GetPostService(idParam)

	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// get the ide off the url
	id := c.Param("id")

	// Get the posts
	err := services.DeletePostService(id)

	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	c.JSON(200, "Deleted Successfully")
}