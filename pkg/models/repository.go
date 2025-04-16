package models

import (
	"context"
	"github-api/pkg/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v50/github"
	"os"
)

// RepositoryModel represents a GitHub repository with its basic details.
// It includes the repository's name, URL, owner, and privacy status.
type RepositoryModel struct {
	*github.Repository
}

// ConvertFromContext creates a new RepositoryModel instance by binding JSON data from the provided
// gin.Context. It parses the incoming JSON payload into a RepositoryModel struct.
//
// Parameters:
//   - c: The gin.Context containing the JSON payload to be bound.
//
// Returns:
//   - Repository: The populated Repository struct.
//   - error: An error if the JSON binding fails, otherwise nil.
func ConvertFromContext(c *gin.Context) (RepositoryModel, error) {
	var repo RepositoryModel

	if err := c.ShouldBindJSON(&repo); err != nil {
		return RepositoryModel{}, err
	}
	return repo, nil
}

// RepoExists checks if the repository exists on GitHub.
//
// Parameters:
//   - client: A GitHub client instance used to interact with the GitHub API.
//
// Returns:
//   - bool: True if the repository exists, false otherwise.
//   - error: An error if the API request fails, or nil if successful.
func (r *RepositoryModel) RepoExists(client interfaces.GitHubClient) (bool, error) {
	username, _, err := client.GetUser(context.Background(), "")
	if err != nil {
		return false, err
	}

	_, resp, err := client.GetRepositories(context.Background(), *username.Login, *r.Name)
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CloneRepo clones the repository specified by the Repository struct to a local directory.
// It uses the repository's Name as the target directory and the URL as the source.
// The cloning process outputs progress to the standard output.
//
// Returns:
//   - error: An error if the cloning process fails, otherwise nil.
func (r *RepositoryModel) CloneRepo() error {
	_, err := git.PlainClone(*r.Name, false, &git.CloneOptions{
		URL:      *r.URL,
		Progress: os.Stdout,
	})
	return err
}

// CreateNew creates a new GitHub repository using the provided GitHub client.
// It initializes a repository object with the current Repository structs
// Name and Private fields, and then sends a request to create the repository
// in the authenticated user's account.
//
// Parameters:
//   - client: A GitHub client instance used to interact with the GitHub API.
//
// Returns:
//   - An error if the repository creation fails, otherwise nil.
func (r *RepositoryModel) CreateNew(client interfaces.GitHubClient) error {

	repo := &github.Repository{
		Name:    github.String(*r.Name),
		Private: github.Bool(*r.Private),
	}

	_, _, err := client.CreateRepository(context.Background(), "", repo)
	return err
}

// DeleteRepo deletes the repository associated with the Repository struct
// using the provided GitHub client. It sends a request to the GitHub API
// to delete the repository identified by its owner and name.
//
// Parameters:
//   - client: A pointer to a github.Client instance used to interact with
//     the GitHub API.
//
// Returns:
//   - error: An error if the deletion fails, or nil if the operation is
//     successful.
func (r *RepositoryModel) DeleteRepo(client interfaces.GitHubClient) error {
	username, _, err := client.GetUser(context.Background(), "")
	if err != nil {
		return err
	}
	_, err = client.DeleteRepository(context.Background(), *username.Login, *r.Name)
	return err
}
