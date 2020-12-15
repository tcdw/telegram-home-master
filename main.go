package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/tcdw/telegram-home-master/config"
	"github.com/tcdw/telegram-home-master/route"
	"os"
	"path/filepath"
)

func fatalError(err error) {
	if err != nil {
		fmt.Printf("Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func PassConfig(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", conf)
		c.Next()
	}
}

func main() {
	var opts struct {
		ConfigPath string `short:"c" long:"config" description:"Config file path" required:"true"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	configPath, err := filepath.Abs(opts.ConfigPath)
	fatalError(err)

	conf, err := config.ReadConfig(configPath)
	fatalError(err)

	r := gin.Default()
	r.Use(PassConfig(*conf))

	v1 := r.Group("/api/v1")
	v1.POST("on", route.TurnOn)

	r.Run()
}