package metrics

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

func WriteMetricsRecordsToDB(db *sql.DB, records []ResourceMetric) error {
	writer, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := writer.Prepare(`
		INSERT INTO metrics (
			metric_name,
			unit,
			resource_attributes,
			scope,
			aggregation_temporality,
			is_monotonic,
			start_time_unix_nano,
			time_unix_nano,
			value,
			data_point_attributes
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, resourceMetric := range records {
		// Convert resource attributes to JSON
		resourceAttributes, _ := json.Marshal(resourceMetric.Resource.Attributes)

		for _, scopeMetric := range resourceMetric.ScopeMetrics {
			scope, _ := json.Marshal(scopeMetric.Scope)

			for _, metric := range scopeMetric.Metrics {
				// For now, only use sum metrics because they are the one in the example, need to change in future!!!
				for _, dp := range metric.Sum.DataPoints {
					// Convert data point attributes to JSON
					dataPointAttributes, _ := json.Marshal(dp.Attributes)

					// Convert value to float as its a string
					var valFloat float64
					if dp.AsInt != "" {
						_, err := fmt.Sscan(dp.AsInt, &valFloat)
						if err != nil {
							// Skip for now without error
							continue
						}
					}

					_, err := stmt.Exec(
						metric.Name,
						metric.Unit,
						string(resourceAttributes),
						string(scope),
						metric.Sum.AggregationTemporality,
						metric.Sum.IsMonotonic,
						dp.StartTimeUnixNano,
						dp.TimeUnixNano,
						valFloat,
						string(dataPointAttributes),
					)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return writer.Commit()
}
