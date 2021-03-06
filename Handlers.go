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
	http.HandleFunc("/currTele", getCurrentTelemetry)
	http.HandleFunc("/position", getGlobalPosition)
	http.HandleFunc("/positionRaw", getRawGps)
	http.HandleFunc("/attitude", getAttitude)
	http.HandleFunc("/hud", getHud)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getGps(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.GpsData); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
}

func getAttitude(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.Attitude); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
}

func getCurrentTelemetry(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.TelemetryData); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//w.WriteHeader(http.StatusOK)
}

func getGlobalPosition(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.GlobalPosition); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getRawGps(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.GpsRaw); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getHud(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(state.HudData); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
