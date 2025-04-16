package mocks

import (
	"context"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/mock"
)

// MockGitHubClient is a mock implementation of the GitHubClient.
type MockGitHubClient struct {
	mock.Mock
	*github.Client
}

func (m *MockGitHubClient) GetUser(ctx context.Context, user string) (*github.User, *github.Response, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*github.User), args.Get(1).(*github.Response), args.Error(2)
}
func (m *MockGitHubClient) GetRepositories(ctx context.Context, owner string, repo string) (*github.Repository, *github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).(*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}
func (m *MockGitHubClient) CreateRepository(ctx context.Context, org string, repo *github.Repository) (*github.Repository, *github.Response, error) {
	args := m.Called(ctx, org, repo)
	return args.Get(0).(*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}
func (m *MockGitHubClient) DeleteRepository(ctx context.Context, owner string, repo string) (*github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).(*github.Response), args.Error(1)
}
func (m *MockGitHubClient) ListPullRequests(ctx context.Context, owner string, repo string, opt *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	args := m.Called(ctx, owner, repo)
	return args.Get(0).([]*github.PullRequest), args.Get(1).(*github.Response), args.Error(2)
}
func (m *MockGitHubClient) ListRepos(ctx context.Context, owner string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error) {
	args := m.Called(ctx, owner, opt)
	return args.Get(0).([]*github.Repository), args.Get(1).(*github.Response), args.Error(2)
}
