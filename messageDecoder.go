package main

import (
	"github.com/ungerik/go-mavlink"
	"log"
)

//TODO remove

var globalPosition = &mavlink.GlobalPositionInt{}

func update(packet *mavlink.MavPacket) {
	switch packet.Header.MessageID {
	case 30:
		return
		break
	case 33:
		updateGlobalPosition(packet)
		break
	}
	return
}

func updateGlobalPosition(packet *mavlink.MavPacket) {
	position := decodeGlobalPosition(packet)
	if position == nil {
		return
	}
	if position.TimeBootMs > globalPosition.TimeBootMs {
		globalPosition = position
	}
}

func decodeGlobalPosition(packet *mavlink.MavPacket) *mavlink.GlobalPositionInt {
	messege, ok := packet.Msg.(*mavlink.GlobalPositionInt)
	if !ok {
		log.Fatal("failed to cast packet.Msg to mavlink.GlobalPositionInt.")
		return nil
	}
	return messege
}
