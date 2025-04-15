package v1

import (
	"github-api/pkg/api/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/repositories/:token", controllers.CreateRepo)
	router.DELETE("/repositories/:token", controllers.DeleteRepo)
	router.GET("/repositories/:token", controllers.ListRepos)
	router.GET("/pull-requests/:username/:repoName/:token", controllers.PullRequests)
	router.GET("/", controllers.Index)
}
