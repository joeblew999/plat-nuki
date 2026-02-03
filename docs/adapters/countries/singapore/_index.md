---
title: "Singapore - Nuki Compatibility Guide"
date: 2024-01-15
draft: false
---

# Singapore - Nuki Compatibility Guide

**Status**: ✅ COMPATIBLE
**Nuki Model**: Smart Lock Pro (EU)
**Confidence**: High

---

## Summary

Singapore is one of the **best markets** for Nuki deployment in Asia. Most residential and commercial doors use **euro profile cylinders** in mortise locks - exactly what Nuki Pro is designed for.

---

## Lock Types in Singapore

### HDB Flats (Public Housing - 80% of population)

| Era | Lock Type | Nuki Compatible? |
|-----|-----------|------------------|
| **2010+** | Euro cylinder mortise | ✅ Yes |
| **2000-2010** | Euro cylinder mortise | ✅ Yes |
| **Pre-2000** | Older mortise, some rim locks | ⚠️ Check cylinder type |

**Most HDB doors** have euro profile cylinders. The standard Nuki Pro should work.

### Condominiums (Private)

| Type | Lock Type | Nuki Compatible? |
|------|-----------|------------------|
| **New launches** | Euro cylinder, often digital-ready | ✅ Yes |
| **Resale** | Euro cylinder mortise | ✅ Yes |
| **Luxury** | High-security euro (Mul-T-Lock, Abloy) | ✅ Yes (check adapter) |

### Commercial / Office

| Type | Lock Type | Notes |
|------|-----------|-------|
| **Office doors** | Euro cylinder | ✅ Compatible |
| **Server rooms** | Access control systems | May already be managed |
| **Shophouses** | Mixed, often euro | ✅ Check first |

---

## Common Cylinder Brands & Adapters

| Brand | Adapter | Adapter Guide | Where Found |
|-------|---------|---------------|-------------|
| **Dormakaba** | Yellow | [yellow-dormakaba](../types/yellow-dormakaba/) | Commercial, condos |
| **Yale** | Magenta | [magenta-yale](../types/magenta-yale/) | HDB, residential |
| **Häfele** | Check | - | Condos |
| **ISEO** | Black | [black-abus](../types/black-abus/) | Building projects |
| **Mul-T-Lock** | Test fit | - | High-security |

### What to Stock

For Singapore deployments, stock these adapters:
- [**Yellow (Dormakaba)**](../types/yellow-dormakaba/) - 40%
- [**Magenta (Yale)**](../types/magenta-yale/) - 40%
- [**Black (ISEO/Abus)**](../types/black-abus/) - 10%
- **Mixed spares** - 10%

---

## Installation Considerations

### HDB-Specific

- **Fire door regulations**: HDB main doors are fire-rated. Nuki mounts on the inside - doesn't affect fire rating.
- **Corridor width**: Check Nuki doesn't protrude too much into common corridor.
- **MCST approval**: Not typically required for internal door modifications, but check with your MCST.

### Condo-Specific

- **MCST rules**: Some condos restrict door modifications. Nuki is non-invasive (no drilling), so usually OK.
- **Intercom integration**: Nuki can integrate with some building systems via API.

---

## Where to Buy

| Source | Notes |
|--------|-------|
| [Nuki Shop](https://shop.nuki.io/en/) | Direct from Nuki, ships internationally |
| Amazon Singapore | Check for genuine stock |
| Lazada/Shopee | Verify seller authenticity |
| Local smart home retailers | Ask about warranty |

**Pricing**: Expect ~S$400-500 for Smart Lock Pro + Bridge

---

## Adapters Needed

![Knob Adapters](../../diagrams/knob-adapters-all.webp)

Most Singapore installations will use:
- [**Yellow adapter (Dormakaba)**](../types/yellow-dormakaba/) - commercial/condo
- [**Magenta adapter (Yale)**](../types/magenta-yale/) - HDB/residential

---

## Pre-Installation Checklist

1. **Identify your cylinder brand**
   - Look for stamped text on cylinder body
   - Check your keys for brand name

2. **Measure cylinder dimensions**
   - Inner length (inside to center screw)
   - Outer length (outside to center screw)
   - Max inner: 68mm, Max outer: 55mm

3. **Check door clearances**
   - Distance from cylinder to frame edge
   - Distance from cylinder to handle

4. **Confirm adapter**
   - Match brand to adapter color
   - Order correct adapter if not included

---

## Deployment Tips

### For Property Managers

- **Bulk ordering**: Contact Nuki for volume pricing
- **Master access**: Use Nuki Web API for centralized management
- **Activity logs**: Useful for security and compliance

### For Airbnb/Short-term Rental

- **Remote access codes**: Generate time-limited codes via app
- **Auto-lock**: Set timer to auto-lock after guests enter
- **Activity log**: Track check-in/check-out times

---

## Support & Warranty

- **Nuki warranty**: 2 years standard
- **Local support**: Via Nuki help center (English)
- **Voltage**: Singapore uses 230V - Nuki Bridge compatible

---

## See Also

- [Adapter Types](../types/) - All adapter documentation
- [Adapters Guide](../../) - Main adapters overview
- [Visual Diagrams](../../diagrams/)
- [Malaysia Guide](../malaysia/) - Similar market
- [Asia/Australia Overview](../../ASIA-AUSTRALIA-GAPS/)
