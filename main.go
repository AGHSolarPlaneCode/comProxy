package main

import (
	"github.com/ungerik/go-mavlink"
	"log"
	"os"
	"strconv"
	"time"
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

	dbFilename := time.Now().Format("2006.01.02 15:04:05") + ".sql"

	if len(os.Args) >= 4 {
		dbFilename = os.Args[3]
	}

	packetChannel := make(chan *mavlink.MavPacket)
	go StartReader(os.Args[1], uint(baudRate), packetChannel)

	stateHolder := stateHolder{}
	go stateHolder.startStateHolder(packetChannel, dbFilename)

	startHttpServer(&stateHolder.stateData)
}
