package traces

import (
	"database/sql"
	"encoding/json"
)

func WriteTraceRecordsToDB(db *sql.DB, records []ResourceSpan) error {
	writer, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := writer.Prepare(`
		INSERT INTO traces (
			start_time_unix_nano,
			end_time_unix_nano,
			name,
			span_id,
			trace_id,
			parent_span_id,
			attributes,
			dropped_attributes_count)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, resourceSpan := range records {
		// Serialize resource-level attributes
		resourceAttrsJSON, _ := json.Marshal(resourceSpan.Resource.Attributes)

		for _, scopeSpan := range resourceSpan.ScopeSpans {
			for _, span := range scopeSpan.Spans {
				// You could also include span-specific attributes if added later
				_, err := stmt.Exec(
					span.StartTimeUnixNano,
					span.EndTimeUnixNano,
					span.Name,
					span.SpanID,
					span.TraceID,
					span.ParentSpanID,
					string(resourceAttrsJSON),
					span.DroppedAttributesCount,
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return writer.Commit()
}
