package telemetry

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres database driver
)

func init() {
	registerTelemDBInit("psql", NewPostgresDB)
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (TelemetryDB, error) {
	connStr := fmt.Sprintf("user='%s' password='%s' dbname='%s' host='%s' port=%d sslmode=disable timezone=UTC",
		dbUser,
		dbPass,
		dbName,
		dbAddr,
		dbPort,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{db: db}, nil
}

func (db *PostgresDB) Save(t *Telemetry) error {
	_, err := db.db.Exec(`INSERT INTO telemetry
		("timestamp", ip, ua, dl, ul, ping, jitter, building, sessionid)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		t.Timestamp,
		t.Remote.String(),
		t.UA,
		t.DL,
		t.UL,
		t.Ping,
		t.Jitter,
		t.Building,
		t.SessionID,
	)
	return err
}
