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
// It is used to simulate the behavior of the repository model in tests.
type MockRepositoryModel struct {
	mock.Mock
}

// CreateNew simulates the creation of a new repository.
// It accepts a client interface and returns an error if the operation fails.
func (m *MockRepositoryModel) CreateNew(client interface{}) error {
	args := m.Called(client)
	return args.Error(0)
}

// GenerateRandomRepoName generates a random repository name by combining two random words.
// The words are selected from a predefined list.
func GenerateRandomRepoName() string {
	words := []string{"alpha", "beta", "gamma", "delta", "omega", "nova", "lunar", "solar", "cosmic", "stellar"}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	return strings.Join([]string{words[rand.Intn(len(words))], words[rand.Intn(len(words))]}, "-")
}

// MockAuth is a mock implementation of the auth package.
// It provides a mock GetClient function to simulate authentication behavior.
var MockAuth = struct {
	GetClient func(token string) (interface{}, error)
}{
	GetClient: func(token string) (interface{}, error) {
		return nil, nil
	},
}

// TestCreateRepo tests the CreateRepo function for various scenarios, including success and failure cases.
func TestCreateRepo(t *testing.T) {
	validToken := os.Getenv("AUTH_TOKEN")

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

	// Test cases for CreateRepo
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
