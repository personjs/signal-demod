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

	// Identification
	Callsign string

	// Position
	Altitude      int
	NICSupplement int
	LatCpr        int
	LonCpr        int
}

func (m *ADSBMessage) String() string {
	return fmt.Sprintf(
		"DF:%d | ICAO:%s | TC:%d | Type:%s | Callsign:%s | Alt:%d | CPR:(%d,%d)",
		m.DownlinkFormat,
		m.ICAO,
		m.TypeCode,
		m.MessageType,
		m.Callsign,
		m.Altitude,
		m.LatCpr,
		m.LonCpr,
	)
}
