package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitializeDatabase(path string) *sql.DB {
	// First, attempt to open the database at the specified path
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

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
