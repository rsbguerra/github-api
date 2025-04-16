package mocks

import (
	"context"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/mock"
)

// MockGitHubClient is a mock implementation of the GitHubClient interface.
// It is used for testing purposes to simulate the behavior of the GitHub API client.
type MockGitHubClient struct {
	mock.Mock
	*github.Client
}

// GetUser mocks the GetUser method of the GitHub client.
// It retrieves a user by their username.
//
// Parameters:
//   - ctx: The context for the request.
//   - user: The username of the user to retrieve.
//
// Returns:
//   - *github.User: The retrieved user object.
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) GetUser(ctx context.Context, user string) (*github.User, *github.Response, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*github.User), args.Get(1).(*github.Response), args.Error(2)
}

// GetRepositories mocks the GetRepositories method of the GitHub client.
// It retrieves a repository by its owner and name.
//
// Parameters:
//   - ctx: The context for the request.
//   - owner: The owner of the repository.
//   - repo: The name of the repository.
//
// Returns:
//   - *github.Repository: The retrieved repository object.
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) GetRepositories(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).(*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}

// CreateRepository mocks the CreateRepository method of the GitHub client.
// It creates a new repository in the specified organization.
//
// Parameters:
//   - ctx: The context for the request.
//   - org: The organization in which to create the repository.
//   - repo: The repository object to create.
//
// Returns:
//   - *github.Repository: The created repository object.
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) CreateRepository(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error) {
	args := m.Called(ctx, org, repo)
	return args.Get(0).(*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}

// DeleteRepository mocks the DeleteRepository method of the GitHub client.
// It deletes a repository by its owner and name.
//
// Parameters:
//   - ctx: The context for the request.
//   - owner: The owner of the repository.
//   - repo: The name of the repository.
//
// Returns:
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) DeleteRepository(ctx context.Context, owner string, repo string) (*github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).(*github.Response), args.Error(1)
}

// ListPullRequests mocks the ListPullRequests method of the GitHub client.
// It lists pull requests for a repository.
//
// Parameters:
//   - ctx: The context for the request.
//   - owner: The owner of the repository.
//   - repo: The name of the repository.
//   - opt: Options for filtering the pull requests.
//
// Returns:
//   - []*github.PullRequest: A list of pull requests.
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) ListPullRequests(ctx context.Context, owner string, repo string, opt *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).([]*github.PullRequest), args.Get(1).(*github.Response), args.Error(2)
}

// ListRepos mocks the ListRepos method of the GitHub client.
// It lists repositories for a user or organization.
//
// Parameters:
//   - ctx: The context for the request.
//   - owner: The owner of the repositories.
//   - opt: Options for filtering the repositories.
//
// Returns:
//   - []*github.Repository: A list of repositories.
//   - *github.Response: The HTTP response from the GitHub API.
//   - error: An error if the operation fails.
func (m *MockGitHubClient) ListRepos(ctx context.Context, owner string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	args := m.Called(ctx, owner, opt)
	return args.Get(0).([]*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}
