package main

import (
	"encoding/json"
	"github.com/ungerik/go-mavlink"
)

const GlobalPositionInt = 33
const AttitudeInt = 30

type stateHolder struct {
	stateData stateData
	db        DbWrapper
}

type stateData struct {
	GlobalPositionInt *mavlink.GlobalPositionInt
	Attitude          *mavlink.Attitude
	TelemetryData     *TelemetryData
}

func (sd *stateData) updateGP() {
	sd.TelemetryData.SetGlobalPosition(sd.GlobalPositionInt)
}

func (sd *stateData) updateAT() {
	sd.TelemetryData.SetAttitude(sd.Attitude)
}

func (s *stateHolder) startStateHolder(packetChan chan *mavlink.MavPacket, dbFilename string) {
	s.stateData = stateData{}
	s.stateData.TelemetryData = &TelemetryData{}
	s.db = DbWrapper{}
	s.db.initialize(dbFilename)

	var packet *mavlink.MavPacket
	for {
		packet = <-packetChan
		s.processPacket(packet)
	}
}

func (s *stateHolder) processPacket(packet *mavlink.MavPacket) {
	switch packet.Msg.ID() {
	case GlobalPositionInt:
		s.stateData.GlobalPositionInt = packet.Msg.(*mavlink.GlobalPositionInt)
		s.stateData.updateGP()
		s.insertIntoDb(s.stateData.GlobalPositionInt, GlobalPositionInt)
		break
	case AttitudeInt:
		s.stateData.Attitude = packet.Msg.(*mavlink.Attitude)
		s.stateData.updateAT()
		s.insertIntoDb(s.stateData.Attitude, AttitudeInt)
		break
	}
	//TODO add information about last update time
}

func (s *stateHolder) insertIntoDb(object interface{}, dataType int) {
	jsonEncoded, err := json.Marshal(s.stateData.TelemetryData)
	if err != nil {
		panic(err)
	}
	s.db.insert(dataType, string(jsonEncoded))
}
