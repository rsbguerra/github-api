package main

import (
	"github-api/pkg/api/v1"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	v1.RegisterRoutes(router)
	log.Println("Server started at :8080")
	router.Run(":8080")
}
