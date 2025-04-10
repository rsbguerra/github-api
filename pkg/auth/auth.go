package auth

import (
	"context"
	"errors"
	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

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
