package repository

import (
	"github.com/hp1234v/go-crud/initializers"
	"github.com/hp1234v/go-crud/models"
)

func CreatePostRepoService(title string, body string) (models.Post, error) {
	post := models.Post{Title: title, Body: body}
	result := initializers.DB.Create(&post)
	return post, result.Error
}

func GetAllPostsRepoService() ([]models.Post, error) {
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	return posts, result.Error
}

func GetPostRepoService(id string) (models.Post, error) {
	var post models.Post
	result := initializers.DB.Find(&post)
	return post, result.Error
}

func UpdatePostRepoService(id string, title string, body string) (models.Post, error) {
	post := models.Post{Title: title, Body: body}
	initializers.DB.First(&post, id)
	result := initializers.DB.Model(&post).Updates(models.Post{
		Title: title,
		Body: body,
	})
	return post, result.Error
}

func DeletePostRepoService(id string) (error) {
	result := initializers.DB.Delete(&models.Post{}, id)
	return result.Error
}