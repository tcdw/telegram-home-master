package route

import (
	"github.com/gin-gonic/gin"
	"github.com/tcdw/telegram-home-master/config"
	"github.com/tcdw/telegram-home-master/wake"
)

type TurnOnParam struct {
	Auth config.Auth `json:"auth" binding:"required"`
	Name string      `json:"name" binding:"required"`
}

func TurnOn(c *gin.Context) {
	var json TurnOnParam
	err := c.ShouldBindJSON(&json)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	conf, exists := c.MustGet("config").(config.Config)
	if !exists {
		c.JSON(400, gin.H{
			"success": false,
		})
		return
	}

	err = wake.ByName(conf, json.Name)
	c.JSON(200, gin.H{
		"success": err == nil,
	})
}