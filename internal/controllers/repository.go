package controllers

import (
	"github-api/pkg/auth"
	"github-api/pkg/models"
	"github-api/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreateRepo handles the creation of a new repository.
// It expects a JSON payload to initialize a repository model and a token parameter
// for authentication.
// The function validates the repository model and token, and if valid,
// it creates a new repository using the provided token.
//
// Responses:
//   - 400 Bad Request: If the repository model is invalid or cannot be created.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 500 Internal Server Error: If an error occurs while creating the repository.
//   - 201 Created: If the repository is successfully created.
func CreateRepo(c *gin.Context) {
	var repo, err = models.NewRepository(c)
	if (repo == models.Repository{}) || err != nil {
		response.StatusBadRequest(c)
		return
	}
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.StatusUnauthorized(c)
		return
	}
	if err := repo.CreateNew(client); err != nil {
		response.StatusInternalServerError(c, err)
		return
	}
	response.StatusCreated(c, repo)
}

// DeleteRepo handles the deletion of a repository.
// It expects a JSON payload to initialize a repository model and a token parameter
// for authentication.
//
// The function validates the repository model and token, and if valid,
// it deletes the specified repository using the provided token.
//
// The function also checks if the repository exists before attempting to delete it.
// Responses:
//   - 400 Bad Request: If the repository model is invalid or cannot be deleted.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 500 Internal Server Error: If an error occurs while deleting the repository.
//   - 204 No Content: If the repository is successfully deleted.
//   - 404 Not Found: If the repository does not exist.
func DeleteRepo(c *gin.Context) {
	var repo, err = models.NewRepository(c)
	if (repo == models.Repository{}) || err != nil {
		response.StatusBadRequest(c)
		return
	}
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.StatusUnauthorized(c)
		return
	}
	if err := repo.DeleteRepo(client); err != nil {
		response.StatusInternalServerError(c, err)
	}
}

// ListRepos handles the retrieval of repositories.
// TODO: Implement the logic to list repositories.
func ListRepos(c *gin.Context) {

}

// PullRequests handles the retrieval of pull requests for a repository.
// TODO: Implement the logic to list pull requests.
func PullRequests(c *gin.Context) {

}
