package main

import (
	"github.com/ungerik/go-mavlink"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		//println("Usage: comProxy port1 port2 baudRate")
		println("Usage: comProxy portName baudRate")
		return
	}

	baudRate, err := strconv.ParseUint(os.Args[2], 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	packetChannel := make(chan *mavlink.MavPacket)
	go StartReader(os.Args[1], uint(baudRate), packetChannel)

	stateHolder := stateHolder{}
	go stateHolder.startStateHolder(packetChannel)

	startHttpServer(&stateHolder.stateData)
}
