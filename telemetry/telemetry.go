package telemetry

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

var (
	dbAddr, dbUser, dbPass, dbName string
	dbPort                         int
	dbFilename                     string

	pow10to4 = 9999.0
)

func init() {
	flag.StringVar(&dbAddr, "dbaddr", "localhost", "Database address")
	flag.StringVar(&dbUser, "dbuser", "speedtest", "Database username")
	flag.StringVar(&dbPass, "dbpass", "", "Database password")
	flag.StringVar(&dbName, "dbname", "speedtest", "Database name")
	flag.IntVar(&dbPort, "dbport", 0, "Database name")
	flag.StringVar(&dbFilename, "dbfile", "telemetry.txt", "Database filepath")
}

type TelemetryDB interface {
	Save(*Telemetry) error
}

type dbInit func() (TelemetryDB, error)

var telemetryDBInits = map[string]dbInit{}

type Telemetry struct {
	Timestamp time.Time `json:"timestamp"`
	Remote    net.IP    `json:"remote"`
	UA        string    `json:"ua"`
	DL        float64   `json:"dl"`
	UL        float64   `json:"ul"`
	Ping      float64   `json:"ping"`
	Jitter    float64   `json:"jitter"`
	Log       string    `json:"log"`
	Building  string    `json:"building"`
	SessionID string    `json:"sessionid"`
}

func (t *Telemetry) String() string {
	return fmt.Sprintf("Telemetry: Client: %s, DL: %f, UL: %f, Ping: %f, Jitter: %f", t.Remote, t.DL, t.UL, t.Ping, t.Jitter)
}

func registerTelemDBInit(name string, i dbInit) {
	telemetryDBInits[name] = i
}

func MakeDB(name string) (TelemetryDB, error) {
	initfn, exists := telemetryDBInits[name]
	if !exists {
		return nil, fmt.Errorf("database type %s doesn't exist", name)
	}

	return initfn()
}

func HTTPHandler(db TelemetryDB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(4096)
		host, _, _ := net.SplitHostPort(r.RemoteAddr)

		sessionID := ""
		sessionIDCookie, err := r.Cookie("st_sessionid")
		if err == nil {
			sessionID = sessionIDCookie.Value
		}

		dlMetric, err := strconv.ParseFloat(r.PostFormValue("dl"), 64)
		if err != nil {
			fmt.Println(err)
		}
		ulMetric, err := strconv.ParseFloat(r.PostFormValue("ul"), 64)
		if err != nil {
			fmt.Println(err)
		}
		pingMetric, err := strconv.ParseFloat(r.PostFormValue("ping"), 64)
		if err != nil {
			fmt.Println(err)
		}
		jitterMetric, err := strconv.ParseFloat(r.PostFormValue("jitter"), 64)
		if err != nil {
			fmt.Println(err)
		}

		tely := &Telemetry{
			Timestamp: time.Now().In(time.UTC),
			Remote:    net.ParseIP(host),
			UA:        r.UserAgent(),
			DL:        truncNumber(dlMetric),
			UL:        truncNumber(ulMetric),
			Ping:      truncNumber(pingMetric),
			Jitter:    truncNumber(jitterMetric),
			Log:       r.PostFormValue("log"),
			Building:  "Unknown",
			SessionID: sessionID,
		}

		if err := db.Save(tely); err != nil {
			fmt.Println(err)
		}
	})
}

// Truncate numbers below 10**4
func truncNumber(in float64) float64 {
	if in > pow10to4 {
		return pow10to4
	}
	return in
}
