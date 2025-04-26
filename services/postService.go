package services

import (
	"github.com/hp1234v/go-crud/models"
	"github.com/hp1234v/go-crud/repository"
)

func CreatePostService(title string, body string) (models.Post, error) {
	post, err := repository.CreatePostRepoService(title, body)
	return post, err
}

func GetAllPostsService() ([]models.Post, error) {
	return repository.GetAllPostsRepoService()
}

func GetPostService(id string) (models.Post, error) {
	return repository.GetPostRepoService(id)
}

func UpdatePostService(id string, title string, body string) (models.Post, error) {
	post, err := repository.UpdatePostRepoService(id, title, body)
	return post, err
}

func DeletePostService(id string) (error){
	return repository.DeletePostRepoService(id)
}