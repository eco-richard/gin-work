package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	err := r.Run() // running on port :8080
	if err != nil {
		log.Fatalf("Error running server: %s", err.Error())
	}
}
