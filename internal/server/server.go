// Package server provides the HTTP API for remote lock control.
package server

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"

	"github.com/joeblew999/plat-nuki/internal/nuki"
)

// Server provides HTTP endpoints for Nuki lock control.
type Server struct {
	client *nuki.Client
	demo   bool
	mu     sync.Mutex
}

// New creates a new HTTP server for the given Nuki client.
// If client is nil, the server runs in demo mode.
func New(client *nuki.Client) *Server {
	return &Server{client: client, demo: client == nil}
}

// Response types for Huma

// StatusOutput is the response for GET /status.
type StatusOutput struct {
	Body struct {
		LockState       string `json:"lock_state" example:"locked" doc:"Current lock state"`
		NukiState       string `json:"nuki_state" example:"ready" doc:"Nuki device state"`
		BatteryCritical bool   `json:"battery_critical" example:"false" doc:"Battery critical warning"`
		BatteryPercent  int    `json:"battery_percent" example:"85" doc:"Battery percentage"`
	}
}

// LockOutput is the response for POST /lock.
type LockOutput struct {
	Body struct {
		Message string `json:"message" example:"Locked" doc:"Operation result"`
	}
}

// UnlockOutput is the response for POST /unlock.
type UnlockOutput struct {
	Body struct {
		Message string `json:"message" example:"Unlocked" doc:"Operation result"`
	}
}

// ListenAndServe starts the HTTP server on the given port.
func (s *Server) ListenAndServe(port int) error {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Nuki BLE Bridge API", "1.0.0"))

	// Register routes
	huma.Get(api, "/status", s.getStatus)
	huma.Post(api, "/lock", s.postLock)
	huma.Post(api, "/unlock", s.postUnlock)

	fmt.Printf("Starting HTTP server on port %d...\n", port)
	if s.demo {
		fmt.Println("Mode: DEMO (no real device)")
	} else {
		fmt.Printf("Device address: %s\n", s.client.DeviceAddr())
	}
	fmt.Println()
	fmt.Println("Endpoints:")
	fmt.Println("  GET  /status  - Get lock status")
	fmt.Println("  POST /lock    - Lock the device")
	fmt.Println("  POST /unlock  - Unlock the device")
	fmt.Println("  GET  /docs    - OpenAPI documentation")
	fmt.Println()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}

func (s *Server) getStatus(ctx context.Context, input *struct{}) (*StatusOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	resp := &StatusOutput{}

	if s.demo {
		resp.Body.LockState = "locked"
		resp.Body.NukiState = "ready"
		resp.Body.BatteryCritical = false
		resp.Body.BatteryPercent = 85
		return resp, nil
	}

	status, err := s.client.GetStatus()
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	resp.Body.LockState = status.LockState
	resp.Body.NukiState = status.NukiState
	resp.Body.BatteryCritical = status.BatteryCritical
	resp.Body.BatteryPercent = status.BatteryPercent
	return resp, nil
}

func (s *Server) postLock(ctx context.Context, input *struct{}) (*LockOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	resp := &LockOutput{}

	if s.demo {
		resp.Body.Message = "Locked (demo)"
		return resp, nil
	}

	err := s.client.Lock()
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	resp.Body.Message = "Locked"
	return resp, nil
}

func (s *Server) postUnlock(ctx context.Context, input *struct{}) (*UnlockOutput, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	resp := &UnlockOutput{}

	if s.demo {
		resp.Body.Message = "Unlocked (demo)"
		return resp, nil
	}

	err := s.client.Unlock()
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	resp.Body.Message = "Unlocked"
	return resp, nil
}
