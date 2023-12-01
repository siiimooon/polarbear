package cmd

import (
	"context"
	"fmt"
	"github.com/siiimooon/go-polar/pkg/h10"
	"github.com/siiimooon/polarbear/internal/cardio"
	"github.com/siiimooon/polarbear/internal/discovery"
	"github.com/spf13/cobra"
)

// ecgStreamCmd represents the ecg stream command
var ecgStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream ECG",
	Run: func(cmd *cobra.Command, args []string) {
		device, err := discovery.New(UUID)
		if err != nil {
			panic(fmt.Errorf("failed at establishing connection to device: %w", err))
		}
		sensorReader := h10.New(device)
		sink := make(chan h10.ECGMeasurement, 1)
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() {
			err := sensorReader.StreamECG(ctx, sink)
			if err != nil {
				panic(fmt.Sprintf("error occured during ecg streaming from Polar H10: %v", err))
			}
		}()

		aggregatedSamples := make([]int, 0)
		for {
			select {
			case measurement := <-sink:
				aggregatedSamples = append(aggregatedSamples, measurement.GetSamples()...)
				heartbeats, cursor := cardio.ExtractHeartbeats(aggregatedSamples)
				if cursor > 0 {
					printHeartbeats(heartbeats)
					aggregatedSamples = aggregatedSamples[cursor-1:]
				}
			}
		}
	},
}

func printHeartbeats(heartbeats []cardio.Heartbeat) {
	for _, heartbeat := range heartbeats {
		cleanSampleRanges := filterVectorNoise(100, compactSamples(heartbeat.Samples))
		fmt.Printf("%v - %v\n", len(cleanSampleRanges), cleanSampleRanges)
	}
}

func compactSamples(samples []int) [][]int {
	if len(samples) < 2 {
		return [][]int{samples}
	}
	compactedSamples := make([][]int, 0)
	currentIsBiggest := samples[1] > samples[0]
	start := samples[0]
	for i := 1; i < len(samples); i++ {
		currentValue := samples[i]
		previousValue := samples[i-1]
		if currentValue == previousValue {
			continue
		} else if currentIsBiggest && (currentValue < previousValue) {
			currentIsBiggest = !currentIsBiggest
			compactedSamples = append(compactedSamples, []int{start, previousValue})
			start = previousValue
		} else if !currentIsBiggest && (currentValue > previousValue) {
			currentIsBiggest = !currentIsBiggest
			compactedSamples = append(compactedSamples, []int{start, previousValue})
			start = previousValue
		}
		if i == len(samples)-1 {
			compactedSamples = append(compactedSamples, []int{start, currentValue})
		}
	}
	return compactedSamples
}

func filterVectorNoise(filter int, vectors [][]int) (filtered [][]int) {
	filtered = make([][]int, 0)

	abs := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	for _, vector := range vectors {
		if abs(vector[1]-vector[0]) > filter {
			filtered = append(filtered, vector)
		}
	}
	return
}
