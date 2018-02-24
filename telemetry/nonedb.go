package telemetry

func init() {
	registerTelemDBInit("none", NewNoneDB)
}

type NoneDB struct{}

func NewNoneDB() (TelemetryDB, error)      { return &NoneDB{}, nil }
func (db *NoneDB) Save(t *Telemetry) error { return nil }
