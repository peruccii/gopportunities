package main

import (
	"github.com/peruccii/gopportunities/config"
	"github.com/peruccii/gopportunities/router"
	
)

var (
	logger  *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil { // if error exists
		logger.ErrF("config initialize error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}