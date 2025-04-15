package v1

import (
	"github-api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/repositories/:token", controllers.CreateRepo)
	router.DELETE("/repositories/:token", controllers.DeleteRepo)
	router.GET("/users/:username/repositories/:token", controllers.ListRepos)
	router.GET("/repositories/:username/:repoName/pull-requests", controllers.PullRequests)
	router.GET("/", controllers.Index)
}
