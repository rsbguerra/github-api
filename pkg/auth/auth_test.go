package auth

import (
	"context"
	"errors"
	"github-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"net/http"
	"testing"
)

// mockOAuth2TokenSource is a mock implementation of the oauth2.TokenSource interface.
// It is used to simulate the behavior of an OAuth2 token source in tests.
type mockOAuth2TokenSource struct {
	mock.Mock
}

// newOAuth2Client is a function that creates a new HTTP client using the provided OAuth2 token source.
// It can be overridden in tests to mock the behavior of the OAuth2 client.
var newOAuth2Client = func(ctx context.Context, src oauth2.TokenSource) *http.Client {
	return oauth2.NewClient(ctx, src)
}

// Token is a mock implementation of the Token method from the oauth2.TokenSource interface.
// It returns a mocked OAuth2 token and an error, if any.
func (m *mockOAuth2TokenSource) Token() (*oauth2.Token, error) {
	args := m.Called()
	return args.Get(0).(*oauth2.Token), args.Error(1)
}

// TestGetClientWithValidToken tests the GetClient function with a valid token.
// It verifies that the function returns a valid GitHubClientWrapper and no error.
func TestGetClientWithValidToken(t *testing.T) {
	token := "valid-token"
	client, err := GetClient(token)

	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.IsType(t, &models.GitHubClientWrapper{}, client)
}

// TestGetClientWithEmptyToken tests the GetClient function with an empty token.
// It verifies that the function returns an error and a nil client.
func TestGetClientWithEmptyToken(t *testing.T) {
	token := ""
	client, err := GetClient(token)

	assert.Error(t, err)
	assert.Nil(t, client)
}

// TestGetClientWithInvalidToken tests the GetClient function with an invalid token.
// It mocks the OAuth2 token source to simulate an invalid token and verifies that
// the function returns an error and a nil client.
func TestGetClientWithInvalidToken(t *testing.T) {
	token := "invalid-token"

	// Mock the OAuth2 token source to simulate an invalid token
	mockTokenSource := new(mockOAuth2TokenSource)
	mockTokenSource.On("Token").Return(nil, errors.New("invalid token"))

	// Override the newOAuth2Client function to return nil for invalid tokens
	newOAuth2Client = func(ctx context.Context, src oauth2.TokenSource) *http.Client {
		return nil
	}

	client, err := GetClient(token)

	assert.Error(t, err)
	assert.Nil(t, client)
}
