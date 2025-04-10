package controllers

import (
	"github-api/pkg/auth"
	"github-api/pkg/models"
	"github-api/pkg/response"
	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var repo, err = models.NewRepository(c)
	if (repo == models.Repository{}) || err != nil {
		response.StatusBadRequest(c)
		return
	}
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.StatusUnauthorized(c)
		return
	}
	if err := repo.CreateNew(client); err != nil {
		response.StatusInternalServerError(c, err)
		return
	}
	response.StatusCreated(c, repo)
}

func DeleteRepo(c *gin.Context) {
	var repo, err = models.NewRepository(c)
	if (repo == models.Repository{}) || err != nil {
		response.StatusBadRequest(c)
		return
	}
	client, err := auth.GetClient(c.Param("token"))
	if err != nil {
		response.StatusUnauthorized(c)
		return
	}
	if err := repo.DeleteRepo(client); err != nil {
		response.StatusInternalServerError(c, err)
	}

}

func ListRepos(c *gin.Context) {

}

func PullRequests(c *gin.Context) {

}
