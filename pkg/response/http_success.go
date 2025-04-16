package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// StatusNoContent sends a HTTP 204 No Content response.
// This is typically used when a request is successful but there is no content to return.
//
// Parameters:
//   - c: The Gin context for the current HTTP request.
func StatusNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

// StatusCreated sends a HTTP 201 Created response with the provided data.
// This is typically used when a resource is successfully created.
//
// Parameters:
//   - c: The Gin context for the current HTTP request.
//   - data: The data to include in the response body.
func StatusCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"data": data})
}

// StatusOK sends a HTTP 200 OK response with the provided data.
// This is typically used for successful requests that return data.
//
// Parameters:
//   - c: The Gin context for the current HTTP request.
//   - data: The data to include in the response body.
func StatusOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
