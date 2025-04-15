package interfaces

import (
	"context"
	"github.com/google/go-github/v50/github"
)

// GitHubClient is an interface that defines methods for interacting with the GitHub API.
type GitHubClient interface {

	// GetUser retrieves a GitHub user by their username.
	// Parameters:
	// - ctx: The context for the request.
	// - user: The username of the GitHub user.
	// Returns:
	// - A pointer to the retrieved GitHub user.
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	GetUser(ctx context.Context, user string) (*github.User, *github.Response, error)

	// GetRepositories retrieves a specific repository by owner and name.
	// Parameters:
	// - ctx: The context for the request.
	// - owner: The owner of the repository.
	// - repo: The name of the repository.
	// Returns:
	// - A pointer to the retrieved repository.
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	GetRepositories(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)

	// CreateRepository creates a new repository under the specified organization.
	// Parameters:
	// - ctx: The context for the request.
	// - org: The organization under which the repository will be created.
	// - repo: A pointer to the repository object to be created.
	// Returns:
	// - A pointer to the created repository.
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	CreateRepository(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error)

	// DeleteRepository deletes a repository by owner and name.
	// Parameters:
	// - ctx: The context for the request.
	// - owner: The owner of the repository.
	// - repo: The name of the repository.
	// Returns:
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	DeleteRepository(ctx context.Context, owner, repo string) (*github.Response, error)

	// ListRepos lists repositories for a specified owner.
	// Parameters:
	// - ctx: The context for the request.
	// - owner: The owner whose repositories will be listed.
	// - opt: Options for listing repositories.
	// Returns:
	// - A slice of pointers to the listed repositories.
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	ListRepos(ctx context.Context, owner string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)

	// ListPullRequests lists pull requests for a specified repository.
	// Parameters:
	// - ctx: The context for the request.
	// - owner: The owner of the repository.
	// - repo: The name of the repository.
	// - opt: Options for listing pull requests.
	// Returns:
	// - A slice of pointers to the listed pull requests.
	// - A pointer to the GitHub API response.
	// - An error, if any occurred.
	ListPullRequests(ctx context.Context, owner, repo string, opt *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error)
}
