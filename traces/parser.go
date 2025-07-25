package traces

import (
	"bufio"
	"encoding/json"
	"os"
)

func ParseTraceFile(filePath string) ([]ResourceSpan, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var resourceSpans []ResourceSpan

	for scanner.Scan() {
		var traceFile TraceFile
		// Unmarshal the JSON data into the TraceFile struct
		if err := json.Unmarshal(scanner.Bytes(), &traceFile); err != nil {
			return nil, err
		}
		resourceSpans = append(resourceSpans, traceFile.ResourceSpans...)
	}

	return resourceSpans, scanner.Err() // Return the slice of ResourceSpans and any error from scanning
}
