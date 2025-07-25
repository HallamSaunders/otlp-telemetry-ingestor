package logs

import (
	"database/sql"
	"encoding/json"
)

func WriteLogRecordsToDB(db *sql.DB, records []LogRecord) error {
	writer, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := writer.Prepare(`
		INSERT INTO logs (
			time_unix_nano,
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
		attrsJSON, _ := json.Marshal(rec.Attributes)

		_, err := stmt.Exec(
			rec.TimeUnixNano,
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
