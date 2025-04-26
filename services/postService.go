package services

import (
	"github.com/hp1234v/go-crud/models"
	"github.com/hp1234v/go-crud/repository"
)

// Define a custom type for CreatePostService result
type CreatePostResult struct {
	Post models.Post
	Err  error
}

func CreatePostService(title string, body string) <-chan CreatePostResult {
	resultChan := make(chan CreatePostResult)

	go func() {
		post, err := repository.CreatePostRepoService(title, body)
		resultChan <- CreatePostResult{Post: post, Err: err}
	}()

	return resultChan
}

// Define a custom type for GetAllPostsService result
type GetAllPostsResult struct {
	Posts []models.Post
	Err   error
}

func GetAllPostsService() <-chan GetAllPostsResult {
	resultChan := make(chan GetAllPostsResult)

	go func() {
		posts, err := repository.GetAllPostsRepoService()
		resultChan <- GetAllPostsResult{Posts: posts, Err: err}
	}()

	return resultChan
}

// Define a custom type for GetPostService result
type GetPostResult struct {
	Post models.Post
	Err  error
}

func GetPostService(id string) <-chan GetPostResult {
	resultChan := make(chan GetPostResult)

	go func() {
		post, err := repository.GetPostRepoService(id)
		resultChan <- GetPostResult{Post: post, Err: err}
	}()

	return resultChan
}

// Define a custom type for UpdatePostService result
type UpdatePostResult struct {
	Post models.Post
	Err  error
}

func UpdatePostService(id string, title string, body string) <-chan UpdatePostResult {
	resultChan := make(chan UpdatePostResult)

	go func() {
		post, err := repository.UpdatePostRepoService(id, title, body)
		resultChan <- UpdatePostResult{Post: post, Err: err}
	}()

	return resultChan
}

// Define a custom type for DeletePostService result
type DeletePostResult struct {
	Err error
}

func DeletePostService(id string) <-chan DeletePostResult {
	resultChan := make(chan DeletePostResult)

	go func() {
		err := repository.DeletePostRepoService(id)
		resultChan <- DeletePostResult{Err: err}
	}()

	return resultChan
}