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
//   - 404 Not Found: If the specified user does not exist or has no repositories.
//   - 409 Conflict: If the repository already exists.
//   - 500 Internal Server Error: If an error occurs while creating the repository.
func CreateRepo(c *gin.Context) {
	// Extract the token from the request parameters
	var token = c.Param("token")
	if token == "" {
		// Response: 400 Bad Request if the token is missing
		response.StatusBadRequestMissingParams(c, []string{token})
		return
	}

	var repo, err = models.ConvertFromContext(c)
	if (repo == models.RepositoryModel{}) || err != nil {
		// Response: 400 Bad Request if the repository model is invalid
		response.StatusBadRequest(c)
		return
	}
	// Check if the token is valid
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		// Response: 401 Unauthorized if the token is invalid
		response.HandleGithubErrors(c, err)
		return
	}

	// Check if the repository model already exists
	exists, err := repo.RepoExists(client)
	if exists || err != nil {
		// Response: 409 Conflict if the repository already exists
		response.StatusConflict(c)
		return
	}

	if err := repo.CreateNew(client); err != nil {
		// Response: 403 Forbidden if the user does not have permission to create the repository
		response.StatusForbidden(c)
		return
	}
	// Response: 201 Created if the repository is successfully created
	response.StatusCreated(c, repo)
}

// DeleteRepo handles the deletion of a repository.
// It expects the following parameters:
//   - token: The authentication token for the GitHub API.
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
		// Response: 400 Bad Request if the token is missing
		response.StatusBadRequestMissingParams(c, []string{token})
		return
	}

	var repo, err = models.ConvertFromContext(c)

	// Check if the repository model is valid
	if (repo == models.RepositoryModel{}) || err != nil {
		// Response: 400 Bad Request if the repository model is invalid
		response.StatusBadRequest(c)
		return
	}

	// Check if the token is valid
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		// Response: 401 Unauthorized if the token is invalid
		response.HandleGithubErrors(c, err)
		return
	}

	// Check if the repository exists
	exists, err := repo.RepoExists(client)
	if exists || err != nil {
		// Response: 404 Not Found if the repository does not exist
		response.HandleGithubErrors(c, err)
	}

	// Delete the repository
	if err := repo.DeleteRepo(client); err != nil {
		// Response: 500 Internal Server Error if an error occurs while deleting the repository
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

	// Check if the token is valid
	var missingParams []string
	for key, value := range params {
		if value == "" {
			missingParams = append(missingParams, key)
		}
	}
	if len(missingParams) > 0 {
		// Response: 400 Bad Request if the token or parameters are missing
		response.StatusBadRequestMissingParams(c, missingParams)
		return
	}

	// Check if the token is valid
	client, err := auth.GetClient(params["token"])
	if err != nil {
		// Response: 401 Unauthorized if the token is invalid
		response.HandleGithubErrors(c, err)
		return
	}

	// Check if repository exists
	_, _, err = client.GetRepositories(c, params["username"], params["repoName"])
	if err != nil {
		// Response: 404 Not Found if the repository does not exist
		response.HandleGithubErrors(c, err)
		return
	}

	opt := &github.PullRequestListOptions{State: "open", Sort: "created", Direction: "desc"}
	pullRequests, _, err := client.ListPullRequests(c, params["username"], params["repoName"], opt)
	if err != nil {
		// Response: 403 Forbidden if the user does not have permission to access the repository
		response.StatusForbidden(c)
	}

	// 200 OK: if the pull requests are successfully retrieved
	response.StatusOK(c, pullRequests)
}

// ListRepos handles the retrieval of repositories for a user.
// It expects the following parameters:
//   - token: The authentication token for the GitHub API.
//   - username: The GitHub username whose repositories are to be listed.
//
// The function authenticates the user using the provided token and retrieves
// the list of repositories owned by the specified username, sorted by the
// last updated time in descending order.
//
// Responses:
//   - 200 OK: If the repositories are successfully retrieved.
//   - 401 Unauthorized: If the provided token is invalid or authentication fails.
//   - 500 Internal Server Error: If an error occurs while retrieving the repositories.
func ListRepos(c *gin.Context) {
	token := c.Param("token")
	client, err := auth.GetClient(token)
	// Check if the token is valid
	if err != nil {
		// Response: 401 Unauthorized if the token is invalid
		response.StatusUnauthorized(c)
		return
	}
	// Check if the user is authenticated
	user, _, err := client.GetUser(c, "")
	if err != nil {
		// Response: 401 an error occoured while retrieving the user
		response.StatusUnauthorized(c)
		return
	}

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.ListRepos(c, *user.Login, opt)

	// Check if the request to list repositories was successful
	if err != nil {
		// Response: 500 Internal Server Error if an error occurs while retrieving the repositories
		response.StatusInternalServerError(c, err)
		return
	}
	// Response: 200 OK if the repositories are successfully retrieved
	response.StatusOK(c, repos)
}

// Index handles the root endpoint of the API.
// It returns a 200 OK response with no content.
func Index(c *gin.Context) {
	response.StatusOK(c, nil)
}
