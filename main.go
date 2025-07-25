package main

import (
	"log"
	"telemetry-ingestor/logs"
)

// Placeholder filepaths for demonstration
var logFilePath = "files/reference/logs-empty.jsonl"

func main() {
	logRecords, err := logs.ParseLogFile(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Parsed Log Records: ", logRecords)
}
