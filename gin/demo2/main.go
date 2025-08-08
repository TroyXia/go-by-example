package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "12345")
		log.Println("set example to 12345")
		c.Next()

		latency := time.Since(t)
		log.Print("cost: ", latency)
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	r := gin.Default()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		log.Println(example)

		c.JSON(200, gin.H{
			"message": example,
		})
	})

	r.Run(":8180")
}
