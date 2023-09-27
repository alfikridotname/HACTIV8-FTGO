package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// LogLevels represents the different log levels.
type LogLevels struct {
	Info  int
	Error int
	Debug int
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error:", r)
		}
	}()

	if len(os.Args) != 2 {
		fmt.Println("Usage: log_analyzer <log_file_path>")
		return
	}

	logFilePath := os.Args[1]

	logFile, err := os.Open(logFilePath)
	if err != nil {
		panic("Error opening log file: " + err.Error())
	}
	defer logFile.Close()

	logLevels := LogLevels{}
	var wg sync.WaitGroup

	scanner := bufio.NewScanner(logFile)

	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1)
		go func(logLine string) {
			defer wg.Done()
			processLogLine(logLine, &logLevels)
		}(line)
	}

	wg.Wait()

	fmt.Println("Log Analysis Results:")
	fmt.Printf("[INFO] Level : %d Occurrences\n", logLevels.Info)
	fmt.Printf("[ERROR] Level : %d Occurrences\n", logLevels.Error)
	fmt.Printf("[DEBUG] Level : %d Occurrences\n", logLevels.Debug)
}

func processLogLine(logLine string, logLevels *LogLevels) {
	// Split the log line into parts: timestamp, log level, and message
	parts := strings.Fields(logLine)
	if len(parts) < 3 {
		return // Invalid log format, skip this line
	}

	logLevel := parts[2]

	switch logLevel {
	case "[INFO]":
		logLevels.Info++
	case "[ERROR]":
		logLevels.Error++
	case "[DEBUG]":
		logLevels.Debug++
	}
}
