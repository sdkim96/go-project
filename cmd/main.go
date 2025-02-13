package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(ErrorHandlerMiddleware())
	router.Use(gin.Recovery())

	router.GET("/panic", func(c *gin.Context) {
		panic("Panic!")
	})

	router.Run(":8080")
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}
}

// main.go
