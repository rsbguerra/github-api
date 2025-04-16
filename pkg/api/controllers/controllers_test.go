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

	println(validToken)

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
			requestBody:    `{"name": "` + repoName + `", "private": false}`,
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
			name:           "Invalid request body",
			token:          validToken,
			requestBody:    `{"invalid": "data"}`,
			expectedStatus: http.StatusBadRequest,
			mockError:      nil,
		},
		{
			name:           "Repository already exists",
			token:          validToken,
			requestBody:    `{"name": "` + repoName + `", "private": false}`,
			expectedStatus: http.StatusConflict,
			mockError:      nil,
		},
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

}
func TestListRepos(t *testing.T)        {}
func TestListPullRequests(t *testing.T) {}
