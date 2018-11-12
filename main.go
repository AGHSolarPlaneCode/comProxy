package main

import (
	"github.com/jacobsa/go-serial/serial"
	"log"
	"os"
	"strconv"
	"io"
	"github.com/ungerik/go-mavlink"
)

func main() {
	if len(os.Args) < 3 {
		//println("Usage: comProxy port1 port2 baudRate")
		println("Usage: comProxy portName baudRate")
		return
	}

	port1Name := os.Args[1]
	//port2Name := os.Args[2]
	baudRate, err := strconv.ParseUint(os.Args[2], 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	options1 := serial.OpenOptions{
		PortName:        port1Name,
		BaudRate:        uint(baudRate),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	//options2 := serial.OpenOptions{
	//	PortName:        port2Name,
	//	BaudRate:        uint(baudRate),
	//	DataBits:        8,
	//	StopBits:        1,
	//	MinimumReadSize: 4,
	//}

	port1, err := serial.Open(options1)

	if err != nil {
		log.Fatal(err)
	}

	defer port1.Close()

	//port2, err := serial.Open(options2)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer port2.Close()

	parser := mavlink.GetMavParser()

	bytes := make([]byte, 1)

	//for {
	//	n := read(port1, &bytes)
	//	if n > 0 {
	//		write(port2, &bytes, n)
	//	}
	//
	//	n = read(port2, &bytes)
	//	if n > 0 {
	//		write(port1, &bytes, n)
	//	}
	//}

	for {
		n := read(port1, &bytes)
		if n > 0 {
			for _, b := range bytes[:n] {
				packet, err := parser(b)
				if err != nil {
					log.Fatal(err)

				} else if packet != nil {
					processPacket(packet)
				}
			}
		}
	}
}

func processPacket(packet *mavlink.MavPacket) {
	println(packet.Header.MessageID)
}

func read(port io.ReadWriteCloser, bytes *[]byte) int {
	n, err := port.Read(*bytes)

	if err != nil {
		log.Fatal(err)
	}

	return n
}

func write(port io.ReadWriteCloser, pbytes *[]byte, n int) {
	bytes := *pbytes
	n, err := port.Write(bytes[:n])

	if err != nil {
		log.Fatal(err)
	}
}
