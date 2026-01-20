# Australia - Nuki Compatibility Guide

**Status**: ⚠️ PARTIAL COMPATIBILITY
**Nuki Model**: Smart Lock Pro (EU) for euro cylinders, Smart Lock (US) for deadbolts
**Confidence**: Medium - significant gaps exist

---

## Summary

Australia is a **challenging market** for Nuki. The most common residential lock - the **Lockwood 001** - is fundamentally incompatible. However, modern apartments with euro cylinders work fine.

**Bottom line**: Target the 25-35% of doors that ARE compatible (apartments, new builds), and have a lock replacement strategy for the rest.

---

## The Lockwood 001 Problem

![Lockwood 001 Rim Lock](images/lockwood-001-rimlock.jpg)

This is Australia's most common residential lock. It is **NOT compatible** with Nuki because:
- It's a **rim lock** (surface-mounted box), not a cylinder
- Has a **snib** (small twist), not a thumbturn Nuki can grip
- No adapter can fix this - the form factor is completely different

---

## Lock Types in Australia

### Breakdown by Property Type

| Property Type | Common Lock | Nuki Compatible? | % of Market |
|---------------|-------------|------------------|-------------|
| **Houses (older)** | Lockwood 001 | ❌ No | ~40% |
| **Houses (newer)** | Lockwood 001 or Euro | ⚠️ Check | ~20% |
| **Apartments (older)** | Mixed | ⚠️ Check | ~10% |
| **Apartments (newer)** | Euro cylinder | ✅ Yes | ~15% |
| **Commercial** | Euro cylinder | ✅ Yes | ~10% |
| **Other** | Various | ⚠️ Check | ~5% |

### Lock Type Details

| Lock Type | What It Looks Like | Nuki Model | Notes |
|-----------|-------------------|------------|-------|
| **Lockwood 001** | Surface-mounted box with snib | ❌ None | Lock replacement required |
| **Euro Cylinder** | Drop-shaped, in mortise lock | ✅ Smart Lock Pro | Modern apartments |
| **Lockwood Oval** | Oval-shaped cylinder | ⚠️ Custom adapter needed | Older/commercial |
| **US Deadbolt** | Rectangular with thumbturn | ✅ Smart Lock (US) | Some imports |

---

## Compatible Scenarios

### Apartments & Units (Best Target)

| City | Building Type | Likely Lock | Compatible? |
|------|---------------|-------------|-------------|
| **Sydney** | High-rise apartment | Euro cylinder | ✅ Yes |
| **Melbourne** | High-rise apartment | Euro cylinder | ✅ Yes |
| **Brisbane** | Modern unit | Euro cylinder | ✅ Yes |
| **Perth** | Apartment | Euro cylinder | ✅ Yes |

**Why apartments work**: Developers increasingly use euro profile mortise locks for:
- Insurance compliance
- Multi-point locking
- Easier key management

### Commercial & Office

Most commercial buildings use euro cylinders. Good fit for Nuki Pro.

### New Residential Construction

Post-2015 houses increasingly use euro cylinder mortise locks, especially in:
- Gated communities
- New housing estates
- Architect-designed homes

---

## Incompatible Scenarios

### Houses with Lockwood 001

**Solutions**:

| Option | Cost | Invasiveness | Result |
|--------|------|--------------|--------|
| **Replace lock** | $200-400 AUD + install | Medium | Full Nuki compatibility |
| **Alternative product** | $300-500 AUD | Low | August, Level, or Tedee |
| **Skip** | $0 | None | No smart lock |

### Lock Replacement Process

To replace a Lockwood 001 with a euro cylinder mortise lock:

1. **Hire a locksmith** - Don't DIY, fire door regulations apply
2. **Choose replacement** - Euro profile mortise lock
3. **Consider brands**: Lockwood Paradigm, Dormakaba, Abloy
4. **Budget**: $200-400 for lock + $100-200 installation

### Alternative Smart Locks

If you can't replace the lock:

| Product | Works With | Notes |
|---------|------------|-------|
| **August WiFi Smart Lock** | US deadbolts | If you have a deadbolt section |
| **Level Lock** | Deadbolts | Invisible from outside |
| **Tedee** | Euro cylinders | European brand |
| **Yale Assure Lock 2** | US deadbolts | Local support |

---

## Adapter Strategy

### For Euro Cylinder Doors

| Brand | Nuki Adapter | Where Found |
|-------|--------------|-------------|
| **Lockwood** (euro style) | Test fit | Residential |
| **Dormakaba** | Yellow | Commercial |
| **Abloy** | Test fit | High-security |
| **Yale** | Magenta | Retail/residential |

### Custom Adapter Opportunity

**Lockwood Oval Cylinder** - Some Australian doors use oval-shaped cylinders. No Nuki adapter exists.

**Opportunity**: Design a 3D-printable adapter for Lockwood oval cylinders.

See [3D Models](../../3d/README.md) for design guidelines.

---

## Regional Considerations

### New South Wales (Sydney)

- High-rise apartments: Euro cylinders, Nuki compatible
- Houses: Mostly Lockwood 001, not compatible
- Commercial: Euro cylinders

### Victoria (Melbourne)

- Similar to NSW
- Older Victorian homes: Lockwood or rim locks
- New apartments: Euro cylinders

### Queensland (Brisbane, Gold Coast)

- More houses, fewer apartments
- Higher % of Lockwood 001
- New developments: Euro cylinders

### Western Australia (Perth)

- Mix of old and new
- Mining company housing: Often standardized, check lock type
- Apartments: Euro cylinders

---

## Where to Buy

| Source | Notes |
|--------|-------|
| [YourSmartLife](https://www.yoursmartlifestore.com.au/collections/nuki) | Australian retailer, ~$350 AUD |
| [Amazon AU](https://www.amazon.com.au) | Ships from Germany |
| [Siegware](https://siegware.com.au/nuki-2-0-accessories/) | Nuki accessories |
| Direct from [Nuki Shop](https://shop.nuki.io/en/) | International shipping |

**Pricing**: Expect A$350-500 for Smart Lock Pro + Bridge

---

## Installation Considerations

### Fire Door Regulations

- Main doors in apartments are often fire-rated
- Nuki mounts internally, doesn't affect fire rating
- Don't modify the door itself

### Strata/Body Corporate

- Check strata rules before installing
- Some buildings restrict door modifications
- Nuki is non-invasive (no drilling on door)

### Power

- Australia uses 230V - Nuki Bridge compatible
- Standard AU power outlet for Bridge

---

## Pre-Installation Checklist

1. **Identify your lock type**
   - Lockwood 001 (rim lock) → NOT compatible
   - Euro cylinder → Compatible
   - Oval cylinder → Custom adapter needed

2. **If euro cylinder, identify brand**
   - Look for stamped text
   - Check keys

3. **Measure clearances**
   - Check Nuki will fit without hitting frame

4. **Order correct adapter**
   - Test fit if brand unknown

---

## Support & Warranty

- **Nuki warranty**: 2 years
- **Local support**: Via Nuki help center (English)
- **Voltage**: 230V compatible
- **Note**: No official Australian distributor - import or buy from AU retailers

---

## See Also

- [Adapters Guide](../../README.md)
- [Visual Diagrams](../../diagrams/README.md) - See Lockwood 001 image
- [3D Models](../../3d/README.md) - Custom adapter designs
- [New Zealand Guide](../new-zealand/) - Similar market
- [Asia/Australia Overview](../../ASIA-AUSTRALIA-GAPS.md)
