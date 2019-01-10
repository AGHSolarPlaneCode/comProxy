package main

import "github.com/ungerik/go-mavlink"

const GlobalPositionInt = 33
const AttitudeInt = 30

type stateHolder struct {
	stateData stateData
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

func (s *stateHolder) startStateHolder(packetChan chan *mavlink.MavPacket) {
	s.stateData = stateData{}
	s.stateData.TelemetryData = &TelemetryData{}
	var packet *mavlink.MavPacket
	for {
		packet = <-packetChan
		processPacket(packet, &s.stateData)
	}
}

func processPacket(packet *mavlink.MavPacket, stateData *stateData) {
	switch packet.Msg.ID() {
	case GlobalPositionInt:
		stateData.GlobalPositionInt = packet.Msg.(*mavlink.GlobalPositionInt)
		stateData.updateGP()
		break
	case AttitudeInt:
		stateData.Attitude = packet.Msg.(*mavlink.Attitude)
		stateData.updateAT()
		break
	}

	//TODO save in database
	//TODO add information about last update time
}
