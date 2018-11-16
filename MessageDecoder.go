package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/ungerik/go-mavlink"
)

var globalPosition = &GlobalPosition{}

func update(packet *mavlink.MavPacket) error {
	switch packet.Header.MessageID {
	case 30:
		return nil
	case 33:
		return updateGlobalPosition(packet)
	}
	return nil
}

func updateGlobalPosition(packet *mavlink.MavPacket) error {
	position, e := decodeGlobalPosition(packet)
	if e != nil {
		return e
	}
	if position.TimeBootMs > globalPosition.TimeBootMs {
		globalPosition = position
	}
	return nil
}

func decodeGlobalPosition(packet *mavlink.MavPacket) (*GlobalPosition, error) {
	if packet.Header.MessageID != 33 {
		return nil, errors.New("message id must be 33")
	}
	body := packet.Bytes()
	if packet.Header.PayloadLength < 28 {
		diff := 28 - packet.Header.PayloadLength
		salt := make([]byte, diff, diff)
		body = append(body, salt...)
	}
	res := GlobalPosition{}
	buff := bytes.NewReader(body)
	err := binary.Read(buff, binary.LittleEndian, &res) //TODO little endian or big endian?
	return &res, err
}
