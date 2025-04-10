package models

import (
    "context"
    "github.com/gin-gonic/gin"
    "github.com/go-git/go-git/v5"
    "github.com/google/go-github/v50/github"
    "os"
)

// Repository represents a GitHub repository with its basic details.
// It includes the repository's name, URL, owner, and privacy status.
type Repository struct {
    Name    string `json:"name"`
    URL     string `json:"url"`
    Owner   string `json:"owner"`
    Private bool   `json:"private"`
}

// NewRepository creates a new Repository instance by binding JSON data from the provided
// gin.Context. It parses the incoming JSON payload into a Repository struct.
//
// Parameters:
//   - c: The gin.Context containing the JSON payload to be bound.
//
// Returns:
//   - Repository: The populated Repository struct.
//   - error: An error if the JSON binding fails, otherwise nil.
func NewRepository(c *gin.Context) (Repository, error) {
    var repo Repository

    if err := c.ShouldBindJSON(&repo); err != nil {
        return Repository{}, err
    }
    return repo, nil
}

// CloneRepo clones the repository specified by the Repository struct to a local directory.
// It uses the repository's Name as the target directory and the URL as the source.
// The cloning process outputs progress to the standard output.
//
// Returns:
//   - error: An error if the cloning process fails, otherwise nil.
func (r *Repository) CloneRepo() error {
    _, err := git.PlainClone(r.Name, false, &git.CloneOptions{
        URL:      r.URL,
        Progress: os.Stdout,
    })
    return err
}

// CreateNew creates a new GitHub repository using the provided GitHub client.
// It initializes a repository object with the current Repository struct's
// Name and Private fields, and then sends a request to create the repository
// in the authenticated user's account.
//
// Parameters:
//   - client: A GitHub client instance used to interact with the GitHub API.
//
// Returns:
//   - An error if the repository creation fails, otherwise nil.
func (r *Repository) CreateNew(client *github.Client) error {

    repo := &github.Repository{
        Name:    github.String(r.Name),
        Private: github.Bool(r.Private),
    }

    _, _, err := client.Repositories.Create(context.Background(), "", repo)
    return err
}

// DeleteRepo deletes the repository associated with the Repository struct
// using the provided GitHub client. It sends a request to the GitHub API
// to delete the repository identified by its owner and name.
//
// Parameters:
//   - client: A pointer to a github.Client instance used to interact with
//             the GitHub API.
//
// Returns:
//   - error: An error if the deletion fails, or nil if the operation is
//            successful.
func (r *Repository) DeleteRepo(client *github.Client) error {
    _, err := client.Repositories.Delete(context.Background(), r.Owner, r.Name)
    return err
}

//func (r *Repository) ListPullRequests(client *github.Client) (*github.Repository, error) {
// TODO
//}
