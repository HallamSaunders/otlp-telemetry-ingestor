package logs

import (
	"bufio"
	"encoding/json"
	"os"
)

func ParseLogFile(filePath string) ([]LogRecord, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err // If there's an error reading the file, return nil and the error
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var logsFromFile []LogRecord

	for scanner.Scan() {
		var logFile LogFile
		// Unmarshal the JSON data into the LogFile struct I created
		if err := json.Unmarshal(scanner.Bytes(), &logFile); err != nil {
			return nil, err
		}
		for _, resourceLog := range logFile.ResourceLogs {
			for _, scopeLog := range resourceLog.ScopeLogs {
				/* Append each LogRecord to the logsFromFile slice
				This assumes the logs are always in the expected format,
				(which they should be since it's OTLP). */
				logsFromFile = append(logsFromFile, scopeLog.LogRecords...)
			}
		}
	}

	return logsFromFile, scanner.Err() // Return the slice of LogRecords and any error from scanning
}
