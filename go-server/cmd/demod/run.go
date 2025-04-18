package demod

import (
	"fmt"
	"os"
	"strings"

	"github.com/personjs/signal-demod/internal/signals/adsb"
	"github.com/personjs/signal-demod/internal/websocket"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <signal>",
	Short: "Run a signal demodulator (e.g., adsb)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		signal := strings.ToLower(args[0])
		switch signal {
		case "adsb":
			// Websocket
			hub := websocket.NewHub()
			go hub.Run()
			go websocket.Start(":8081", hub)

			// Signal
			samples := make(chan complex64, 8192)
			go adsb.StartSDR(samples)
			adsb.Run(samples, hub)
		default:
			fmt.Printf("❌ Unknown signal: %s\n", signal)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
