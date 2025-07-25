package metrics

import (
	"bufio"
	"encoding/json"
	"os"
)

func ParseMetricsFile(filePath string) ([]ResourceMetric, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var resourceMetrics []ResourceMetric

	for scanner.Scan() {
		var traceFile MetricsFile
		// Unmarshal the JSON data into the TraceFile struct
		if err := json.Unmarshal(scanner.Bytes(), &traceFile); err != nil {
			return nil, err
		}
		resourceMetrics = append(resourceMetrics, traceFile.ResourceMetrics...)
	}

	return resourceMetrics, scanner.Err() // Return the slice of ResourceSpans and any error from scanning
}
