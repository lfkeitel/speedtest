package telemetry

import (
	"fmt"
)

func init() {
	registerTelemDBInit("console", NewLogDB)
}

type LogDB struct{}

func NewLogDB() (TelemetryDB, error) {
	return &LogDB{}, nil
}

func (db *LogDB) Save(t *Telemetry) error {
	fmt.Println(t.String())
	return nil
}
