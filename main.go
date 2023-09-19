package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var started = time.Now()

func healthz(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(started)

	if duration.Seconds() > 20 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	dbURL := os.Getenv("DB_URL")
	io.WriteString(w, fmt.Sprintf("[v4] Hello, Kubernetes! host: %s, DataBase URL: %s", host, dbURL))
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
