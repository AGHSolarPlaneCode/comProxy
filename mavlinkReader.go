package main

import (
	"github.com/jacobsa/go-serial/serial"
	"github.com/ungerik/go-mavlink"
	"io"
	"log"
)

const frameSize = 1
const frameStart = 0xFD

func StartReader(portName string, baudRate uint, outChan chan *mavlink.MavPacket) {

	options1 := serial.OpenOptions{
		PortName:        portName,
		BaudRate:        baudRate,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options1)

	if err != nil {
		log.Fatal(err)
	}

	defer port.Close()

	parser := mavlink.GetMavParser()

	bytes := make([]byte, frameSize)

	startFound := false
	for {
		n := read(port, &bytes)
		if n > 0 {
			for _, b := range bytes[:n] {
				if !startFound {
					if b == frameStart {
						startFound = true
					}
				}
				var packet *mavlink.MavPacket
				var err error
				if startFound {
					packet, err = parser(b)
				}
				if err != nil {
				//	log.Fatal(err)

				} else if packet != nil {
					outChan <- packet
					startFound = false
				}
			}
		}
	}
}

func read(port io.ReadWriteCloser, bytes *[]byte) int {
	n, err := port.Read(*bytes)

	if err != nil {
		log.Fatal(err)
	}

	return n
}
