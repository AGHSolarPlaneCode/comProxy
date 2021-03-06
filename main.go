package main

import (
	"github.com/gswly/gomavlib"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		//println("Usage: comProxy port1 port2 baudRate")
		println("Usage: comProxy portName baudRate [dbFilename]")
		return
	}

	baudRate, err := strconv.ParseUint(os.Args[2], 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	dbFilename := "flightData.sql"
	if len(os.Args) >= 4 {
		dbFilename = os.Args[3]
	}

	packetChannel := make(chan *gomavlib.EventFrame)
	go StartReader(os.Args[1], int(baudRate), packetChannel)

	stateHolder := stateHolder{}
	go stateHolder.startStateHolder(packetChannel, dbFilename)

	startHttpServer(&stateHolder.stateData)
}
