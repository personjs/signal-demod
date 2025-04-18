package demod

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "signal-demod",
	Short: "RTL-SDR signal demodulation CLI",
	Long:  "A modular RTL-SDR demodulator that supports multiple signal types like ADS-B.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌", err)
		os.Exit(1)
	}
}
