package sdr

import (
	rtlsdr "github.com/jpoirier/gortlsdr"
	"github.com/personjs/signal-demod/internal/services"
)

func Start(out chan<- complex64) {
	dev, err := rtlsdr.Open(0)
	if err != nil {
		services.Logger.Fatal().Err(err).Msg("failed to open rtl-sdr")
	}
	defer dev.Close()

	if err := dev.SetCenterFreq(1090000000); err != nil {
		services.Logger.Fatal().Err(err).Msg("failed to set center freq")
	}
	if err := dev.SetSampleRate(2000000); err != nil {
		services.Logger.Fatal().Err(err).Msg("failed to set sample rate")
	}
	if err := dev.SetTunerGainMode(true); err != nil { // auto gain
		services.Logger.Error().Err(err).Msg("failed to set auto gain")
	}
	if err := dev.SetTunerGain(490); err != nil { // 49.0 dB
		services.Logger.Fatal().Err(err).Msg("failed to set tuner gain")
	}
	if err := dev.ResetBuffer(); err != nil {
		services.Logger.Fatal().Err(err).Msg("failed to reset buffer")
	}

	buf := make([]byte, 16*16384) // I/Q interleaved, 8-bit unsigned
	for {
		n, err := dev.ReadSync(buf, len(buf))
		if err != nil || n == 0 {
			services.Logger.Error().Err(err).Msg("ReadSync failure")
			continue
		}

		// Convert each I/Q pair into complex64 (centered at 0)
		for i := 0; i < n; i += 2 {
			iSample := float32(buf[i]) - 127.5
			qSample := float32(buf[i+1]) - 127.5
			out <- complex(iSample, qSample)
		}
	}
}
