package adsb

import (
	"github.com/personjs/signal-demod/internal/services"
)

// HasPreamble checks for the presence of a Mode S preamble in 16 signal magnitude samples
func HasPreamble(mags []float32) bool {
	if len(mags) < 16 {
		return false
	}

	// Safely compute noise floor from trailing samples
	var noiseSamples = 20
	if len(mags) < noiseSamples {
		noiseSamples = len(mags)
	}

	var sum float32
	for _, v := range mags[len(mags)-noiseSamples:] {
		sum += v
	}
	noise := sum / float32(noiseSamples)
	threshold := noise*2.0 + 1.0

	services.Logger.Trace().Msgf("🔍 Noise floor: %.2f | Threshold: %.2f", noise, threshold)
	services.Logger.Trace().Msgf("🔬 Preamble window: %v", mags[:16])

	// --- Pulse pattern pass ---
	pulsePattern := mags[0] > threshold &&
		mags[1] < noise &&
		mags[2] > threshold &&
		mags[3] < noise &&
		mags[7] > threshold &&
		mags[8] > threshold &&
		mags[9] > threshold &&
		mags[11] > threshold &&
		mags[12] > threshold &&
		mags[13] > threshold

	// --- Heuristic 1: total energy ---
	var totalEnergy float32
	for _, v := range mags[:16] {
		totalEnergy += v
	}

	// --- Heuristic 2: shape deltas ---
	delta1 := mags[0] - mags[1]
	delta2 := mags[2] - mags[3]

	services.Logger.Trace().Msgf("⚡️ Total energy: %.2f | Δ0-1: %.2f | Δ2-3: %.2f", totalEnergy, delta1, delta2)

	// --- Heuristic 2: expected pulses should be significantly higher than neighbors ---
	if (mags[0]-mags[1] < noise) || (mags[2]-mags[3] < noise) {
		return false
	}

	// --- Final check ---
	if pulsePattern &&
		totalEnergy > 25.0 &&
		delta1 > 1.0 &&
		delta2 > 1.0 {
		services.Logger.Debug().Msg("📡 PREAMBLE DETECTED ✅")
		return true
	}

	return true
}
