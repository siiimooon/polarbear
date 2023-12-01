package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var UUID string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "polarbear",
	Short: "A CLI for interacting with a Polar H10 nearby you",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
