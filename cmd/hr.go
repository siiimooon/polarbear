package cmd

import (
	"github.com/spf13/cobra"
)

// hrCmd represents the hr command
var hrCmd = &cobra.Command{
	Use:   "hr",
	Short: "HR commands",
}

func init() {
	rootCmd.AddCommand(hrCmd)
	hrCmd.AddCommand(hrStreamCmd)
	hrStreamCmd.Flags().StringVarP(&UUID, "device-uuid", "u", "", "UUID of Bluetooth device.")
	_ = hrStreamCmd.MarkFlagRequired("device-uuid")
}
