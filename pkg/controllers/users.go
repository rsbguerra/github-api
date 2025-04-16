package controllers

import (
	"github-api/pkg/auth"
	"github-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
)

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
	//token, username := c.Param("token"), c.Param("username")

	// Extract the token and parameters from the request
	params := map[string]string{
		"token":    c.Param("token"),
		"username": c.Param("username"),
	}

	client, err := auth.GetClient(params["token"])

	// Check if the token is valid
	if err != nil {
		response.StatusUnauthorized(c)
		return
	}

	opt := &github.RepositoryListOptions{Type: "owner", Sort: "updated", Direction: "desc"}
	repos, _, err := client.Repositories.List(c, params["username"], opt)

	// Check if the request to list repositories was successful
	if err != nil {
		response.StatusInternalServerError(c, err)
		return
	}

	response.StatusOK(c, repos)
}
