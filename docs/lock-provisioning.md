---
title: "How Smart Locks Get Provisioned"
date: 2024-01-15
draft: false
---

# How Smart Locks Get Provisioned

Provisioning is where the real differences show up between cheap/consumer locks and commercial/hospitality systems.

## 1. Consumer Locks (Budget Options)

These typically use **app-based provisioning**:

### Setup Flow
1. Download manufacturer's app (eufy Security, Wyze, TEEHO, etc.)
2. Create account
3. Put lock in pairing mode (usually hold a button)
4. App discovers lock via **Bluetooth**
5. Connect lock to your WiFi (if supported)
6. You're the admin — now you can add codes/fingerprints

### Adding Users
- **Fingerprints**: Enrolled directly on the lock (press fingerprint sensor multiple times)
- **PIN codes**: Created in app, synced to lock
- **Guest access**: Generate temporary/one-time codes in app, share via text/email

### Limitations
- Single admin account (or limited sharing)
- No central management for multiple locks
- No API — can't automate provisioning
- Each lock is its own island

## 2. Matter-Enabled Locks (Aqara, SwitchBot, etc.)

### Setup Flow
1. Scan QR code on lock (or enter setup code)
2. Add to your Matter controller (Apple Home, Google Home, SmartThings, Home Assistant)
3. Lock joins your Thread/WiFi network
4. Manage through your smart home platform

### Adding Users
- Through the Matter controller app
- Apple Home Key: Users get keys automatically when added to your "Home"
- PIN codes: Created in app, pushed to lock

### Better because
- Works across ecosystems
- Can be managed alongside other smart home devices
- Home Assistant = full API access for automation

## 3. Hospitality/Commercial Locks (Goki, SALTO KS, etc.)

This is where it gets **actually scalable**.

### Architecture

```
PMS (Property Management System)
         ↓ API
   Lock Platform (Cloud)
         ↓ BLE/WiFi/Thread
      Door Locks
```

### Provisioning Flow

1. **Lock Installation**
   - Install lock on door
   - Lock connects to gateway/hub (BLE, WiFi, or cellular)
   - Lock registers with cloud platform

2. **Guest Provisioning (Automated)**
   - Guest makes reservation in PMS
   - PMS sends booking data to lock platform via API
   - Lock platform generates credentials (mobile key, PIN, etc.)
   - Guest receives access automatically (email, SMS, app)

3. **Access Activation**
   - Credentials only work during stay dates
   - Automatically revoked at checkout
   - No manual intervention required

### Example — Goki

```
Guest books room → PMS (e.g., Mews, Cloudbeds)
                 → Goki API
                 → Generates DoorCode + Mobile Key
                 → Guest gets SMS/email with access
                 → Lock validates credentials locally
```

## Comparison Table

| Aspect | Consumer Lock | Matter Lock | Hospitality Lock |
|--------|--------------|-------------|------------------|
| **Setup** | App + Bluetooth | QR code scan | Cloud dashboard + gateway |
| **Admin** | Single account | Home members | Role-based (admin, staff, guest) |
| **Add users** | Manual in app | Through home app | API / PMS integration |
| **Guest codes** | Manual create/share | Manual or automation | Auto-generated per booking |
| **Expiring access** | Some support it | Yes | Yes, tied to reservation |
| **Multi-lock management** | One at a time | Per "home" | Centralized dashboard |
| **API access** | No | Via Home Assistant | Yes, REST/webhooks |
| **Scalability** | 1-5 locks | 1-20 locks | Unlimited |

## If You Want API/Automation with Cheap Locks

### Options

1. **Home Assistant + Matter lock**
   - SwitchBot Lock Pro or Aqara U100/U200 (not the U400)
   - Full API access via Home Assistant
   - Can build your own provisioning logic

2. **TTLock-based locks** (many cheap Chinese locks use this)
   - TTLock has a developer API
   - Can integrate with custom systems
   - Quality varies wildly

3. **Seam API**
   - Universal API for smart locks
   - Supports August, Yale, Schlage, some others
   - Not the cheapest locks though

## Bottom Line

| Use Case | Solution |
|----------|----------|
| Personal use | Consumer locks — manual provisioning is fine |
| Small rental/Airbnb | Matter locks + Home Assistant — DIY automation possible |
| Scale (hotels, multiple properties) | Hospitality platforms (Goki, SALTO KS) — API-first, auto-provisioning |
