package cardio

func ExtractHeartbeats(ecgSamples []int) ([]Heartbeat, int) {
	if len(ecgSamples) < 50 {
		return []Heartbeat{}, -1
	}
	ecgSamplesDerivatives := make([]int, len(ecgSamples))
	for i := 0; i < len(ecgSamples)-1; i++ {
		ecgSamplesDerivatives[i] = ecgSamples[i+1] - ecgSamples[i]
	}
	heartbeatsRaw := make([]int, 0)
	for i := 0; i < len(ecgSamplesDerivatives); i++ {
		if ecgSamplesDerivatives[i] < -300 {
			heartbeatsRaw = append(heartbeatsRaw, i)
		}
	}
	if len(heartbeatsRaw) < 3 {
		return []Heartbeat{}, -1
	}
	heartbeats := make([]Heartbeat, 0)
	heartbeat := Heartbeat{}
	r := heartbeatsRaw[0]
	for i := 0; i < len(heartbeatsRaw); i++ {
		if i > 0 {
			if heartbeatsRaw[i]-heartbeatsRaw[i-1] < 3 {
			} else {
				heartbeat.Samples = ecgSamples[r : heartbeatsRaw[i]+1]
				heartbeats = append(heartbeats, heartbeat)
				r = heartbeatsRaw[i]
			}
		}
	}
	cursor := 0
	if len(heartbeatsRaw) > 0 {
		cursor = heartbeatsRaw[len(heartbeatsRaw)-1]
	}
	if len(heartbeatsRaw) > 2 {
		cursor = heartbeatsRaw[len(heartbeatsRaw)-2]
	}
	return heartbeats, cursor
}
