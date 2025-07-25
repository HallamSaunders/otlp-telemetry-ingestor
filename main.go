package main

import (
	"log"
	"telemetry-ingestor/db"
	"telemetry-ingestor/logs"
)

// Placeholder filepaths for demonstration
var logFilePath = "files/reference/logs-empty.jsonl"

func main() {
	// Parse the logs file
	logRecords, err := logs.ParseLogFile(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Parsed Log Records: ", logRecords)

	// Set up the database
	db := db.InitializeDatabase("logs.db")
	defer db.Close()

	// Insert the parsed log records into the database
	if err := logs.WriteLogRecordsToDB(db, logRecords); err != nil {
		log.Fatal("Failed to write log records to database: ", err)
	}
	log.Print("Successfully wrote log records to database")
}
