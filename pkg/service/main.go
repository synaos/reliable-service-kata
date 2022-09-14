package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var hostname string

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.Debug("Determining hostname")
	h, err := os.Hostname()
	if err != nil {
		log.Panicf("Cannot determine the actual hostname: %v", err)
	} else {
		hostname = h
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleAnyRoot)
	r.HandleFunc("/crash", handleAnyCrash)

	address := os.Getenv("KATA_ADDRESS")
	if address == "" {
		address = "127.0.0.1:8000"
	}
	srv := &http.Server{
		Handler: r,
		Addr:    address,
	}
	log.Infof("Binding to address %v", address)
	log.Fatal(srv.ListenAndServe())
}

func handleAnyRoot(w http.ResponseWriter, r *http.Request) {
	log.WithField("remote_address", r.RemoteAddr).
		Info("Got / request")
	_, _ = fmt.Fprintf(w, "Hello world from %s", hostname)
}

func handleAnyCrash(w http.ResponseWriter, r *http.Request) {
	log.WithField("remote_address", r.RemoteAddr).
		Warn("Got /crash request, will be happening in 1s...")
	_, _ = fmt.Fprint(w, "I'm going to crash now...")

	go crashMe()
}

func crashMe() {
	time.Sleep(1 * time.Second)
	log.Fatal("Going to crash, bye! :-(")
}
