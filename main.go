package main

import (
	"github.com/FelipeDaros/gooportunes/config"
	"github.com/FelipeDaros/gooportunes/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")
	err := config.Init()

	if err != nil {
		logger.Errorf("config", err)
		return
	}

	router.Initialize()
}
