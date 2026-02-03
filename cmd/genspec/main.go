//go:build ignore

// Command genspec generates the OpenAPI spec from the Huma API definition.
// Run with: go run cmd/genspec/main.go > openapi.json
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

// Response types - must match internal/server types

type StatusOutput struct {
	Body struct {
		LockState       string `json:"lock_state" example:"locked" doc:"Current lock state"`
		NukiState       string `json:"nuki_state" example:"ready" doc:"Nuki device state"`
		BatteryCritical bool   `json:"battery_critical" example:"false" doc:"Battery critical warning"`
		BatteryPercent  int    `json:"battery_percent" example:"85" doc:"Battery percentage"`
	}
}

type LockOutput struct {
	Body struct {
		Message string `json:"message" example:"Locked" doc:"Operation result"`
	}
}

type UnlockOutput struct {
	Body struct {
		Message string `json:"message" example:"Unlocked" doc:"Operation result"`
	}
}

func main() {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Nuki BLE Bridge API", "1.0.0"))

	// Register routes with dummy handlers (just for spec generation)
	huma.Get(api, "/status", func(ctx context.Context, input *struct{}) (*StatusOutput, error) { return nil, nil })
	huma.Post(api, "/lock", func(ctx context.Context, input *struct{}) (*LockOutput, error) { return nil, nil })
	huma.Post(api, "/unlock", func(ctx context.Context, input *struct{}) (*UnlockOutput, error) { return nil, nil })

	// Output OpenAPI 3.0.3 spec (downgraded from 3.1 for codegen compatibility)
	spec, err := api.OpenAPI().Downgrade()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downgrading spec: %v\n", err)
		os.Exit(1)
	}

	var prettyJSON map[string]any
	json.Unmarshal(spec, &prettyJSON)
	output, _ := json.MarshalIndent(prettyJSON, "", "  ")
	fmt.Println(string(output))
}
