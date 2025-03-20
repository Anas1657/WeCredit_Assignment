package main

import (
	"log"
	"os"
	"time"
)

var (
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
)

func init() {
	// Create a log file
	file, err := os.OpenFile("/var/log/go-logging-app/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create different loggers for different levels
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger = log.New(file, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	for {
		infoLogger.Println("This is an INFO message.")
		warnLogger.Println("This is a WARNING message.")
		debugLogger.Println("This is a DEBUG message.")
		errorLogger.Println("This is an ERROR message.")

		time.Sleep(5 * time.Second) // Log every 5 seconds
	}
}
