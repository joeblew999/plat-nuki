// Command bridge runs the Nuki BLE bridge on an embedded device.
// It provides CLI commands for BLE operations and an HTTP server for remote control.
package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/joeblew999/plat-nuki/internal/nuki"
	"github.com/joeblew999/plat-nuki/internal/server"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})))

	scanCmd := flag.NewFlagSet("scan", flag.ExitOnError)
	scanDuration := scanCmd.Duration("t", 10*time.Second, "scan duration")

	pairCmd := flag.NewFlagSet("pair", flag.ExitOnError)
	pairAddr := pairCmd.String("addr", "", "device BLE address (required)")
	pairPin := pairCmd.String("pin", "", "security PIN (if set on device)")

	lockCmd := flag.NewFlagSet("lock", flag.ExitOnError)
	lockAddr := lockCmd.String("addr", "", "device BLE address")

	unlockCmd := flag.NewFlagSet("unlock", flag.ExitOnError)
	unlockAddr := unlockCmd.String("addr", "", "device BLE address")

	statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
	statusAddr := statusCmd.String("addr", "", "device BLE address")

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	servePort := serveCmd.Int("port", 8080, "HTTP server port")
	serveAddr := serveCmd.String("addr", "", "device BLE address")
	serveDemo := serveCmd.Bool("demo", false, "run in demo mode (no real device)")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "scan":
		scanCmd.Parse(os.Args[2:])
		doScan(*scanDuration)
	case "pair":
		pairCmd.Parse(os.Args[2:])
		if *pairAddr == "" {
			fmt.Println("Error: -addr is required")
			pairCmd.Usage()
			os.Exit(1)
		}
		doPair(*pairAddr, *pairPin)
	case "lock":
		lockCmd.Parse(os.Args[2:])
		doLock(*lockAddr)
	case "unlock":
		unlockCmd.Parse(os.Args[2:])
		doUnlock(*unlockAddr)
	case "status":
		statusCmd.Parse(os.Args[2:])
		doStatus(*statusAddr)
	case "serve":
		serveCmd.Parse(os.Args[2:])
		doServe(*serveAddr, *servePort, *serveDemo)
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`Usage: nuki-bridge <command> [options]

Commands:
  scan     Scan for nearby Nuki devices
  pair     Pair with a Nuki device (put lock in pairing mode first)
  lock     Lock the paired device
  unlock   Unlock the paired device
  status   Get current lock status
  serve    Start HTTP server for remote control

Examples:
  nuki-bridge scan -t 15s
  nuki-bridge pair -addr "XX:XX:XX:XX:XX:XX"
  nuki-bridge status
  nuki-bridge unlock
  nuki-bridge lock
  nuki-bridge serve -port 8080`)
}

func doScan(duration time.Duration) {
	fmt.Printf("Scanning for Nuki devices for %v...\n\n", duration)

	results, err := nuki.Scan(duration)
	if err != nil {
		fmt.Printf("Scan error: %v\n", err)
	}

	for _, r := range results {
		fmt.Printf("Found: %s\n", r.Address)
		fmt.Printf("  Name: %s\n", r.Name)
		fmt.Printf("  RSSI: %d dBm\n", r.RSSI)
		fmt.Println()
	}

	fmt.Printf("Scan complete. Found %d device(s).\n", len(results))
}

func doPair(addr, pin string) {
	fmt.Printf("Pairing with %s...\n", addr)
	fmt.Println("Make sure the lock is in pairing mode (hold button for 5 seconds).")
	fmt.Println()

	err := nuki.Pair(addr, pin)
	if err != nil {
		fmt.Printf("Pairing failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Credentials saved to nuki-credentials.yaml")
	fmt.Println("Pairing successful!")
}

func doStatus(addr string) {
	client, err := nuki.NewClient(addr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connecting to %s...\n", client.DeviceAddr())

	status, err := client.GetStatus()
	if err != nil {
		fmt.Printf("Failed to read status: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nLock State: %s\n", status.LockState)
	fmt.Printf("Nuki State: %s\n", status.NukiState)
	fmt.Printf("Battery Critical: %v\n", status.BatteryCritical)
	fmt.Printf("Battery: %d%%\n", status.BatteryPercent)
}

func doLock(addr string) {
	client, err := nuki.NewClient(addr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connecting to %s...\n", client.DeviceAddr())
	fmt.Println("Locking...")

	err = client.Lock()
	if err != nil {
		fmt.Printf("Lock failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Locked!")
}

func doUnlock(addr string) {
	client, err := nuki.NewClient(addr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Connecting to %s...\n", client.DeviceAddr())
	fmt.Println("Unlocking...")

	err = client.Unlock()
	if err != nil {
		fmt.Printf("Unlock failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Unlocked!")
}

func doServe(addr string, port int, demo bool) {
	var client *nuki.Client
	var err error

	if demo {
		fmt.Println("Running in DEMO mode - no real device connection")
		client = nil
	} else {
		client, err = nuki.NewClient(addr)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("Hint: Use -demo flag to run without a paired device")
			os.Exit(1)
		}
	}

	srv := server.New(client)
	if err := srv.ListenAndServe(port); err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}
