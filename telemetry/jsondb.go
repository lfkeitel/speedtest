package telemetry

import (
	"encoding/json"
	"errors"
	"os"
)

func init() {
	registerTelemDBInit("json", NewJSONLog)
}

type JSONLog struct{}

func NewJSONLog() (TelemetryDB, error) {
	if dbFilename == "" {
		return nil, errors.New("dbfilename must be non-empty")
	}
	return &JSONLog{}, nil
}

func (db *JSONLog) Save(t *Telemetry) error {
	file, err := os.OpenFile(dbFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	j, _ := json.Marshal(t)
	_, err = file.Write(j)
	file.Write([]byte{'\n'})
	return err
}
