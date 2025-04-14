package auth

import (
	"context"
	"errors"
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
func GetClient(accessToken string) (*github.Client, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, _, err := client.Users.Get(ctx, "")
	if err != nil {
		return nil, errors.New("invalid access token")
	}

	return client, nil
}
