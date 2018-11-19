package main

import "github.com/ungerik/go-mavlink"

const GlobalPositionInt = 33

type stateHolder struct {
	stateData stateData
}

type stateData struct {
	GlobalPositionInt *mavlink.GlobalPositionInt
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
	}

	//TODO save in database
	//TODO add information about last update time
}
