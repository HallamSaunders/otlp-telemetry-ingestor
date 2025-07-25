package db

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

func OverwriteCurrentDBFile(path string) *sql.DB {
	// Remove the file if it exists
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		log.Fatalf("Failed to remove existing database file: %v", err)
	}

	// Create a new database file
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	return db
}

func InitializeLogsDatabase(path string) *sql.DB {
	db := OverwriteCurrentDBFile(path)

	// === Create the logs table based on Logs type defined if it doesn't exist ===
	// TODO: update Body field to match ValueField, Attributes to match []Attribute type
	createLogsTable := `
	CREATE TABLE IF NOT EXISTS logs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		time_unix_nano TEXT NOT NULL,
		severity_number INTEGER NOT NULL,
		severity_text TEXT NOT NULL,
		body TEXT NOT NULL,
		attributes TEXT,
		dropped_attributes_count INTEGER DEFAULT 0,
		trace_id TEXT,
		span_id TEXT
	);`

	if _, err := db.Exec(createLogsTable); err != nil {
		log.Fatalf("Failed to create logs table: %v", err)
	}

	return db
}

func InitializeTracesDatabase(path string) *sql.DB {
	db := OverwriteCurrentDBFile(path)

	// === Create the traces table based on Traces type defined if it doesn't exist ===
	// TODO: update Body field to match ValueField, Attributes to match []Attribute type
	createTracesTable := `
	CREATE TABLE IF NOT EXISTS traces (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		start_time_unix_nano TEXT NOT NULL,
		end_time_unix_nano TEXT NOT NULL,
		name TEXT NOT NULL,
		span_id TEXT NOT NULL,
		trace_id TEXT NOT NULL,
		parent_span_id TEXT,
		attributes TEXT,
		dropped_attributes_count INTEGER DEFAULT 0
	);`

	if _, err := db.Exec(createTracesTable); err != nil {
		log.Fatalf("Failed to create logs table: %v", err)
	}

	return db
}

func InitializeMetricsDatabase(path string) *sql.DB {
	db := OverwriteCurrentDBFile(path)

	// === Create the metrics table based on Metrics type defined if it doesn't exist ===
	createMetricsTable := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		metric_name TEXT NOT NULL,
		unit TEXT,
		resource_attributes TEXT,
		scope TEXT,
		aggregation_temporality TEXT,
		is_monotonic BOOLEAN,
		start_time_unix_nano TEXT,
		time_unix_nano TEXT NOT NULL,
		value REAL,
		data_point_attributes TEXT 
	);`

	if _, err := db.Exec(createMetricsTable); err != nil {
		log.Fatalf("Failed to create metrics table: %v", err)
	}

	return db
}
