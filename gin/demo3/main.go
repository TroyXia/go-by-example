package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", pingHandler)
		v1.GET("/healthz", healthzHandler)
	}

	r.Run(":8180")
}

func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func healthzHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ready",
	})
}
