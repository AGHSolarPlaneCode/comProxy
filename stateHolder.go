package main

import (
	"encoding/json"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/common"
)

const GlobalPositionInt = 33
const AttitudeInt = 30

type stateHolder struct {
	stateData stateData
	db        DbWrapper
}

type stateData struct {
	GlobalPosition *common.MessageGlobalPositionInt
	Attitude       *common.MessageAttitude
	TelemetryData  *TelemetryData
}

func (s *stateHolder) startStateHolder(packetChan chan *gomavlib.EventFrame, dbFilename string) {
	s.stateData = stateData{}
	s.db = DbWrapper{}
	s.db.initialize(dbFilename)

	var packet *gomavlib.EventFrame
	for {
		packet = <-packetChan
		s.processPacket(packet)
	}
}

func (s *stateHolder) processPacket(packet *gomavlib.EventFrame) {
	if gps, ok := packet.Message().(*common.MessageGlobalPositionInt); ok {
		s.stateData.GlobalPosition = gps
		s.stateData.TelemetryData.SetGlobalPosition(gps)
		s.insertIntoDb(s.stateData.GlobalPosition, GlobalPositionInt)
	}

	if att, ok := packet.Message().(*common.MessageAttitude); ok {
		s.stateData.Attitude = att
		s.stateData.TelemetryData.SetAttitude(att)
		s.insertIntoDb(s.stateData.Attitude, AttitudeInt)
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
