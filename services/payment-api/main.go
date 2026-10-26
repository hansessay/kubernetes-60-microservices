package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type PlatformInfo struct {
	Service  string `json:"service"`
	Version  string `json:"version"`
	Platform string `json:"platform"`
	Hostname string `json:"hostname"`
}

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "OpenShift Platform Engineering API",
		})
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	mux.HandleFunc("/api/v1/platform", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()

		info := PlatformInfo{
			Service:  "platform-api",
			Version:  "0.1.0",
			Platform: "OpenShift",
			Hostname: hostname,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(info)
	})

	return mux
}

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           newRouter(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("platform-api listening on :8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
