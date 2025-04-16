package auth

import (
	"context"
	"github-api/pkg/interfaces"
	"github-api/pkg/models"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

// GetClient creates a new GitHub client using the provided access token.
// It validates the access token by making a request to the GitHub API.
// If the token is valid, it returns a GitHub client.
// If the token is invalid, it returns an error.
//
// Parameters:
//   - accessToken: The GitHub access token to authenticate the client.
//
// Returns:
//   - *github.Client: A GitHub client authenticated with the provided access token.
//   - error: An error if the access token is invalid or if there was an issue creating the client.
func GetClient(token string) (interfaces.GitHubClient, error) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)
	client := github.NewClient(tc)

	// Validate the token by making a test request to the GitHub API
	_, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err // Return the error if the token is invalid
	}

	return &models.GitHubClientWrapper{Client: client}, nil
}
