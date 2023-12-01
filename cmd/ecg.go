package cmd

import (
	"github.com/spf13/cobra"
)

// ecgCmd represents the ecg command
var ecgCmd = &cobra.Command{
	Use:   "ecg",
	Short: "ECG commands",
}

func init() {
	rootCmd.AddCommand(ecgCmd)
	ecgCmd.AddCommand(ecgStreamCmd)
	ecgStreamCmd.Flags().StringVarP(&UUID, "device-uuid", "u", "", "UUID of Bluetooth device.")
	_ = ecgStreamCmd.MarkFlagRequired("device-uuid")
}
