package models

import (
	"bytes"
	"github-api/pkg/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v50/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestConvertFromContext tests the ConvertFromContext function.
// It verifies that the function correctly parses a JSON request body into a RepositoryModel object.
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

// TestRepoExists tests the RepoExists method of the RepositoryModel struct.
// It verifies that the method correctly checks if a repository exists using a mocked GitHub client.
func TestRepoExists(t *testing.T) {
	mockClient := new(mocks.MockGitHubClient)
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

// TestCloneRepo tests the CloneRepo method of the RepositoryModel struct.
// It verifies that the method attempts to clone a repository and expects an error in the test environment.
func TestCloneRepo(t *testing.T) {
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo"), URL: github.String("https://github.com/test/test-repo")}}

	err := repo.CloneRepo()
	assert.Error(t, err) // Expecting an error since the URL is not accessible in the test environment
}

// TestCreateNew tests the CreateNew method of the RepositoryModel struct.
// It verifies that the method successfully creates a new repository using a mocked GitHub client.
func TestCreateNew(t *testing.T) {
	mockClient := new(mocks.MockGitHubClient)
	repo := RepositoryModel{Repository: &github.Repository{Name: github.String("test-repo"), Private: github.Bool(true)}}

	mockClient.On("CreateRepository", mock.Anything, "", mock.Anything).Return(
		&github.Repository{},
		&github.Response{},
		nil)

	err := repo.CreateNew(mockClient)
	assert.NoError(t, err)
}

// TestDeleteRepo tests the DeleteRepo method of the RepositoryModel struct.
// It verifies that the method successfully deletes a repository using a mocked GitHub client.
func TestDeleteRepo(t *testing.T) {
	mockClient := new(mocks.MockGitHubClient)
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
