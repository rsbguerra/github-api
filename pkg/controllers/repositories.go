package controllers

import (
	"github-api/pkg/auth"
	"github-api/pkg/models"
	"github-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
)

// CreateRepo handles the creation of a new repository.
// It expects the following parameters:
//   - token: The authentication token for the GitHub API.
//
// The function validates the repository model and token, and if valid,
// it creates a new repository using the provided token.
//
// Responses:
//   - 201 Created: If the repository is successfully created.
//   - 400 Bad Request: If the repository model is invalid or cannot be created.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 500 Internal Server Error: If an error occurs while creating the repository.
func CreateRepo(c *gin.Context) {
	// Extract the token from the request parameters
	var token = c.Param("token")
	if token == "" {
		response.StatusBadRequestMissingParams(c, []string{token})
		return
	}

	var repo, err = models.ConvertFromContext(c)
	if (repo == models.RepositoryModel{}) || err != nil {
		response.HandleGithubErrors(c, err)
		return
	}
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.HandleGithubErrors(c, err)
		return
	}
	if err := repo.CreateNew(client); err != nil {
		response.HandleGithubErrors(c, err)
		return
	}
	response.StatusCreated(c, repo)
}

// DeleteRepo handles the deletion of a repository.
// It expects the following parameters:
//   - token: The authentication token for the GitHub API.
//   - repoName: The name of the repository to be deleted.
//
// The function validates the repository model and token, and if valid,
// it deletes the specified repository using the provided token.
//
// The function also checks if the repository exists before attempting to delete it.
// Responses:
//   - 204 No Content: If the repository is successfully deleted.
//   - 400 Bad Request: If the repository model is invalid, cannot be deleted or if there are no request parameters.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 404 Not Found: If the repository does not exist.
//   - 500 Internal Server Error: If an error occurs while deleting the repository.
func DeleteRepo(c *gin.Context) {
	// Extract the token from the request parameters
	var token = c.Param("token")
	if token == "" {
		response.StatusBadRequestMissingParams(c, []string{token})
		return
	}

	var repo, err = models.ConvertFromContext(c)

	// Check if the repository model is valid
	if (repo == models.RepositoryModel{}) || err != nil {
		response.StatusBadRequest(c)
		return
	}

	// Check if the token is valid
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.HandleGithubErrors(c, err)
		return
	}

	// Check if the repository exists
	if err := repo.DeleteRepo(client); err != nil {
		response.HandleGithubErrors(c, err)
	}

	// If the repository was deleted successfully, return a 204 No Content response
	response.StatusNoContent(c)
}

// PullRequests handles the retrieval of open pull requests for a repository.
// It expects the following parameters:
//   - token: The authentication token for the GitHub API.
//   - username: The GitHub username who owns the repository.
//   - repoName: The name of the repository whose pull requests are to be listed.
//
// The function authenticates the user using the provided token and retrieves
// the list of open pull requests for the specified repository, sorted by
// the creation date in descending order.
//
// Responses:
//   - 200 OK: If the pull requests are successfully retrieved.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 403 Forbidden: If the user does not have permission to access the repository.
//   - 404 Not Found: If the specified repository does not exist.
//   - 500 Internal Server Error: If an error occurs while retrieving the pull requests.
func PullRequests(c *gin.Context) {

	// Extract the token and parameters from the request
	params := map[string]string{
		"token":    c.Param("token"),
		"username": c.Param("username"),
		"repoName": c.Param("repoName"),
	}

	var missingParams []string
	for key, value := range params {
		if value == "" {
			missingParams = append(missingParams, key)
		}
	}
	if len(missingParams) > 0 {
		response.StatusBadRequestMissingParams(c, missingParams)
		return
	}

	// Check if the token is valid
	client, err := auth.GetClient(params["token"])
	if err != nil {
		response.HandleGithubErrors(c, err)
		return
	}

	opt := &github.PullRequestListOptions{State: "open", Sort: "created", Direction: "desc"}
	pullRequests, _, err := client.PullRequests.List(c, params["username"], params["repoName"], opt)
	if err != nil {
		response.HandleGithubErrors(c, err)
	}

	// 200 OK: if the pull requests are successfully retrieved
	response.StatusOK(c, pullRequests)
}

// Index handles the root endpoint of the API.
// It returns a 200 OK response with no content.
func Index(c *gin.Context) {
	response.StatusOK(c, nil)
}
