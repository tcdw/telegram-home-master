package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tcdw/telegram-home-master/config"
)

type TurnOnParam struct {
	Auth config.Auth `json:"user" binding:"required"`
	Name string `json:"password" binding:"required"`
}

func TurnOn(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": true,
	})
}