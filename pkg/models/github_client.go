package models

import (
	"context"
	"github.com/google/go-github/v50/github"
)

// GitHubClientWrapper is a wrapper around the GitHub client to provide
// convenient methods for interacting with the GitHub API.
type GitHubClientWrapper struct {
	Client *github.Client
}

// GetUser retrieves a GitHub user by their username.
// Parameters:
// - ctx: The context for the request.
// - user: The username of the GitHub user to retrieve.
// Returns:
// - A pointer to the GitHub user object.
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) GetUser(ctx context.Context, user string) (*github.User, *github.Response, error) {
	return w.Client.Users.Get(ctx, user)
}

// GetRepositories retrieves a specific repository by owner and name.
// Parameters:
// - ctx: The context for the request.
// - owner: The owner of the repository.
// - repo: The name of the repository.
// Returns:
// - A pointer to the GitHub repository object.
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) GetRepositories(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	return w.Client.Repositories.Get(ctx, owner, repo)
}

// CreateRepository creates a new repository under the specified organization.
// Parameters:
// - ctx: The context for the request.
// - org: The organization under which the repository will be created.
// - repo: A pointer to the repository object to be created.
// Returns:
// - A pointer to the created GitHub repository object.
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) CreateRepository(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error) {
	return w.Client.Repositories.Create(ctx, org, repo)
}

// DeleteRepository deletes a repository by owner and name.
// Parameters:
// - ctx: The context for the request.
// - owner: The owner of the repository.
// - repo: The name of the repository.
// Returns:
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) DeleteRepository(ctx context.Context, owner, repo string) (*github.Response, error) {
	return w.Client.Repositories.Delete(ctx, owner, repo)
}

// ListPullRequests lists all pull requests for a specific repository.
// Parameters:
// - ctx: The context for the request.
// - owner: The owner of the repository.
// - repo: The name of the repository.
// Returns:
// - A slice of pointers to GitHub pull request objects.
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) ListPullRequests(ctx context.Context, owner, repo string, opt *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	return w.Client.PullRequests.List(ctx, owner, repo, opt)
}

// ListRepos lists all repositories for a specific owner.
// Parameters:
// - ctx: The context for the request.
// - owner: The owner whose repositories will be listed.
// - opt: Options for listing repositories.
// Returns:
// - A slice of pointers to GitHub repository objects.
// - A pointer to the GitHub response object.
// - An error, if any occurred.
func (w *GitHubClientWrapper) ListRepos(ctx context.Context, owner string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	return w.Client.Repositories.List(ctx, owner, opt)
}
