package adsb

import (
	"encoding/json"

	"github.com/personjs/signal-demod/internal/sdr"
	"github.com/personjs/signal-demod/internal/services"
	"github.com/personjs/signal-demod/internal/websocket"
	"github.com/personjs/signal-demod/internal/models"
)

const (
	sampleRate     = 2000000 // 2 MHz
	samplesPerBit  = 2       // ADS-B 1µs per bit
	messageBits    = 112
	preambleLength = 16 // 8µs preamble * 2 samples/µs
	totalLength    = preambleLength + messageBits*samplesPerBit
)

func StartSDR(samples chan<- complex64) {
	go sdr.Start(samples)
}

func Run(samples <-chan complex64, hub *websocket.Hub) {
	buffer := make([]complex64, 0, totalLength)

	var total, valid, corrected int

	for sample := range samples {
		buffer = append(buffer, sample)
		if len(buffer) > totalLength {
			buffer = buffer[1:]
		}

		if len(buffer) == totalLength {
			mags := make([]float32, len(buffer))
			for i, c := range buffer {
				mags[i] = real(c)*real(c) + imag(c)*imag(c)
			}

			if HasPreamble(mags[:preambleLength]) {
				bits := ExtractBits(mags[preambleLength:], samplesPerBit)

				total++

				// Run initial CRC check
				if !CheckCRC(bits) {
					if fixed := TryFix1BitError(bits); fixed != nil {
						bits = fixed
						corrected++
					} else {
						// Skip this message, CRC invalid
						continue
					}
				}

				valid++
				msg := ParseMessage(bits)
				if msg != nil {
					// Holy shit I got one!
					services.Logger.Info().Msgf("✈️ %s", msg)
					services.DB.Create(msg)

					// Convert to plane and broadcast
					plane := models.ToPlane(msg)
					data, _ := json.Marshal(plane)
					hub.Broadcast(data)
				}

				// Print stats every 50 messages
				if total%50 == 0 {
					services.Logger.Info().Msgf("📈 Stats — Valid: %d / %d (%.1f%%), Corrected: %d", valid, total, float64(valid)/float64(total)*100, corrected)
				}

				// Prevent re-processing overlapping samples
				buffer = buffer[totalLength/2:]
			}
		}
	}
}
