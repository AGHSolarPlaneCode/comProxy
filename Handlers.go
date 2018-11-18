package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var state *stateData

func startHttpServer(data *stateData) {
	state = data
	http.HandleFunc("/gps", getGps)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getGps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(state.GlobalPositionInt); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getAttitude(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
