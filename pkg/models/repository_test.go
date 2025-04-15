package models

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGitHubClient is a mock implementation of the GitHubClient.
type MockGitHubClient struct {
	mock.Mock
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

func TestConvertFromContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	body := `{"name": "test-repo", "url": "https://github.com/test/test-repo", "private": true}`
	c.Request = httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
	c.Request.Header.Set("Content-Type", "application/json")

	repo, err := ConvertFromContext(c)
	assert.NoError(t, err)
	assert.Equal(t, "test-repo", *repo.Name)
	assert.Equal(t, "https://github.com/test/test-repo", *repo.URL)
	assert.True(t, *repo.Private)
}

func TestRepoExists(t *testing.T) {
	mockClient := new(MockGitHubClient)
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo")}}

	mockClient.On("GetUser", mock.Anything, "").Return(
		&github.User{Login: github.String("test-user")},
		&github.Response{},
		nil)
	mockClient.On("GetRepositories", mock.Anything, "test-user", "test-repo").Return(
		&github.Repository{},
		&github.Response{},
		nil)

	exists, err := repo.RepoExists(mockClient)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestCloneRepo(t *testing.T) {
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo"), URL: github.String("https://github.com/test/test-repo")}}

	err := repo.CloneRepo()
	assert.Error(t, err) // Expecting an error since the URL is not accessible in the test environment
}

func TestCreateNew(t *testing.T) {
	mockClient := new(MockGitHubClient)
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo"), Private: github.Bool(true)}}

	mockClient.On("CreateRepository", mock.Anything, "", mock.Anything).Return(
		&github.Repository{},
		&github.Response{},
		nil)

	err := repo.CreateNew(mockClient)
	assert.NoError(t, err)
}

func TestDeleteRepo(t *testing.T) {
	mockClient := new(MockGitHubClient)
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo")}}

	mockClient.On("GetUser", mock.Anything, "").Return(
		&github.User{Login: github.String("test-user")},
		&github.Response{},
		nil)
	mockClient.On("DeleteRepository", mock.Anything, "test-user", "test-repo").Return(
		&github.Response{},
		nil)
	err := repo.DeleteRepo(mockClient)
	assert.NoError(t, err)
}
