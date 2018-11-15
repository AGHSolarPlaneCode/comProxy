package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/ungerik/go-mavlink"
)

func decodeGlobalPosition(packet *mavlink.MavPacket) (*mavlink.GlobalPositionInt, error) {
	if packet.Header.MessageID != 33 {
		return nil, errors.New("message id must be 33")
	}
	body := packet.Bytes()
	if packet.Header.PayloadLength < 28 {
		diff := 28 - packet.Header.PayloadLength
		salt := make([]byte, diff, diff)
		body = append(body, salt...)
	}
	res := mavlink.GlobalPositionInt{}
	buff := bytes.NewReader(body)
	err := binary.Read(buff, binary.LittleEndian, &res) //TODO little endian or big endian?
	return &res, err
}
