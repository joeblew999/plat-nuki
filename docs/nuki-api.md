---
title: "Nuki API Options"
date: 2024-01-15
draft: false
---

# Nuki API Options

Nuki is one of the most open smart locks on the market. You can talk directly to the lock without going through their cloud.

## Pricing (US, December 2025)

| Product | Price | Notes |
|---------|-------|-------|
| **Nuki Smart Lock** (lock only) | **$199** | Retrofit, Matter/Thread, WiFi built-in |
| **Nuki Smart Lock + Keypad 2** (bundle) | **$269** | Lock + fingerprint/PIN keypad |
| **Keypad 2** (fingerprint + codes) | ~$149 | Add-on |
| **Keypad** (codes only) | ~$89 | Add-on |
| **Door Sensor** | ~$59 | Detects open/closed |
| **Nuki Bridge** | ~$99 | For local HTTP API |

### No More Subscription

Nuki dropped the subscription entirely — all features included in the hardware price:
- WiFi remote access
- Real-time notifications
- Full app features
- Matter over Thread

### What You Get

- **Retrofit install**: Clamps over existing deadbolt, ~15 min install
- **Open APIs**: Local BLE, Bridge HTTP, Matter — no cloud required
- **Matter/Thread**: Works with Apple Home, Google, Alexa, Home Assistant
- **No account required**: Can use without ever creating an account
- **End-to-end encryption**
- **Brushless motor**: Fast (<1 sec unlock), quieter than previous gen

### Europe vs US Pricing

| Region | Lock Only | Lock + Keypad 2 |
|--------|-----------|-----------------|
| **Europe** | €269 (~$280) | €349 (~$365) |
| **US** | $199 | $269 |

US is actually cheaper now.

## Communication Options

### 1. Local Bluetooth API — Direct to Lock

Communicate directly with the Nuki lock over **Bluetooth Low Energy (BLE)** using their published protocol.

- **Fully documented**: https://developer.nuki.io/page/nuki-smart-lock-api-2/2/
- No cloud, no internet required
- You authenticate with the lock directly
- Full control: lock, unlock, get status, manage users

**What you can do:**
- Pair your own app/device with the lock
- Send lock/unlock commands
- Read lock state
- Manage authorization entries (PINs, users)
- Get activity logs

**Protocol**: BLE with encrypted challenge-response (NaCl/libsodium crypto)

### 2. Nuki Bridge HTTP API — Local Network

If you add a **Nuki Bridge** (~$100), you get a local HTTP API on your network:

```
Your App → HTTP → Nuki Bridge → BLE → Lock
```

**API endpoint** (on your LAN):
```
http://<bridge-ip>:8080/lockAction?nukiId=<id>&action=1&token=<token>
```

**Actions:**

| Action | Code |
|--------|------|
| Unlock | 1 |
| Lock | 2 |
| Unlatch | 3 |
| Lock 'n' Go | 4 |
| Lock 'n' Go with Unlatch | 5 |

**Get lock state:**
```
GET http://<bridge-ip>:8080/lockState?nukiId=<id>&token=<token>
```

**Response:**
```json
{
  "state": 1,
  "stateName": "locked",
  "batteryCritical": false,
  "success": true
}
```

Full Bridge API docs: https://developer.nuki.io/page/nuki-bridge-http-api-1-13/4/

### 3. Nuki Web API — Cloud (if you want it)

They also have a REST API through their cloud:
```
https://api.nuki.io/
```

- OAuth2 authentication
- Webhooks for events
- Manage locks remotely
- Create/revoke access

But you **don't need this** — the local options work fine offline.

### 4. Matter (Nuki Smart Lock 4.0 Pro)

The newest Nuki locks support **Matter over Thread**:

- Works with Apple Home, Google Home, Home Assistant
- Local control via Matter protocol
- No Nuki cloud required

## Example: Talking to Nuki from Go

Using the Bridge HTTP API:

```go
package main

import (
    "fmt"
    "io"
    "net/http"
)

const (
    bridgeIP   = "192.168.1.50"
    bridgePort = "8080"
    nukiID     = "12345678"
    token      = "your-bridge-token"
)

func lockDoor() error {
    url := fmt.Sprintf("http://%s:%s/lockAction?nukiId=%s&action=2&token=%s",
        bridgeIP, bridgePort, nukiID, token)

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    return nil
}

func getState() error {
    url := fmt.Sprintf("http://%s:%s/lockState?nukiId=%s&token=%s",
        bridgeIP, bridgePort, nukiID, token)

    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    fmt.Println(string(body))
    return nil
}
```

## What You Need

| Setup | Components | Cost | API Type |
|-------|-----------|------|----------|
| **Direct BLE** | Nuki Lock + your BLE code | ~$150-250 | Bluetooth |
| **Local HTTP** | Nuki Lock + Nuki Bridge | ~$250-350 | REST (LAN) |
| **Matter** | Nuki 4.0 Pro + Thread border router | ~$250 | Matter |

## TL;DR

**Nuki is fully open:**

- Direct BLE protocol documented
- Local HTTP API via Bridge (no cloud)
- Matter support on new models
- You own the communication

This is exactly why Nuki is popular with hackers and home automation folks — you're not locked into their ecosystem.

## Go Libraries

### Official Nuki Go Libraries

| Library | URL | Description |
|---------|-----|-------------|
| **nuki-io/go-nuki** | https://github.com/nuki-io/go-nuki | Official Web API Client (cloud) |
| **nuki-io/nuki-cli** | https://github.com/nuki-io/nuki-cli | CLI tool (Bluetooth + Web API) |

### Community Go Libraries

| Library | URL | Description |
|---------|-----|-------------|
| **qvest-digital/go-nuki** | https://github.com/qvest-digital/go-nuki | Direct BLE — local control, no cloud |
| **christianschmizz/go-nukibridgeapi** | https://github.com/christianschmizz/go-nukibridgeapi | Bridge HTTP API wrapper |

### Library Comparison

| Library | Connection | Cloud Required | Use Case |
|---------|-----------|----------------|----------|
| **nuki-io/go-nuki** | Web API | Yes | Remote control via Nuki cloud |
| **nuki-io/nuki-cli** | BLE + Web | Optional | CLI tool for either |
| **qvest-digital/go-nuki** | BLE direct | No | Local control, no cloud |
| **go-nukibridgeapi** | HTTP (LAN) | No | Local via Bridge on network |

### Direct BLE Example (qvest-digital/go-nuki)

```go
package main

import (
    "context"
    "crypto/rand"

    "github.com/tarent/go-nuki"
    "github.com/tarent/go-nuki/communication/command"
    "github.com/go-ble/ble"
    "github.com/go-ble/ble/linux"
    "github.com/kevinburke/nacl/box"
)

func main() {
    // Create BLE device
    device, err := linux.NewDevice()
    if err != nil {
        panic(err)
    }

    nukiClient := nuki.NewClient(device)
    defer nukiClient.Close()

    // Connect to lock by MAC address
    // Find MAC with: hcitool lescan
    nukiDeviceAddr := ble.NewAddr("54:D2:AA:BB:CC:DD")
    err = nukiClient.EstablishConnection(context.Background(), nukiDeviceAddr)
    if err != nil {
        panic(err)
    }

    // Generate keypair for pairing
    publicKey, privateKey, err := box.GenerateKey(rand.Reader)
    if err != nil {
        panic(err)
    }

    // Pair with lock (lock must be in pairing mode)
    err = nukiClient.Pair(
        context.Background(),
        privateKey,
        publicKey,
        13,  // app ID
        command.ClientIdTypeApp,
        "MyGoApp",
    )
    if err != nil {
        panic(err)
    }

    // Now you can send commands...
}
```

### Integration Options

**Option A: Direct BLE** (qvest-digital/go-nuki)
- No extra hardware needed
- Your server needs Bluetooth
- Fully local, no cloud

**Option B: Bridge HTTP** (go-nukibridgeapi or roll your own)
- Need Nuki Bridge (~$99)
- Simple REST calls over LAN
- Server doesn't need Bluetooth

### Matter Option

The newest Nuki locks support **Matter over Thread**. If you go that route, you'd integrate via Home Assistant or a Matter controller rather than these Go libraries directly. Home Assistant has a good REST API you could call from Go.

## Resources

- BLE API: https://developer.nuki.io/page/nuki-smart-lock-api-2/2/
- Bridge HTTP API: https://developer.nuki.io/page/nuki-bridge-http-api-1-13/4/
- Web API: https://api.nuki.io/
