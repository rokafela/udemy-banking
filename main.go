package main

import (
	"github.com/rokafela/udemy-banking/app"
	"github.com/rokafela/udemy-banking/logger"
)

func main() {
	logger.Info("application starting")
	app.Start()
}
