package v1

import (
	"github-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/create-repo/:token", controllers.CreateRepo)
	router.DELETE("/delete-repo/:token", controllers.DeleteRepo)
	router.GET("/list-repos/:username/:token", controllers.ListRepos)
	router.GET("/pull-requests/:username/:repoName/:token", controllers.PullRequests)
	router.GET("/", controllers.Index)
}
