package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tcdw/telegram-home-master/functions"
)

func main() {
	a := functions.Computer{
		Name: "damn",
		Mac: "11:22:33:44:55:66",
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data": a,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}