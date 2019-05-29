package main

import (
	dialect "github.com/gswly/gomavlib/dialects/common"
)

//TODO remove
/*
Not so fast.
Json field names created using mavlink.GlobalPositionInt class have capitalised first letters.
*/
//TODO remove remove

type TelemetryData struct {
	Lat         int32   `json:"lat"`
	Lon         int32   `json:"lon"`
	Alt         int32   `json:"alt"`
	RelativeAlt int32   `json:"relativealt"`
	Vx          int16   `json:"vx"`
	Vy          int16   `json:"vy"`
	Vz          int16   `json:"vz"`
	Hdg         uint16  `json:"hdg"`
	Roll        float32 `json:"roll"`
	Pitch       float32 `json:"pitch"`
	Yaw         float32 `json:"yaw"`
	Rollspeed   float32 `json:"rollspeed"`
	Pitchspeed  float32 `json:"pitchspeed"`
	Yawspeed    float32 `json:"yawspeed"`
}

func (td *TelemetryData) SetGlobalPosition(gps *dialect.MessageGlobalPositionInt) {
	td.Lat = gps.Lat
	td.Lon = gps.Lon
	td.Alt = gps.Alt
	td.RelativeAlt = gps.RelativeAlt
	td.Vx = gps.Vx
	td.Vy = gps.Vy
	td.Vz = gps.Vz
	td.Hdg = gps.Hdg
}

func (td *TelemetryData) SetAttitude(at *dialect.MessageAttitude) {
	td.Roll = at.Roll
	td.Pitch = at.Pitch
	td.Yaw = at.Yaw
	td.Rollspeed = at.Rollspeed
	td.Pitchspeed = at.Pitchspeed
	td.Yawspeed = at.Yawspeed
}
