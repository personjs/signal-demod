package adsb

import (
	"math"
	"time"

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

		latCpr := bitsToInt(bits[54:71])
		lonCpr := bitsToInt(bits[71:88])
		isOdd := bits[53]

		now := time.Now().UnixMilli() // nanoseconds for higher precision

		// Cache it
		current := cprFrame{
			Timestamp: now,
			LatCpr:    latCpr,
			LonCpr:    lonCpr,
			IsOdd:     isOdd,
		}

		cache := cprCache[msg.ICAO]

		if isOdd {
			cache[1] = current
		} else {
			cache[0] = current
		}

		cprCache[msg.ICAO] = cache

		even := cache[0]
		odd := cache[1]

		if even.Timestamp > 0 && odd.Timestamp > 0 {
			lat, lon, ok := DecodeCPR(
				even.LatCpr, even.LonCpr,
				odd.LatCpr, odd.LonCpr,
				even.Timestamp, odd.Timestamp,
				false, // surface = false for aircraft
			)
			if ok {
				msg.Latitude = lat
				msg.Longitude = lon
			}
		}
	case tc == 19:
		msg.MessageType = "Airborne Velocity"
		msg.Speed, msg.Heading = decodeVelocity(bits)
	default:
		msg.MessageType = "Unknown"
	}

	return msg
}

func decodeVelocity(bits []bool) (speed float64, heading float64) {
	subtype := bitsToInt(bits[37:40])

	if subtype != 1 {
		return 0, 0 // unsupported subtype for now
	}

	// Ground speed and track angle
	ewDir := bits[45]
	ewVel := bitsToInt(bits[46:56]) // 10 bits
	nsDir := bits[56]
	nsVel := bitsToInt(bits[57:67]) // 10 bits

	if ewVel == 0 && nsVel == 0 {
		return 0, 0
	}

	vEw := float64(ewVel)
	if ewDir {
		vEw *= -1
	}
	vNs := float64(nsVel)
	if nsDir {
		vNs *= -1
	}

	// Speed (knots), Heading (degrees)
	speed = math.Sqrt(vEw*vEw + vNs*vNs)
	heading = math.Mod(math.Atan2(vEw, vNs)*180/math.Pi, 360)

	return speed, heading
}
