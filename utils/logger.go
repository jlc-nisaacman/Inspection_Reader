package utils

import (
	"log"
	"os"
	"sync"
)

// Global mutex for thread-safe logging
var logMutex sync.Mutex

// Setup logging to a file
func SetupLogging(logFile string) {
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if err != nil {
		log.Fatalf("Could not open log file: %v", err)
	}
	log.SetOutput(file)
}

// Thread-safe log function
func LogSafe(format string, v ...interface{}) {
	logMutex.Lock()
	defer logMutex.Unlock()
	log.Printf(format, v...)
}
