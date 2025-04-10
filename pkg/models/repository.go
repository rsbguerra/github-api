package models

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v50/github"
	"os"
)

type Repository struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Owner   string `json:"owner"`
	Private bool   `json:"private"`
}

func NewRepository(c *gin.Context) (Repository, error) {
	var repo Repository

	if err := c.ShouldBindJSON(&repo); err != nil {
		return Repository{}, err
	}
	return repo, nil
}

func (r *Repository) CloneRepo() error {
	_, err := git.PlainClone(r.Name, false, &git.CloneOptions{
		URL:      r.URL,
		Progress: os.Stdout,
	})
	return err
}

func (r *Repository) CreateNew(client *github.Client) error {

	repo := &github.Repository{
		Name:    github.String(r.Name),
		Private: github.Bool(r.Private),
	}

	_, _, err := client.Repositories.Create(context.Background(), "", repo)
	return err
}

func (r *Repository) DeleteRepo(client *github.Client) error {
	_, err := client.Repositories.Delete(context.Background(), r.Owner, r.Name)
	return err
}

//func (r *Repository) ListPullRequests(client *github.Client) (*github.Repository, error) {
// TODO
//}
