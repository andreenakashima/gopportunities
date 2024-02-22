package main

import (
	"github.com/andreenakashima/gopportunities.git/config"
	"github.com/andreenakashima/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initalization error: %v", err)
		return
	}

	router.Initialize()
}