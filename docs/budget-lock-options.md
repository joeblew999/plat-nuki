# Budget Lock Options with API Access

Goki isn't cheap — **$300-500+ per lock** plus monthly platform fees. Here are cheaper alternatives with similar API capabilities.

## 1. TTLock-Based Locks — The Budget Hack

**Price: $50–$150 per lock**

TTLock is a Chinese platform that powers hundreds of cheap smart locks on Amazon/AliExpress. The key thing is they have an **open API**.

| Lock Examples | Price | Features |
|--------------|-------|----------|
| **Smonet** | ~$80 | Fingerprint, keypad, app, TTLock |
| **Catchface** | ~$60 | Keypad, Bluetooth, TTLock |
| **Hornbill** | ~$70 | Fingerprint, keypad, TTLock |
| **Generic "TTLock" branded** | ~$50 | Basic keypad + app |

### How provisioning works

```
Your System → TTLock API → Lock
```

TTLock API lets you:
- Create/delete passcodes programmatically
- Set time-limited access (check-in/check-out dates)
- Generate one-time codes
- Get unlock logs
- Manage multiple locks from one account

### Downsides
- Chinese servers (data sovereignty concerns)
- Hardware quality varies
- Some locks need a WiFi gateway (~$30) for remote access
- API documentation is rough

## 2. Tuya-Based Locks — Similar Story

**Price: $40–$120**

Tuya is another Chinese IoT platform with an API. Many budget locks use it.

- Search Amazon for "Tuya smart lock" or "Smart Life lock"
- Tuya has a developer portal with REST API
- Can integrate with Home Assistant via Tuya integration

## 3. Nuki — European, API-First

**Price: ~$150–250**

Nuki is a retrofit lock (clamps onto existing deadbolt) popular in Europe:

- **Proper REST API** with good documentation
- Webhooks for events
- Web API for remote access
- Works with Home Assistant, Airbnb integration available
- Matter support on newer models

Better quality than TTLock stuff, proper company behind it.

## 4. Build Your Own with Home Assistant

**Cost: Lock ($80-150) + Home Assistant (~$50 for a Pi or use existing server)**

### Stack

```
Your Booking System
      ↓ webhook
Home Assistant
      ↓ integration
Matter/Zigbee/WiFi Lock
```

- Use a **SwitchBot Lock Pro** (~$130) or **Aqara U100** (~$150)
- Home Assistant gives you full automation API
- Build your own provisioning logic
- Can integrate with any PMS via webhooks/REST

## 5. RemoteLock — Budget Commercial Option

**Locks: $150-300 | Software: ~$5-10/lock/month**

RemoteLock is a platform that works with multiple lock brands:
- Kwikset
- Yale
- Schlage
- August

They provide the **cloud platform + API** so you get Goki-like functionality but can use cheaper commodity locks.

## Quick Comparison

| Option | Lock Cost | Monthly Fee | API | Quality |
|--------|-----------|-------------|-----|---------|
| **Goki** | $300-500 | Yes | Yes | High |
| **TTLock locks** | $50-150 | Free | Yes (rough) | Variable |
| **Tuya locks** | $40-120 | Free | Yes | Variable |
| **Nuki** | $150-250 | Optional | Yes (good) | Good |
| **Home Assistant + Matter lock** | $80-150 + HA | Free | Yes (DIY) | Good |
| **RemoteLock** | $150-300 | ~$5-10/lock | Yes | Good |

## Recommendations by Stage

| Stage | Solution | Why |
|-------|----------|-----|
| **Prototype/MVP** | TTLock locks ($60-80) + TTLock API | Cheap, functional, prove the concept |
| **Production/Quality** | Nuki or SwitchBot + Home Assistant | Better hardware, cleaner integration |
| **Scale** | RemoteLock or similar platform | Proper commercial support |
