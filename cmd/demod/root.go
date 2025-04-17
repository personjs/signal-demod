package demod

import (
	"fmt"
	"os"

	"github.com/personjs/signal-demod/internal/config"
	"github.com/personjs/signal-demod/internal/services"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "signal-demod",
	Short: "RTL-SDR signal demodulation CLI",
	Long:  "A modular RTL-SDR demodulator that supports multiple signal types like ADS-B.",
}

func Execute() {
	if err := config.Load(); err != nil {
		fmt.Println("failed to load config:", err)
		os.Exit(1)
	}

	services.InitLogger(config.App.Log)
	services.InitDatabase(config.App.DB)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
