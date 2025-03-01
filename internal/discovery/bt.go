package discovery

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"tinygo.org/x/bluetooth"
)

func New(uuid string) (*bluetooth.Device, error) {
	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	if err != nil {
		return nil, fmt.Errorf("failed at enabling default bt adapter: %w", err)
	}

	hrId, err := bluetooth.ParseUUID(uuid)
	if err != nil {
		return nil, fmt.Errorf("failed to parse device UUID: %w", err)
	}
	address := bluetooth.Address{
		UUID: hrId,
	}
	device, err := adapter.Connect(address, bluetooth.ConnectionParams{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to H10: %w", err)
	}

	return &device, nil
}

func Scan(duration time.Duration) ([]string, error) {
	adapter := bluetooth.DefaultAdapter
	err := adapter.Enable()
	if err != nil {
		return nil, fmt.Errorf("failed at enabling default bt adapter: %w", err)
	}

	devicesNearby, err := obtainUUIDsOfDevicesNearby(adapter, duration)
	if err != nil {
		return nil, fmt.Errorf("failed at scanning for devices nearby: %w", err)
	}
	keys := make([]string, 0)
	for key := range devicesNearby {
		keys = append(keys, key)
	}
	return keys, nil
}

func obtainUUIDsOfDevicesNearby(adapter *bluetooth.Adapter, duration time.Duration) (map[string]int16, error) {
	mut := sync.Mutex{}
	devices := map[string]int16{}
	timeout := time.Now().Add(duration)

	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		mut.Lock()
		defer mut.Unlock()
		deviceName := fmt.Sprintf("%s:%s", strings.ToLower(device.LocalName()), device.Address.String())
		if device.LocalName() == "" {
			deviceName = fmt.Sprintf("%s:%s", "undefined", device.Address.String())
		}
		if currentDevice, exists := devices[deviceName]; exists {
			if currentDevice > device.RSSI {
				devices[deviceName] = device.RSSI
			}
		} else {
			devices[deviceName] = device.RSSI
		}
		if time.Now().After(timeout) {
			err := adapter.StopScan()
			if err != nil {
				panic(fmt.Errorf("failed at stopping scan for devices: %w", err))
			}
		}
	})
	if err != nil {
		return nil, fmt.Errorf("failed to perform scan for device: %w", err)
	}

	return devices, nil
}
