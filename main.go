package main

import (
	"github.com/luizengdev/gopportunities/config"
	"github.com/luizengdev/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// initialize Logger
	logger = config.GetLogger("main")

	// initialize Config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
