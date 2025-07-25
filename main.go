package main

import (
	"log"
	"telemetry-ingestor/db"
	"telemetry-ingestor/logs"
	"telemetry-ingestor/traces"
)

// Placeholder filepaths for demonstration
var logFilePath = "files/reference/logs-empty.jsonl"
var traceFilePath = "files/reference/traces-empty.jsonl"

func main() {
	logsSetup()
	tracesSetup()
}

func logsSetup() {
	// Parse logs file
	logRecords, err := logs.ParseLogFile(logFilePath)
	if err != nil {
		log.Fatal("Error parsing log file: ", err)
	}
	log.Print("Parsed Log Records: ", logRecords)

	// Setup logs database
	dbLogs := db.InitializeLogsDatabase("logs.db")
	defer dbLogs.Close()

	// Write log records to database
	if err := logs.WriteLogRecordsToDB(dbLogs, logRecords); err != nil {
		log.Fatal("Failed to write log records to database: ", err)
	}
	log.Print("Successfully wrote log records to database")
}

func tracesSetup() {
	// Parse logs file
	traceRecords, err := traces.ParseTraceFile(traceFilePath)
	if err != nil {
		log.Fatal("Error parsing trace file: ", err)
	}
	log.Print("Parsed Trace Records: ", traceRecords)

	// Setup logs database
	dbTraces := db.InitializeTracesDatabase("traces.db")
	defer dbTraces.Close()

	// Write log records to database
	if err := traces.WriteTraceRecordsToDB(dbTraces, traceRecords); err != nil {
		log.Fatal("Failed to write trace records to database: ", err)
	}
	log.Print("Successfully wrote trace records to database")
}
