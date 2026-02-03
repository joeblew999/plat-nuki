//go:build ignore

// Example usage of the generated Nuki API client
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joeblew999/plat-nuki/pkg/client"
)

func main() {
	// Create client pointing to your Nuki bridge server
	c, err := client.NewClient("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get lock status
	status, err := c.GetStatus(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Type switch to handle the response
	switch v := status.(type) {
	case *client.StatusOutputBody:
		fmt.Printf("Lock State: %s\n", v.GetLockState())
		fmt.Printf("Nuki State: %s\n", v.GetNukiState())
		fmt.Printf("Battery: %d%%\n", v.GetBatteryPercent())
		fmt.Printf("Battery Critical: %v\n", v.GetBatteryCritical())
	case *client.ErrorModelStatusCode:
		fmt.Printf("Error: %s\n", v.Response.GetDetail().Or("unknown error"))
	}

	// Lock the device
	lockRes, err := c.PostLock(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if body, ok := lockRes.(*client.LockOutputBody); ok {
		fmt.Printf("Lock result: %s\n", body.GetMessage())
	}

	// Unlock the device
	unlockRes, err := c.PostUnlock(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if body, ok := unlockRes.(*client.UnlockOutputBody); ok {
		fmt.Printf("Unlock result: %s\n", body.GetMessage())
	}
}
