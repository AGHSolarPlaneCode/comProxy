package main

type GlobalPosition struct {
	TimeBootMs  uint32 `json:"timebootms"`
	Lat         int32  `json:"lat"`
	Lon         int32  `json:"lon"`
	Alt         int32  `json:"alt"`
	RelativeAlt int32  `json:"relativealt"`
	Vx          int16  `json:"vx"`
	Vy          int16  `json:"vy"`
	Vz          int16  `json:"vz"`
	Hdg         uint16 `json:"hdg"`
}
