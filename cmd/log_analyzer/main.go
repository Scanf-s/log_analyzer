package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"log_analyzer/internal/runner"
)

func main() {
	projectRoot, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current working directory: %s", err)
		return
	}

	defaultLogPath := filepath.Join(projectRoot, "testdata")
	logPath := flag.String("log-path", defaultLogPath, "Absolute Path to the log file directory to analyze")
	flag.Parse()

	log.Println("Simple Log Analyzer")
	log.Println("Version 0.0.1")
	log.Println("Go version 1.25.1")
	log.Printf("Starting Log Analyzer with log path: %s\n", *logPath)

	runner.Run(*logPath)
}
