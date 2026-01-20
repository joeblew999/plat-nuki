# How Smart Locks Work When Battery Dies

## 1. Apple Home Key / NFC Wallet Keys — Power Reserve

Apple specifically designed this for exactly this problem:

- **Power Reserve mode** keeps the NFC chip functional for **up to 5 hours** after your iPhone dies
- Uses the phone's emergency power reserve (same tech that lets you use transit cards when dead)
- Works on **iPhone XS and later**
- Also works on Apple Watch

With the **Aqara U400** or any Apple Home Key lock, you can still tap-to-unlock even with a dead phone — the NFC passive chip doesn't need battery power to respond to the reader.

## 2. Lock-Side Emergency Power

Most modern locks have a **USB-C emergency port** on the exterior:

| Lock | Emergency Power Method |
|------|----------------------|
| **Aqara U400** | USB-C port on outside — plug in any power bank for 10 seconds, unlock, then charge properly inside |
| **Goki** | Same — external USB-C |
| **Most smart locks** | 9V battery terminal on exterior (touch a 9V battery to the contacts) |

## 3. Backup Access Methods

The Aqara U400 and similar locks have multiple fallbacks:

1. **Physical key** — yes, there's still a keyhole hidden under a cover
2. **PIN code on keypad** — doesn't require your phone at all
3. **DoorCode** (Goki) — staff can remotely generate a one-time code
4. **NFC card/tag** — separate from your phone

## 4. For Hotels Specifically

With **Goki** and similar hospitality systems:

- Guest gets a **DoorCode** (PIN) alongside their mobile key automatically
- If phone dies, they just punch in the code
- Staff can also **remotely unlock** from the dashboard
- Physical keycard backup still works

## The Aqara U400 Specifically

From the specs:

> *"In case you forget to recharge, the U400 lock features a USB-C Emergency Charging Port, allowing you to power up with a power bank or a phone that supports USB-C to USB-C power supply."*

You can literally use someone else's phone as a temporary power source to unlock, then charge the lock properly from inside.

## TL;DR

Between Apple's Power Reserve (5 hours after phone death), external USB-C emergency ports, PIN codes, and physical keys — you're covered. The days of being locked out because of dead batteries are pretty much solved.
