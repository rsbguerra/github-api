package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// StatusBadRequest sends a 400 Bad Request response with a generic error message.
// This is used when the request payload is invalid.
func StatusBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
}

// StatusBadRequestMissingParams sends a 400 Bad Request response with a message
// indicating missing required parameters.
// Parameters:
// - c: The Gin context.
// - missingParams: A slice of strings representing the missing parameters.
func StatusBadRequestMissingParams(c *gin.Context, missingParams []string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":          "Missing required parameters",
		"missing_params": missingParams,
	})
}

// StatusUnauthorized sends a 401 Unauthorized response with a generic error message.
// This is used when the access token is invalid or missing.
func StatusUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
}

// StatusInternalServerError sends a 500 Internal Server Error response with a detailed error message.
// Parameters:
// - c: The Gin context.
// - err: The error that occurred.
func StatusInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: " + err.Error()})
}

// StatusNotFound sends a 404 Not Found response with a generic error message.
// This is used when the requested resource cannot be found.
func StatusNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
}

// StatusForbidden sends a 403 Forbidden response with a generic error message.
// This is used when the user does not have permission to access the resource.
func StatusForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
}

// StatusUnprocessableEntity sends a 422 Unprocessable Entity response with a detailed error message.
// Parameters:
// - c: The Gin context.
// - err: The error that occurred.
func StatusUnprocessableEntity(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Unprocessable entity: " + err.Error()})
}

// StatusConflict sends a 409 Conflict response with a generic error message.
// This is used when a conflict occurs, such as when a repository already exists.
func StatusConflict(c *gin.Context) {
	c.JSON(http.StatusConflict, gin.H{"error": "Conflict, repository already exists"})
}
