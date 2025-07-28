package traces

import (
	"database/sql"
	"strconv"
)

func WriteTraceRecordsToDB(db *sql.DB, records []ResourceSpan) error {
	writer, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := writer.Prepare(`
		INSERT INTO traces (
			spanID,
			parentSpanID,
			traceID,
			startTime,
			duration,
			serviceName,
			operationName)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, resourceSpan := range records {
		// Serialize resource-level attributes
		for _, scopeSpan := range resourceSpan.ScopeSpans {
			for _, span := range scopeSpan.Spans {
				// Find duration in nanoseconds
				startTime, err := strconv.ParseInt(span.StartTimeUnixNano, 10, 64)
				if err != nil {
					return err
				}
				endTime, err := strconv.ParseInt(span.EndTimeUnixNano, 10, 64)
				if err != nil {
					return err
				}
				duration := endTime - startTime
				startTimeMilli := startTime / 1_000_000
				durationMilli := duration / 1_000_000

				_, err = stmt.Exec(
					span.SpanID,
					span.ParentSpanID,
					span.TraceID,
					startTimeMilli,
					durationMilli,
					span.Name,
					span.Name,
				)
				if err != nil {
					return err
				}
			}
		}
	}
	return writer.Commit()
}
