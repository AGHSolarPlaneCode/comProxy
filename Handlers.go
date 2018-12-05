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
	http.HandleFunc("/attitude", getAttitude)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getGps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(state.GlobalPositionInt); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getAttitude(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(state.Attitude); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
