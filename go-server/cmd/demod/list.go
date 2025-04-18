package demod

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available signal demodulators",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available demodulators:")
		fmt.Println("  • adsb")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
