package main

import (
	"os"
	"path"
)

// Run executres the code needed to generate the static site from the source files
func Run(logger Logger) {
	logger.LogInformation("Reading configuration")
	config, err := ReadConfig(path.Join(inputDir, "config.yaml"))
	if err != nil {
		logger.LogFatal(err.Error())
		os.Exit(1)
	}
	logger.LogInformation(config.Title)
}

// Serve starts a localhost for development previewing
func Serve(logger Logger) {
	logger.LogInformation("Inside Serve")
}
