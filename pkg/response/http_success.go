package response

import (
	"github-api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusCreated(c *gin.Context, repo models.Repository) {
	c.JSON(http.StatusCreated, gin.H{"data": repo})
}
func StatusOK(c *gin.Context, repo models.Repository) {
	c.JSON(http.StatusOK, gin.H{"data": repo})
}
