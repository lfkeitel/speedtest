package telemetry

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"time"
)

func init() {
	registerTelemDBInit("csv", NewCSVLog)
}

type CSVLog struct{}

func NewCSVLog() (TelemetryDB, error) {
	if dbFilename == "" {
		return nil, errors.New("dbfilename must be non-empty")
	}
	return &CSVLog{}, nil
}

func (db *CSVLog) Save(t *Telemetry) error {
	file, err := os.OpenFile(dbFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{
		t.Timestamp.Format(time.RFC3339),
		t.Remote.String(),
		t.UA,
		strconv.FormatFloat(t.DL, 'f', 2, 64),
		strconv.FormatFloat(t.UL, 'f', 2, 64),
		strconv.FormatFloat(t.Ping, 'f', 2, 64),
		strconv.FormatFloat(t.Jitter, 'f', 2, 64),
		t.Building,
		t.SessionID,
	})
	writer.Flush()

	return err
}
