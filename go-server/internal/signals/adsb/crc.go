package adsb

// CRC-24 polynomial used in Mode S (ADS-B)
const crcPolynomial uint32 = 0xFFF409 // 24-bit polynomial
const crcBits = 24
const totalBits = 112
const crcStartBit = totalBits - crcBits

// CheckCRC verifies if the last 24 bits match the CRC of the first 88 bits
func CheckCRC(bits []bool) bool {
	if len(bits) != totalBits {
		return false
	}

	dataBits := bits[:crcStartBit]
	expectedCRC := bitsToInt(bits[crcStartBit:])
	calculatedCRC := computeCRC(dataBits)

	// if calculatedCRC != expectedCRC {
	// 	fmt.Printf("  ❌ Expected CRC: %06X | Got: %06X\n", expectedCRC, calculatedCRC)
	// }

	return calculatedCRC == expectedCRC
}

// computeCRC calculates the 24-bit CRC over the given bits
func computeCRC(bits []bool) int {
	crc := uint32(0)

	for _, bit := range bits {
		msb := (crc >> 23) & 1
		crc <<= 1

		if bit != (msb == 1) {
			crc ^= crcPolynomial
		}
	}

	return int(crc & 0xFFFFFF)
}

// TryFix1BitError attempts to correct a single-bit error in the first 88 bits.
func TryFix1BitError(bits []bool) []bool {
	for i := range bits[:88] {
		flipped := make([]bool, len(bits))
		copy(flipped, bits)
		flipped[i] = !flipped[i]
		if CheckCRC(flipped) {
			return flipped
		}
	}
	return nil
}
