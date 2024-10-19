package main

import (
	"github.com/jamadeu/api-cars/config"
	"github.com/jamadeu/api-cars/router"
)

func main() {
	logger := config.GetLogger("main")
	// Initialize configs
	if err := config.Init(); err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}
