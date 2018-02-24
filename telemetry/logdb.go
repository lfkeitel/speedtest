package telemetry

import (
	"encoding/json"
	"fmt"
)

func init() {
	registerTelemDBInit("log", NewLogDB)
}

type LogDB struct{}

func NewLogDB() (TelemetryDB, error) {
	return &LogDB{}, nil
}

func (db *LogDB) Save(t *Telemetry) error {
	log, _ := json.Marshal(t)
	fmt.Printf("%s\n", log)
	return nil
}
