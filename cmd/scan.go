package cmd

import (
	"fmt"
	"github.com/siiimooon/polarbear/internal/discovery"
	"github.com/spf13/cobra"
	"time"
)

var scanDuration time.Duration

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for Bluetooth devices nearby",
	Run: func(cmd *cobra.Command, args []string) {
		results, err := discovery.Scan(scanDuration)
		if err != nil {
			panic(fmt.Errorf("failed at scanning for bluetooth devices nearby: %w", err))
		}
		for _, result := range results {
			fmt.Println(result)
		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().DurationVarP(&scanDuration, "duration", "d", 10*time.Second, "Duration of scan")
}
