package adsb

import "fmt"

// ExtractBits takes a magnitude slice (post-preamble) and returns 112 bits based on PPM decoding.
// Each bit is expected to span `samplesPerBit` samples (typically 2 at 2 MHz).
func ExtractBits(mags []float32, samplesPerBit int) []bool {
	numBits := 112
	bits := make([]bool, numBits)

	for i := range numBits {
		offset := i * samplesPerBit

		if offset+samplesPerBit > len(mags) {
			break
		}

		// Pulse position modulation:
		// - Pulse in first half = 1
		// - Pulse in second half = 0
		first := mags[offset]
		second := mags[offset+1]

		bits[i] = first > second
	}

	return bits
}

// BitsToString prints a binary string representation of bits for debugging.
func BitsToString(bits []bool) string {
	s := make([]rune, len(bits))
	for i, b := range bits {
		if b {
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}
	return string(s)
}

var charset = []rune("#ABCDEFGHIJKLMNOPQRSTUVWXYZ#####_###############0123456789######")

func decodeCallsign(bits []bool) string {
	callsign := make([]rune, 8)
	for i := range 8 {
		start := i*6 + 0
		end := start + 6
		if end > len(bits) {
			break
		}
		val := bitsToInt(bits[start:end])
		if val < len(charset) {
			callsign[i] = charset[val]
		} else {
			callsign[i] = '_'
		}
	}
	return string(callsign)
}

func decodeAltitude(bits []bool) int {
	if len(bits) < 52 {
		return -1 // insufficient data
	}

	// Simplified Gillham encoding to 25 ft units
	qBit := bits[47]
	if qBit {
		// 13 bits used when Q=1: bits 40–52 except 47
		n := (bitsToInt(bits[40:47]) << 4) | bitsToInt(bits[48:52])
		return n*25 - 1000
	}
	return -1 // unsupported encoding
}

// bitsToInt converts a slice of bits to an integer (MSB first)
func bitsToInt(bits []bool) int {
	result := 0
	for _, b := range bits {
		result <<= 1
		if b {
			result |= 1
		}
	}
	return result
}

// bitsToHex converts a slice of bits to a hex string
func bitsToHex(bits []bool) string {
	val := bitsToInt(bits)
	return formatHex(val, len(bits)/4)
}

func formatHex(val int, width int) string {
	return fmt.Sprintf("%0*X", width, val)
}
