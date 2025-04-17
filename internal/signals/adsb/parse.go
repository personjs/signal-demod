package adsb

import (
	"github.com/personjs/signal-demod/internal/models"
)

// ParseMessage extracts structured ADS-B info from raw 112 bits.
// Currently supports Downlink Format 17 (Extended Squitter).
func ParseMessage(bits []bool) *models.ADSBMessage {
	if len(bits) != 112 {
		return nil
	}

	df := bitsToInt(bits[0:5])
	if df != 17 {
		// Only handling DF17 for now
		return nil
	}

	icao := bitsToHex(bits[8:32])
	tc := bitsToInt(bits[32:37])

	msg := &models.ADSBMessage{
		DownlinkFormat: df,
		ICAO:           icao,
		TypeCode:       tc,
	}

	switch {
	case tc >= 1 && tc <= 4:
		// Aircraft Identification
		msg.MessageType = "Identification"
		msg.Callsign = decodeCallsign(bits[40:88])

	case tc >= 9 && tc <= 18:
		// Airborne position
		msg.MessageType = "Airborne Position"
		msg.NICSupplement = bitsToInt(bits[37:40])
		msg.Altitude = decodeAltitude(bits)
		msg.LatCpr = bitsToInt(bits[54:71])
		msg.LonCpr = bitsToInt(bits[71:88])

	default:
		msg.MessageType = "Unknown"
	}

	return msg
}
