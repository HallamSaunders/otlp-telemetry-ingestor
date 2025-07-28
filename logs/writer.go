package logs

import (
	"database/sql"
	"encoding/json"
	"strconv"
	"time"
)

func WriteLogRecordsToDB(db *sql.DB, records []LogRecord) error {
	writer, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := writer.Prepare(`
		INSERT INTO logs (
			timestamp,
			severity_number,
			severity_text,
			body,
			attributes,
			dropped_attributes_count,
			trace_id,
			span_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Now to actually write the records to the database
	for _, rec := range records {
		attrsJSON, _ := json.Marshal(rec.Attributes) // Needs an update to handle different types of attribute

		// Get time as timestamp for ease of viewing
		unixNano, err := strconv.ParseInt(rec.TimeUnixNano, 10, 64)
		if err != nil {
			return err
		}
		timestamp := time.Unix(0, unixNano).Format(time.RFC3339Nano)

		_, err = stmt.Exec(
			timestamp,
			rec.SeverityNumber,
			rec.SeverityText,
			rec.Body.StringValue,
			rec.DroppedAttributesCount,
			rec.TraceID,
			rec.SpanID,
			string(attrsJSON),
		)
		if err != nil {
			return err
		}
	}
	return writer.Commit()
}
