package main

import (
	"github.com/hchaudhari73/banking/app"
	"github.com/hchaudhari73/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
