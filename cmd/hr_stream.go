package cmd

import (
	"context"
	"fmt"
	"github.com/siiimooon/go-polar/pkg/h10"
	"github.com/siiimooon/polarbear/internal/discovery"

	"github.com/spf13/cobra"
)

// hrStreamCmd represents the hr stream command
var hrStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream HR",
	Run: func(cmd *cobra.Command, args []string) {
		device, err := discovery.New(UUID)
		if err != nil {
			panic(fmt.Errorf("failed at establishing connection to device: %w", err))
		}
		sensorReader := h10.New(device)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		sink := make(chan h10.HeartRateMeasurement, 1)
		go func() {
			err := sensorReader.StreamHeartRate(ctx, sink)
			if err != nil {
				panic(fmt.Sprintf("failed at streaming hr: %v", err))
			}
		}()

		for {
			select {
			case measurement := <-sink:
				fmt.Println(measurement.String())
			}
		}
	},
}
