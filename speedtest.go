package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/lfkeitel/speedtest/telemetry"
	"github.com/satori/go.uuid"
)

const garbageLen = 1048576

var (
	contentTypes = map[string]string{
		".css":  "text/css",
		".js":   "text/javascript",
		".html": "text/html",
	}

	garbage []byte

	httpAddress     string
	telemetryDBType string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&httpAddress, "addr", ":8080", "HTTP address and port to bind")
	flag.StringVar(&telemetryDBType, "db", "none", "Telemetry database type")
}

func main() {
	flag.Parse()
	generateGarbage()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/empty.php", emptyHandler)
	http.HandleFunc("/getIP.php", getIPHandler)
	http.HandleFunc("/garbage.php", garbageHandler)

	teleDB, err := telemetry.MakeDB(telemetryDBType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	http.HandleFunc("/telemetry.php", telemetry.HTTPHandler(teleDB))

	if httpAddress[0] == ':' {
		fmt.Printf("Now listening on http://localhost%s\n", httpAddress)
	} else {
		fmt.Printf("Now listening on http://%s\n", httpAddress)
	}

	if err := http.ListenAndServe(httpAddress, nil); err != nil {
		fmt.Println(err)
	}
}

func nocacheHeaders(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
}

func generateGarbage() {
	garbage = make([]byte, garbageLen)
	n, err := rand.Read(garbage)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if n != garbageLen {
		fmt.Println("Failed to generate garbage")
		os.Exit(1)
	}
}

func setSessionID(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("st_sessionid")
	if err != nil || sessionID.Value == "" {
		http.SetCookie(w, &http.Cookie{
			Name:     "st_sessionid",
			Value:    uuid.NewV4().String(),
			HttpOnly: true,
		})
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	setSessionID(w, r)

	file := r.URL.Path
	if file == "/" {
		file = "/index.html"
	}
	if strings.Contains(file, "..") {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	file = filepath.Join("public", file)
	http.ServeFile(w, r, file)
}

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	setSessionID(w, r)

	nocacheHeaders(w)
	w.Header().Add("Connection", "keep-alive")

	// The request was failing, reading the request to /dev/null fixed it
	io.Copy(ioutil.Discard, r.Body)
	w.WriteHeader(http.StatusOK)
}

func getIPHandler(w http.ResponseWriter, r *http.Request) {
	setSessionID(w, r)

	w.Header().Add("Content-Type", "text/plain; charset=utf-8")

	if r.Header.Get("X-Real-IP") != "" {
		w.Write([]byte(r.Header.Get("X-Real-IP")))
	} else {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		w.Write([]byte(host))
	}
}

func garbageHandler(w http.ResponseWriter, r *http.Request) {
	setSessionID(w, r)

	nocacheHeaders(w)

	w.Header().Add("Content-Description", "File Transfer")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=random.dat")
	w.Header().Add("Content-Transfer-Encoding", "binary")

	chunksGet := r.URL.Query().Get("ckSize")
	chunks, err := strconv.Atoi(chunksGet)
	if err != nil || chunks < 0 {
		chunks = 4
	}
	if chunks > 100 {
		chunks = 100
	}

	for i := 1; i < chunks; i++ {
		w.Write(garbage)
	}
}
