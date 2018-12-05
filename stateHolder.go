package main

import "github.com/ungerik/go-mavlink"

const GlobalPositionInt = 33
const Attitude = 30

type stateHolder struct {
	stateData stateData
}

type stateData struct {
	GlobalPositionInt *mavlink.GlobalPositionInt
	Attitude          *mavlink.Attitude
}

func (s *stateHolder) startStateHolder(packetChan chan *mavlink.MavPacket) {
	s.stateData = stateData{}
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
		break
	case Attitude:
		stateData.Attitude = packet.Msg.(*mavlink.Attitude)
		break
	}

	//TODO save in database
	//TODO add information about last update time
}
