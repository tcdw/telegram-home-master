package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tcdw/telegram-home-master/route"
)

func main() {
	r := gin.Default()
	r.POST("/on", route.TurnOn)
	r.Run()
}