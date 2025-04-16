package response

import (
	"github-api/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
}
func StatusBadRequestMissingParams(c *gin.Context, missingParams []string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":          "Missing required parameters",
		"missing_params": missingParams,
	})
}
func StatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
}
func StatusInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
}
func StatusNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
}
func StatusForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
}
func StatusUnprocessableEntity(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unprocessable entity: " + err.Error()})
}

func StatusConflict(c *gin.Context, repo models.RepositoryModel) {
	c.JSON(http.StatusConflict, gin.H{"error": "Conflict, repository already exists", "repo": repo})
}
