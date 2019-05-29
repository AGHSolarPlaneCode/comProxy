package main

import (
	"fmt"
	"github.com/gswly/gomavlib"
	"github.com/gswly/gomavlib/dialects/ardupilotmega"
	"strconv"
)

func StartReader(portName string, baudRate int, outChan chan *gomavlib.EventFrame) {

	node, err := gomavlib.NewNode(gomavlib.NodeConf{
		Endpoints: []gomavlib.EndpointConf{
			gomavlib.EndpointSerial{Address: portName + ":" + strconv.Itoa(baudRate)},
		},
		Dialect:     ardupilotmega.Dialect,
		OutSystemId: 10,
	})
	if err != nil {
		panic(err)
	}
	defer node.Close()

	for {
		for evt := range node.Events() {
			if frm, ok := evt.(*gomavlib.EventFrame); ok {
				fmt.Printf("received: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
				outChan <- frm
			}
		}
	}
}
