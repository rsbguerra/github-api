package v1

import (
	"github-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/:token/create-repo/", controllers.CreateRepo)
	router.DELETE("/:token/delete-repo/", controllers.DeleteRepo)
	router.GET("/:token/list-repos/", controllers.ListRepos)
}
