package models

import (
	"fmt"

	"gorm.io/gorm"
)

type ADSBMessage struct {
	gorm.Model

	DownlinkFormat int
	ICAO           string
	TypeCode       int
	MessageType    string
	Callsign       string
	Altitude       int
	NICSupplement  int
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	Heading        float64
	Speed          float64
}

func (m *ADSBMessage) String() string {
	return fmt.Sprintf(
		"DF:%d | ICAO:%s | TC:%d | Type:%s | Callsign:%s | Alt:%d | Lat:%f | Lon:%f  | Hdg:%f | Spd:%f",
		m.DownlinkFormat,
		m.ICAO,
		m.TypeCode,
		m.MessageType,
		m.Callsign,
		m.Altitude,
		m.Latitude,
		m.Longitude,
		m.Heading,
		m.Speed,
	)
}
