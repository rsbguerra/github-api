package auth

import (
	"context"
	"errors"
	"github-api/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
	"net/http"
	"os"
	"testing"
)

type mockOAuth2TokenSource struct {
	mock.Mock
}

var newOAuth2Client = func(ctx context.Context, src oauth2.TokenSource) *http.Client {
	return oauth2.NewClient(ctx, src)
}

func (m *mockOAuth2TokenSource) Token() (*oauth2.Token, error) {
	args := m.Called()
	return args.Get(0).(*oauth2.Token), args.Error(1)
}

func TestGetClientWithValidToken(t *testing.T) {
	validToken := os.Getenv("TEST_AUTH_TOKEN")
	client, err := GetClient(validToken)

	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.IsType(t, &models.GitHubClientWrapper{}, client)
}

func TestGetClientWithEmptyToken(t *testing.T) {
	token := ""
	client, err := GetClient(token)

	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestGetClientWithInvalidToken(t *testing.T) {
	token := "invalid-token"

	// Mock the OAuth2 token source to simulate an invalid token
	mockTokenSource := new(mockOAuth2TokenSource)
	mockTokenSource.On("Token").Return(nil, errors.New("invalid token"))

	newOAuth2Client = func(ctx context.Context, src oauth2.TokenSource) *http.Client {
		return nil
	}

	client, err := GetClient(token)

	assert.Error(t, err)
	assert.Nil(t, client)
}
