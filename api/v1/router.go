package v1

import (
	controllers2 "github-api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/repositories/:token", controllers2.CreateRepo)
	router.DELETE("/repositories/:token", controllers2.DeleteRepo)
	router.GET("/users/:username/repositories/:token", controllers2.ListRepos)
	router.GET("/repositories/:username/:repoName/pull-requests", controllers2.PullRequests)
	router.GET("/", controllers2.Index)
}
