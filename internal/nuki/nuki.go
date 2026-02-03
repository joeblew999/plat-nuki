// Package nuki provides BLE communication with Nuki smart locks.
package nuki

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"tinygo.org/x/bluetooth"

	"github.com/nuki-io/nuki-cli/pkg/blecommands"
	"github.com/nuki-io/nuki-cli/pkg/bleflows"
	"github.com/nuki-io/nuki-cli/pkg/nukible"
)

// Status represents the current state of a Nuki lock.
type Status struct {
	LockState       string
	NukiState       string
	BatteryCritical bool
	BatteryPercent  int
}

// Client provides operations for interacting with a Nuki lock over BLE.
type Client struct {
	deviceAddr string
}

// NewClient creates a new Nuki BLE client.
// If addr is empty, it attempts to load the address from saved credentials.
func NewClient(addr string) (*Client, error) {
	if addr == "" {
		addr = getDeviceAddrFromConfig()
		if addr == "" {
			return nil, fmt.Errorf("no device address specified and no paired devices found")
		}
	}
	return &Client{deviceAddr: addr}, nil
}

// DeviceAddr returns the configured device address.
func (c *Client) DeviceAddr() string {
	return c.deviceAddr
}

// getDeviceAddrFromConfig loads the device address from saved credentials.
func getDeviceAddrFromConfig() string {
	viper.SetConfigName("nuki-credentials")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	auths := viper.GetStringMap("authorizations")
	for id := range auths {
		return id
	}
	return ""
}

// ScanResult represents a discovered Nuki device.
type ScanResult struct {
	Address string
	Name    string
	RSSI    int16
}

// Scan searches for nearby Nuki devices.
func Scan(duration time.Duration) ([]ScanResult, error) {
	_, err := nukible.NewNukiBle()
	if err != nil {
		return nil, fmt.Errorf("failed to enable Bluetooth: %w", err)
	}

	var results []ScanResult
	devices := make(map[string]bool)

	adapter := bluetooth.DefaultAdapter
	timeout := time.AfterFunc(duration, func() { adapter.StopScan() })

	err = adapter.Scan(func(a *bluetooth.Adapter, sr bluetooth.ScanResult) {
		if !strings.HasPrefix(sr.LocalName(), "Nuki") {
			return
		}
		addr := sr.Address.String()
		if !devices[addr] {
			devices[addr] = true
			results = append(results, ScanResult{
				Address: addr,
				Name:    sr.LocalName(),
				RSSI:    sr.RSSI,
			})
		}
	})

	timeout.Stop()

	if err != nil {
		return results, fmt.Errorf("scan error: %w", err)
	}

	return results, nil
}

// Pair pairs with a Nuki device and saves credentials.
func Pair(addr, pin string) error {
	ble, err := nukible.NewNukiBle()
	if err != nil {
		return fmt.Errorf("failed to enable Bluetooth: %w", err)
	}

	err = ble.ScanForDevice(addr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	flow, err := bleflows.NewUnauthenticatedFlow(ble, addr)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer flow.DisconnectDevice()

	err = flow.Authorize(pin)
	if err != nil {
		return fmt.Errorf("pairing failed: %w", err)
	}

	err = viper.WriteConfigAs("nuki-credentials.yaml")
	if err != nil {
		return fmt.Errorf("could not save credentials: %w", err)
	}

	return nil
}

// GetStatus retrieves the current lock status.
func (c *Client) GetStatus() (*Status, error) {
	ble, err := nukible.NewNukiBle()
	if err != nil {
		return nil, fmt.Errorf("Bluetooth error: %w", err)
	}

	err = ble.ScanForDevice(c.deviceAddr, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	flow, err := bleflows.NewAuthenticatedFlow(ble, c.deviceAddr)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}
	defer flow.DisconnectDevice()

	state, err := flow.GetStatus()
	if err != nil {
		return nil, fmt.Errorf("status failed: %w", err)
	}

	return &Status{
		LockState:       state.LockState.String(),
		NukiState:       state.NukiState.String(),
		BatteryCritical: state.BatteryStateCritical,
		BatteryPercent:  state.BatteryPercentage,
	}, nil
}

// Lock locks the device.
func (c *Client) Lock() error {
	return c.performOperation(blecommands.Lock)
}

// Unlock unlocks the device.
func (c *Client) Unlock() error {
	return c.performOperation(blecommands.Unlock)
}

func (c *Client) performOperation(op blecommands.Action) error {
	ble, err := nukible.NewNukiBle()
	if err != nil {
		return fmt.Errorf("Bluetooth error: %w", err)
	}

	err = ble.ScanForDevice(c.deviceAddr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("scan failed: %w", err)
	}

	flow, err := bleflows.NewAuthenticatedFlow(ble, c.deviceAddr)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer flow.DisconnectDevice()

	err = flow.PerformLockOperation(op)
	if err != nil {
		return fmt.Errorf("operation failed: %w", err)
	}

	return nil
}
