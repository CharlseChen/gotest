package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	router := r.Group("/needToken")
	router.Use(JwtMiddleware())
	{
		r.GET("/test", func(s *gin.Context) {
			s.JSON(http.StatusOK, gin.H{"message": "This is protected data"})
		})
	}
	r.Group("/noNeedToken", func(c *gin.Context) {
		print("test")
	})
	r.Run(":8080")
}
