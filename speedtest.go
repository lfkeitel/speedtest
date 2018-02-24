package main

import (
	"encoding/json"
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
)

const garbageLen = 1048576

var (
	contentTypes = map[string]string{
		".css":  "text/css",
		".js":   "text/javascript",
		".html": "text/html",
	}

	garbage []byte

	httpAddress string
)

type telemetry struct {
	Remote string `json:"remote"`
	DL     string `json:"dl"`
	UL     string `json:"ul"`
	Ping   string `json:"ping"`
	Jitter string `json:"jitter"`
	Log    string `json:"log"`
}

func init() {
	rand.Seed(time.Now().UnixNano())

	flag.StringVar(&httpAddress, "addr", "localhost:8080", "HTTP address and port to bind")
}

func main() {
	flag.Parse()
	generateGarbage()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		file := r.URL.Path
		if file == "/" {
			file = "/index.html"
		}
		if strings.Contains(file, "..") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		file = filepath.Join("public", file)
		fmt.Printf("Request for %s\n", file)
		http.ServeFile(w, r, file)
	})

	http.HandleFunc("/empty.php", func(w http.ResponseWriter, r *http.Request) {
		nocacheHeaders(w)
		w.Header().Add("Connection", "keep-alive")

		// The request was failing, reading the request to /dev/null fixed it
		io.Copy(ioutil.Discard, r.Body)
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/getIP.php", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")

		if r.Header.Get("X-Real-IP") != "" {
			w.Write([]byte(r.Header.Get("X-Real-IP")))
		} else {
			host, _, _ := net.SplitHostPort(r.RemoteAddr)
			w.Write([]byte(host))
		}
	})

	http.HandleFunc("/garbage.php", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/telemetry.php", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(4096)
		host, _, _ := net.SplitHostPort(r.RemoteAddr)

		tely := &telemetry{
			Remote: host,
			DL:     r.PostFormValue("dl"),
			UL:     r.PostFormValue("ul"),
			Ping:   r.PostFormValue("ping"),
			Jitter: r.PostFormValue("jitter"),
			Log:    r.PostFormValue("log"),
		}

		log, _ := json.Marshal(tely)
		fmt.Printf("%s\n", log)
	})

	fmt.Printf("Now listening on http://%s\n", httpAddress)
	if err := http.ListenAndServe(httpAddress, nil); err != nil {
		fmt.Println(err)
	}
}

func nocacheHeaders(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0, post-check=0, pre-check=0")
	w.Header().Add("Pragma", "no-cache")
}

func generateGarbage() {
	fmt.Println("Generating garbage")

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
