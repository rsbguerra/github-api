package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusNoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}
func StatusCreated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"data": data})
}
func StatusOK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"data": data})
}
