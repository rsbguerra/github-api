package controllers

import (
	"bytes"
	"errors"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepositoryModel is a mock implementation of the RepositoryModel.
type MockRepositoryModel struct {
	mock.Mock
}

func (m *MockRepositoryModel) CreateNew(client interface{}) error {
	args := m.Called(client)
	return args.Error(0)
}

func GenerateRandomRepoName() string {
	words := []string{"alpha", "beta", "gamma", "delta", "omega", "nova", "lunar", "solar", "cosmic", "stellar"}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return strings.Join([]string{words[rand.Intn(len(words))], words[rand.Intn(len(words))]}, "-")
}

// MockAuth is a mock implementation of the auth package.
var MockAuth = struct {
	GetClient func(token string) (interface{}, error)
}{
	GetClient: func(token string) (interface{}, error) {
		return nil, nil
	},
}

func TestCreateRepo(t *testing.T) {
	validToken := os.Getenv("TEST_AUTH_TOKEN")

	// Set up Gin context and recorder
	gin.SetMode(gin.TestMode)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing or return a 204 No Content
		w.WriteHeader(http.StatusNoContent)
	})
	router := gin.Default()
	router.POST("/repositories/:token", CreateRepo)

	// Mock dependencies
	mockRepo := new(MockRepositoryModel)
	mockRepo.On("CreateNew", mock.Anything).Return(nil)

	MockAuth.GetClient = func(token string) (interface{}, error) {
		if token == validToken {
			return &struct{}{}, nil
		}
		return nil, errors.New("unauthorized")
	}

	repoName := GenerateRandomRepoName()

	// Test cases
	tests := []struct {
		name           string
		token          string
		requestBody    string
		expectedStatus int
		mockError      error
	}{
		{
			name:           "Successful repository creation",
			token:          validToken,
			requestBody:    `{"name": test-repo", "private": false}`,
			expectedStatus: http.StatusCreated,
			mockError:      nil,
		},
		{
			name:           "Invalid token",
			token:          "invalid-token",
			requestBody:    `{"name": "` + repoName + `", "private": false}`,
			expectedStatus: http.StatusUnauthorized,
			mockError:      nil,
		},
		{
			name:           "Empty token",
			token:          "",
			requestBody:    `{"name": "` + repoName + `", "private": false}`,
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
		},
		{
			name:           "Invalid request body",
			token:          validToken,
			requestBody:    `{"invalid": "data"}`,
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
		},
		{
			name:           "Repository already exists",
			token:          validToken,
			requestBody:    `{"name": "` + "github-api" + `", "private": false}`,
			expectedStatus: http.StatusConflict,
			mockError:      nil,
		},
		// Note: Test fails because the mock error is not being returned correctly
		{
			name:           "Error creating repository",
			token:          validToken,
			requestBody:    `{"name": "` + repoName + `", "private": false}`,
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("creation error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock behavior
			mockRepo.On("CreateNew", mock.Anything).Return(tt.mockError)

			// Create request and response recorder
			req, _ := http.NewRequest(http.MethodPost, "/repositories/"+tt.token, bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rec, req)

			// Assert the response status
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
func TestDeleteRepo(t *testing.T) {
	validToken := os.Getenv("TEST_AUTH_TOKEN")
	// Set up Gin context and recorder
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/repositories/:token", DeleteRepo)

	// Mock dependencies
	mockRepo := new(MockRepositoryModel)
	mockRepo.On("DeleteRepo", mock.Anything).Return(nil)

	_, err := http.NewRequest(http.MethodPost, "/repositories/"+validToken, bytes.NewBufferString(`{"name": "test-repo"}`))
	if err != nil {
		return
	}

	MockAuth.GetClient = func(token string) (interface{}, error) {
		if token == validToken {
			return &struct{}{}, nil
		}
		return nil, errors.New("unauthorized")
	}

	// Test cases
	tests := []struct {
		name           string
		token          string
		requestBody    string
		expectedStatus int
		mockError      error
	}{
		{
			name:           "Successful repository deletion",
			token:          validToken,
			requestBody:    `{"name": "test-repo"}`,
			expectedStatus: http.StatusNoContent,
			mockError:      nil,
		},
		{
			name:           "Invalid token",
			token:          "invalid-token",
			requestBody:    `{"name": "test-repo"}`,
			expectedStatus: http.StatusUnauthorized,
			mockError:      nil,
		},
		{
			name:           "Repository not found",
			token:          validToken,
			requestBody:    `{"name": "nonexistent-repo"}`,
			expectedStatus: http.StatusNotFound,
			mockError:      errors.New("repository not found"),
		},
		// Note: Test fails because the mock error is not being returned correctly
		{
			name:           "Error deleting repository",
			token:          validToken,
			requestBody:    `{"name": "test-repo"}`,
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("deletion error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock behavior
			mockRepo.On("DeleteRepo", mock.Anything).Return(tt.mockError)

			// Create request and response recorder
			req, _ := http.NewRequest(http.MethodDelete, "/repositories/"+tt.token, bytes.NewBufferString(tt.requestBody))
			rec := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rec, req)

			// Assert the response status
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestListRepos(t *testing.T) {
	validToken := os.Getenv("TEST_AUTH_TOKEN")

	// Set up Gin context and recorder
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/repositories/:token", ListRepos)

	// Mock dependencies
	MockAuth.GetClient = func(token string) (interface{}, error) {
		if token == validToken {
			return &struct{}{}, nil
		}
		return nil, errors.New("unauthorized")
	}

	// Test cases
	tests := []struct {
		name           string
		token          string
		expectedStatus int
		mockError      error
	}{
		{
			name:           "Successful repository retrieval",
			token:          validToken,
			expectedStatus: http.StatusOK,
			mockError:      nil,
		},
		{
			name:           "Invalid token",
			token:          "invalid-token",
			expectedStatus: http.StatusUnauthorized,
			mockError:      nil,
		},
		// Note: Test fails because the mock error is not being returned correctly
		{
			name:           "Internal server error",
			token:          validToken,
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("internal error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock behavior for ListRepos
			MockAuth.GetClient = func(token string) (interface{}, error) {
				if token == validToken {
					return &struct{}{}, nil
				}
				return nil, errors.New("unauthorized")
			}

			// Create request and response recorder
			req, _ := http.NewRequest(http.MethodGet, "/repositories/"+tt.token, bytes.NewBufferString(""))
			rec := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rec, req)

			// Assert the response status
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestListPullRequests(t *testing.T) {
	validToken := os.Getenv("TEST_AUTH_TOKEN")
	// Set up Gin context and recorder
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/pull-requests/:username/:repoName/:token", PullRequests)

	// Mock dependencies
	MockAuth.GetClient = func(token string) (interface{}, error) {
		if token == validToken {
			return &struct{}{}, nil
		}
		return nil, errors.New("unauthorized")
	}

	// Test cases
	tests := []struct {
		name           string
		token          string
		username       string
		repoName       string
		expectedStatus int
		mockError      error
		mockPRs        []string
	}{
		{
			name:           "Successful pull request retrieval",
			token:          validToken,
			username:       "rsbguerra",
			repoName:       "github-api",
			expectedStatus: http.StatusOK,
			mockError:      nil,
			mockPRs:        []string{"PR1", "PR2"},
		},
		{
			name:           "Invalid token",
			token:          "invalid-token",
			username:       "test-user",
			repoName:       "test-repo",
			expectedStatus: http.StatusUnauthorized,
			mockError:      nil,
			mockPRs:        nil,
		},
		// Note: Test fails because the mock error is not being returned correctly
		{
			name:           "Internal server error",
			token:          validToken,
			username:       "test-user",
			repoName:       "test-repo",
			expectedStatus: http.StatusInternalServerError,
			mockError:      errors.New("internal error"),
			mockPRs:        nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock behavior for PullRequests
			MockAuth.GetClient = func(token string) (interface{}, error) {
				if token == validToken {
					return &struct{}{}, nil
				}
				return nil, errors.New("unauthorized")
			}

			// Create request and response recorder
			req, _ := http.NewRequest(http.MethodGet, "/pull-requests/"+tt.username+"/"+tt.repoName+"/"+tt.token, nil)
			rec := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rec, req)

			// Assert the response status
			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
