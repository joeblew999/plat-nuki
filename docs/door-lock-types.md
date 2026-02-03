---
title: "Door Lock Types & Adapters"
date: 2024-01-15
draft: false
---

# Door Lock Types & Adapters

Understanding lock types is essential for smart lock compatibility. The US and Europe use fundamentally different systems.

## US vs European Locks

| Feature | US Deadbolt | European Euro Cylinder |
|---------|-------------|------------------------|
| **Shape** | Rectangular bolt | Drop-shaped cylinder |
| **Mechanism** | Single bolt into frame | Often multi-point (3+ bolts) |
| **Thumbturn** | Above lever, 90° turn | Below lever, 360° turn |
| **Modularity** | Replace entire lockset | Replace just the cylinder |
| **Installation** | DIY, 15 min | Precise measurement required |

### US Deadbolt

```
┌─────────────────┐
│     DOOR        │
│                 │
│  ┌───────────┐  │
│  │ DEADBOLT  │──┼──► Bolt extends into frame
│  │ (inside)  │  │
│  │           │  │
│  │ [THUMB]   │  │    Thumbturn on inside
│  │  TURN     │  │
│  └───────────┘  │
│                 │
│  ┌───────────┐  │
│  │  HANDLE   │  │
│  └───────────┘  │
└─────────────────┘
```

- **How it works**: Turn thumbturn → bolt slides into door frame
- **Smart lock mounts**: Over the thumbturn on the inside
- **Common brands**: Schlage, Kwikset, Baldwin, Defiant

### European Euro Profile Cylinder

```
┌─────────────────┐
│     DOOR        │
│                 │
│  ┌───────────┐  │
│  │  EURO     │  │    Drop-shaped cylinder
│  │ CYLINDER  │──┼──► Key or thumbturn inside
│  │   ◉       │  │
│  └───────────┘  │
│                 │
│  ══════════════ │    Multi-point locking bar
│       │         │    (engages at 3+ points)
│  ══════════════ │
│                 │
│  ┌───────────┐  │
│  │  HANDLE   │  │    Lift handle to engage
│  └───────────┘  │
└─────────────────┘
```

- **How it works**: Turn key/thumbturn 360° → multiple bolts engage
- **Smart lock replaces**: The cylinder, or mounts over existing
- **Common brands**: Abus, EVVA, DOM, Dormakaba, KESO

---

## The Tailpiece Problem

The **tailpiece** is the metal bar that connects the cylinder to the lock mechanism. It's why adapters exist.

### US Deadbolt Tailpiece

```
    ┌─────┐
    │     │  ← Thumbturn (you grip this)
    └──┬──┘
       │
    ┌──┴──┐
    │     │  ← Mounting plate
    └──┬──┘
       │
    ╔══╧══╗
    ║     ║  ← TAILPIECE (turns the bolt)
    ║     ║
    ╚═════╝
       │
    [BOLT]
```

**Tailpiece shapes vary by manufacturer:**

| Shape | Looks Like | Common Brands |
|-------|------------|---------------|
| Square | ▢ | Schlage, some Kwikset |
| D-shaped (DD) | ◗◖ | Many US brands |
| Round | ○ | Less common |
| Slotted | ═ | Older locks |

### European Cylinder Tailpiece

```
    ┌─────────────────────┐
    │   EURO CYLINDER     │
    │         ◉           │  ← Keyhole
    └─────────┬───────────┘
              │
           ┌──┴──┐
           │     │  ← CAM (flat metal piece)
           └─────┘
              │
           [LOCK MECHANISM]
```

**Euro cylinder tailpiece (cam) shapes:**

| Type | Inner Diameter | Common Brands |
|------|----------------|---------------|
| Standard | 7.5mm - 8mm | Most European |
| Large | 8mm - 9mm | Some Abus, EVVA |
| Knob | N/A (has knob) | Various |

---

## Nuki Adapters

### US Model - 3 Adapters Included

The Nuki Smart Lock (US) ships with **3 adapters** for different tailpiece shapes:

```
Adapter 1        Adapter 2        Adapter 3
┌───────┐       ┌───────┐       ┌───────┐
│  ▢    │       │  ◗◖   │       │  ═    │
│       │       │       │       │       │
└───────┘       └───────┘       └───────┘
 Square          D-shape         Slotted
```

**Fits most**: Schlage, Kwikset, Baldwin, Defiant, Weiser

### European Model - Knob Cylinder Adapters

Nuki sells **8 color-coded adapters** for European knob cylinders:

| Color | Fits |
|-------|------|
| Black | Abus, CISA, Bricard, AXA, ISEO |
| Green | CES, EVVA |
| Blue | DOM, Winkhaus |
| Yellow | Dormakaba (Kaba) |
| Red | KESO |
| White | Nemef |
| Magenta | Yale, M&C |
| Orange | Mauer |

**Why so many?** Each manufacturer uses slightly different knob dimensions.

---

## When You Need a Custom Adapter

### Signs the included adapters won't work:

1. **Tailpiece too long** - Extends past the adapter
2. **Tailpiece too short** - Doesn't engage the adapter
3. **Wrong shape** - Spins freely, doesn't grip
4. **Knob cylinder** - European style with a knob, not a key

### Solutions:

| Problem | Solution |
|---------|----------|
| Wrong shape | 3D print custom adapter |
| Too long | Cut tailpiece (carefully!) |
| Too short | Use extension plate (included) |
| Knob cylinder | Buy matching Nuki adapter or print one |

---

## Measuring Your Lock

### US Deadbolt

```
        ← 1.38" min →
    ┌───────┬───────┬───────┐
    │ FRAME │ DEAD  │ HANDLE│
    │       │ BOLT  │       │
    └───────┴───────┴───────┘
            ↑
        Center of deadbolt
```

**For inward-opening doors:**
- Deadbolt center to frame: > 1.38" (35mm)
- Deadbolt center to handle: > 1.38" (35mm)

**For outward-opening doors:**
- Deadbolt center to frame: > 1.58" (40mm)
- Deadbolt center to handle: > 1.38" (35mm)

### European Euro Cylinder

```
    ← Inner length →   ← Outer length →
    ┌────────────────┬────────────────┐
    │     INSIDE     │    OUTSIDE     │
    │                │                │
    └────────────────┴────────────────┘
                     ↑
              Retaining screw
```

**Measure from center screw:**
- Inner length: Should not exceed 68mm
- Outer length: Should not exceed 55mm

---

## Compatibility Quick Reference

### Nuki Smart Lock (US)

| Lock Brand | Compatible | Notes |
|------------|------------|-------|
| Schlage | ✅ Yes | Most models |
| Kwikset | ✅ Yes | Most models |
| Baldwin | ✅ Yes | Most models |
| Defiant | ✅ Yes | Home Depot brand |
| Weiser | ✅ Yes | Canadian brand |
| Yale (US) | ✅ Yes | US deadbolt models |
| **Euro cylinder** | ❌ No | Wrong lock type |

### Nuki Smart Lock Pro (EU)

| Lock Brand | Compatible | Adapter Needed |
|------------|------------|----------------|
| Abus | ✅ Yes | Black adapter |
| EVVA | ✅ Yes | Green adapter |
| DOM | ✅ Yes | Blue adapter |
| Dormakaba | ✅ Yes | Yellow adapter |
| KESO | ✅ Yes | Red adapter |
| **US Deadbolt** | ❌ No | Wrong lock type |

---

## Lock Standards by Country

The door industry is fragmented by region. What works in New York won't fit in Berlin, Stockholm, or Tokyo.

### World Map of Lock Types

| Region | Primary Lock Type | Standard | Nuki Model |
|--------|-------------------|----------|------------|
| **USA** | Deadbolt | ANSI/BHMA | Smart Lock (US) |
| **Canada** | Deadbolt | ANSI/BHMA | Smart Lock (US) |
| **Mexico** | Deadbolt | ANSI/BHMA | Smart Lock (US) |
| **UK** | Euro Cylinder + Night Latch | BS 3621 | Smart Lock Pro |
| **Ireland** | Euro Cylinder | BS 3621 | Smart Lock Pro |
| **Germany** | Euro Cylinder | DIN 18252 | Smart Lock Pro |
| **Austria** | Euro Cylinder | DIN/ÖNORM | Smart Lock Pro |
| **France** | Euro Cylinder | NF | Smart Lock Pro |
| **Italy** | Euro Cylinder | UNI | Smart Lock Pro |
| **Spain** | Euro Cylinder | UNE | Smart Lock Pro |
| **Netherlands** | Euro Cylinder | NEN | Smart Lock Pro |
| **Belgium** | Euro Cylinder | NBN | Smart Lock Pro |
| **Switzerland** | Euro Cylinder | SN | Smart Lock Pro |
| **Sweden** | Scandinavian Oval | SS | Smart Lock Ultra Nordics |
| **Norway** | Scandinavian Oval | NS | Smart Lock Ultra Nordics |
| **Denmark** | Scandinavian Oval | DS | Smart Lock Ultra Nordics |
| **Finland** | Scandinavian Oval/Round | SFS | Smart Lock Ultra Nordics |
| **Australia** | Lockwood/Oval | AS 4145 | ⚠️ Smart Lock Pro (some) |
| **New Zealand** | Lockwood/Oval | NZS | ⚠️ Smart Lock Pro (some) |
| **Singapore** | Euro Cylinder/Mortise | SS | ✅ Smart Lock Pro |
| **Malaysia** | Euro Cylinder/Mortise | MS | ✅ Smart Lock Pro |
| **Hong Kong** | Euro Cylinder/Mortise | - | ✅ Smart Lock Pro |
| **Taiwan** | Euro Cylinder/Mortise | CNS | ✅ Smart Lock Pro |
| **Thailand** | Euro Cylinder/Mortise | TIS | ✅ Smart Lock Pro |
| **Philippines** | Deadbolt/Euro | PNS | ⚠️ Smart Lock (US) or Pro |
| **Indonesia** | Euro Cylinder/Mixed | SNI | ⚠️ Smart Lock Pro |
| **Vietnam** | Euro Cylinder/Mixed | TCVN | ⚠️ Smart Lock Pro |
| **Japan** | Mortise (MIWA style) | JIS | ❌ Different system |
| **South Korea** | Digital/Mortise | KS | ❌ Different system |
| **China** | Various | GB | ⚠️ Varies widely |

### Regional Details

#### North America (USA, Canada, Mexico)
```
Standard: ANSI/BHMA (American National Standards Institute)
Lock type: Single-cylinder deadbolt with thumbturn
Common brands: Schlage, Kwikset, Baldwin, Defiant, Weiser
Nuki model: Smart Lock (US) - 3 adapters included
```

**Characteristics:**
- Deadbolt + separate handle/knob
- Single-point locking (one bolt)
- Thumbturn on inside, key on outside
- DIY installation culture
- Big box store availability (Home Depot, Lowe's)

#### Western Europe (Germany, Austria, France, Italy, Spain, Netherlands, Belgium, Switzerland)
```
Standard: DIN 18252 (Germany), various national standards
Lock type: Euro profile cylinder in multipoint lock
Common brands: Abus, EVVA, DOM, Dormakaba, KESO, Winkhaus
Nuki model: Smart Lock Pro + color-coded adapters
```

**Characteristics:**
- Drop-shaped euro cylinder
- Multi-point locking (3+ bolts engage)
- Lift handle to engage all bolts
- Cylinder is modular (replace without changing lock)
- Professional locksmith culture

#### United Kingdom & Ireland
```
Standard: BS 3621 (British Standard)
Lock type: Euro cylinder OR night latch (Yale lock)
Common brands: Yale, ERA, Union, Chubb
Nuki model: Smart Lock Pro
```

**UK has TWO common lock types:**

1. **Euro Cylinder** (modern doors)
   - Same as Western Europe
   - Often in UPVC doors
   - Nuki Pro compatible

2. **Night Latch / Rim Lock** (older doors)
   - Surface-mounted on inside
   - "Yale lock" colloquially
   - Nuki does NOT fit these directly

**Insurance note:** UK insurers often require BS 3621 locks.

#### Scandinavia (Sweden, Norway, Denmark, Finland)
```
Standard: SS/NS/DS/SFS (national standards)
Lock type: Scandinavian oval or round cylinder
Common brands: ASSA, Abloy, Ruko, Trioving
Nuki model: Smart Lock Ultra Nordics
```

**Characteristics:**
- Oval-shaped cylinder (not euro profile!)
- 360° key rotation
- Very high security standards
- ASSA/Abloy dominance
- Nuki made a specific "Nordics" model for this market

#### Australia & New Zealand
```
Standard: AS 4145.2 (Australia), NZS (New Zealand)
Lock type: Lockwood 001, oval cylinder, euro cylinder (newer)
Common brands: Lockwood, Gainsborough, Yale, Legge
Nuki model: Smart Lock Pro (for euro cylinder doors)
```

**Lock types found in AU/NZ:**

| Type | Where Found | Nuki Compatible? |
|------|-------------|------------------|
| **Lockwood 001** | Most common residential | ❌ No - rim lock style |
| **Euro Cylinder** | Modern doors, apartments | ✅ Yes - Smart Lock Pro |
| **Oval Cylinder** | Older/commercial | ⚠️ Possibly with adapter |
| **Deadbolt** | Some US-style imports | ✅ Yes - Smart Lock (US) |

**The Lockwood 001 problem:**
- Australia's #1 residential lock
- Rim-mounted deadlatch (not cylinder-based)
- Cannot use Nuki directly
- **Solution**: Replace with euro cylinder mortise lock, or use August/Level

**Where to buy Nuki in Australia:**
- [YourSmartLife](https://www.yoursmartlifestore.com.au/collections/nuki) - ~$350 AUD
- [Amazon AU](https://www.amazon.com.au) - Ships from Germany
- [Siegware](https://siegware.com.au/nuki-2-0-accessories/)

**Adapter opportunity:** 🖨️ Custom adapters for Lockwood oval cylinders

---

#### Southeast Asia (Singapore, Malaysia, Hong Kong, Taiwan, Thailand)
```
Standard: Various (SS, MS, CNS, TIS)
Lock type: Euro cylinder mortise locks (dominant)
Common brands: Dormakaba, Häfele, Yale, Gateman, Kaadas
Nuki model: Smart Lock Pro
```

**Good news:** Southeast Asia widely uses **euro profile cylinders** in mortise locks, especially in:
- HDB flats (Singapore)
- Condominiums
- Commercial buildings
- Modern residential

**Common setup:**
```
┌─────────────────┐
│     DOOR        │
│                 │
│  ┌───────────┐  │
│  │  EURO     │  │  ← Euro cylinder (Nuki Pro fits!)
│  │ CYLINDER  │  │
│  └───────────┘  │
│                 │
│  ┌───────────┐  │
│  │  MORTISE  │  │  ← Mortise lock body
│  │   LOCK    │  │
│  └───────────┘  │
│                 │
│  ┌───────────┐  │
│  │  HANDLE   │  │
│  └───────────┘  │
└─────────────────┘
```

**Cylinder brands in SEA:**
- Dormakaba (German)
- Häfele (German)
- Yale (International)
- ISEO (Italian)
- Mul-T-Lock (Israeli)

**Adapter situation:** Most should work with standard Nuki Pro adapters since they use euro profile. Check cylinder brand against Nuki's color-coded adapter chart.

**Where to buy:**
- Direct from [Nuki shop](https://shop.nuki.io/en/) (ships internationally)
- Amazon Singapore/Malaysia
- Local smart home retailers

---

#### Philippines, Indonesia, Vietnam
```
Standard: PNS, SNI, TCVN
Lock type: Mixed - both US deadbolt and euro cylinder
Common brands: Corona, Yale, Häfele, local brands
Nuki model: Depends on lock type
```

**Mixed market:**
- **US-style deadbolts**: Common in houses (US influence)
- **Euro cylinders**: Common in condos, commercial
- **Local variations**: May need custom adapters

**Decision tree:**
```
What lock do you have?
         │
    ┌────┴────┐
    ▼         ▼
Deadbolt?   Euro Cylinder?
    │              │
    ▼              ▼
Smart Lock    Smart Lock
   (US)          Pro
```

**Adapter opportunity:** 🖨️ Local cylinder brands may need custom adapters

---

#### Japan
```
Standard: JIS (Japanese Industrial Standards)
Lock type: Mortise lock (MIWA, GOAL, ALPHA)
Common brands: MIWA, GOAL, ALPHA, Keiden
Nuki model: ❌ Not compatible
```

**Characteristics:**
- Pin tumbler mortise locks
- MIWA LA-type very common
- Digital locks increasingly popular
- Different form factor than US/EU
- Local brands dominate (Samsung, Epic not common)

**Smart lock options:** MIWA, GOAL, Qrio, Sesame (local brands)

#### South Korea
```
Standard: KS (Korean Standards)
Lock type: Digital mortise lock
Common brands: Samsung, EPIC, Gateman, iRevo
Nuki model: ❌ Not compatible
```

**Characteristics:**
- Digital locks are the DEFAULT
- Keypad/fingerprint/card standard in apartments
- Mortise-style integration
- Most advanced digital lock market
- Push-pull handles common

**Note:** Korea skipped mechanical smart locks entirely.

---

## Why This Matters for Nuki

### Which Nuki Do You Need?

```
┌─────────────────────────────────────────────────────────┐
│                    WHERE ARE YOU?                        │
└─────────────────────────────────────────────────────────┘
                           │
         ┌─────────────────┼─────────────────┐
         ▼                 ▼                 ▼
    ┌─────────┐      ┌───────────┐    ┌────────────┐
    │   USA   │      │  EUROPE   │    │ SCANDINAVIA│
    │ Canada  │      │   (most)  │    │  (Nordic)  │
    │ Mexico  │      │           │    │            │
    └────┬────┘      └─────┬─────┘    └──────┬─────┘
         │                 │                  │
         ▼                 ▼                  ▼
  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐
  │ Smart Lock   │  │ Smart Lock   │  │ Smart Lock   │
  │    (US)      │  │    Pro       │  │Ultra Nordics │
  │              │  │              │  │              │
  │ 3 adapters   │  │ 8 adapters   │  │ Oval/round   │
  │ included     │  │ (buy extra)  │  │ cylinders    │
  └──────────────┘  └──────────────┘  └──────────────┘
```

### Countries Where Nuki Won't Work (Without Major Changes)

| Country | Why | Alternative |
|---------|-----|-------------|
| Japan | MIWA mortise standard | Qrio, Sesame, MIWA smart |
| South Korea | Digital mortise default | Samsung, EPIC, Gateman |
| Australia (Lockwood 001) | Rim lock, not cylinder | Replace lock or use August |

---

## Adapter Sourcing by Region

### Official Nuki Adapters

| Adapter | Cylinder Brands | Where to Buy | Price |
|---------|-----------------|--------------|-------|
| Black | Abus, CISA, Bricard, AXA, ISEO | [Nuki Shop](https://nuki.io/en-at/products/spare-parts/adapter-smart-locks) | ~€10 |
| Green | CES, EVVA | Nuki Shop | ~€10 |
| Blue | DOM, Winkhaus | Nuki Shop | ~€10 |
| Yellow | Dormakaba (Kaba) | Nuki Shop | ~€10 |
| Red | KESO | Nuki Shop | ~€10 |
| White | Nemef | Nuki Shop | ~€10 |
| Magenta | Yale, M&C | Nuki Shop | ~€10 |
| Orange | Mauer | Nuki Shop | ~€10 |

### 3D Printable Adapters (Fill the Gaps)

| Gap in Market | Solution | Print Files |
|---------------|----------|-------------|
| Abus door knob | Custom adapter | [MakerWorld](https://makerworld.com/en/models/1373160) |
| EVVA knob cylinder | Gen 5 mount kit | [Printables](https://www.printables.com/model/1264912) |
| Better adhesion | Spacer plate | [Printables](https://www.printables.com/model/379910) |
| Various inner diameters | Multi-size set | [Cults3D](https://cults3d.com/en/tags/nuki) |
| Lockwood (AU) oval | **NEEDED** | 🖨️ Design opportunity |
| MIWA (Japan) | **NEEDED** | 🖨️ Design opportunity |
| Local SEA brands | **NEEDED** | 🖨️ Design opportunity |

### Regional Adapter Strategy

#### North America
- **Status**: ✅ Covered
- **Included**: 3 adapters in box (square, D-shape, slotted)
- **Fits**: Schlage, Kwikset, Baldwin, Defiant, Weiser
- **Gaps**: Rare tailpiece shapes only

#### Western Europe
- **Status**: ✅ Covered
- **Official**: 8 color-coded adapters from Nuki
- **Fits**: All major euro cylinder brands
- **Gaps**: Obscure local brands

#### Scandinavia
- **Status**: ✅ Covered
- **Model**: Smart Lock Ultra Nordics
- **Fits**: ASSA, Abloy, Ruko, Trioving
- **Gaps**: None significant

#### Australia/New Zealand
- **Status**: ⚠️ Gaps exist
- **Works**: Euro cylinder doors (Smart Lock Pro)
- **Problem**: Lockwood 001 (most common) is rim lock
- **Needed**:
  - Lockwood oval cylinder adapter
  - Or: Replace Lockwood 001 with euro mortise lock
- **3D print opportunity**: Oval cylinder adapters

#### Southeast Asia
- **Status**: ⚠️ Mostly covered
- **Works**: Euro cylinder doors (very common)
- **Check**: Cylinder brand against Nuki adapter chart
- **Potential gaps**: Local/regional cylinder brands
- **Recommendation**: Test fit before bulk deployment

#### Japan
- **Status**: ❌ Not compatible
- **Problem**: MIWA LA-type mortise standard
- **Solution**: Use local smart locks (Qrio, Sesame)
- **Future**: Adapter would require significant engineering

#### South Korea
- **Status**: ❌ Not applicable
- **Why**: Digital locks are already standard
- **Market**: Already post-smart-lock adoption

### Adapter Design Priorities

If you're deploying Nuki in new markets, here are the gaps to fill:

| Priority | Market | Lock Type | Difficulty |
|----------|--------|-----------|------------|
| 🔴 High | Australia | Lockwood oval | Medium |
| 🟡 Medium | SEA | Local brands | Low-Medium |
| 🟡 Medium | UK | Night latch bypass | High |
| 🟢 Low | Europe | Obscure brands | Low |
| ⚫ Hard | Japan | MIWA mortise | Very High |

### How to Design a Custom Adapter

1. **Measure the cylinder**
   - Inner diameter (where key goes)
   - Outer diameter (body of cylinder)
   - Knob dimensions (if applicable)

2. **Get the Nuki adapter interface dimensions**
   - The part that connects to the Nuki
   - Match existing adapter as template

3. **Model in CAD**
   - OpenSCAD for parametric (adjustable)
   - Fusion 360 for precise modeling
   - TinkerCAD for simple shapes

4. **Print and test**
   - PLA for prototyping
   - PETG for durability
   - 0.2mm layer height
   - 40% infill minimum

5. **Share on Printables/Cults3D**
   - Help the community
   - Get feedback

See [3D Models](adapters/3d/) for existing designs.

---

## Resources

- [Nuki US Compatibility Guide](https://nuki.io/en-us/products/product-details/smart-lock-us-compatibility-guide)
- [Nuki EU Adapters](https://nuki.io/en-at/products/spare-parts/adapter-smart-locks)
- [Nuki Nordics](https://nuki.io/en-001/products/smart-lock-ultra-nordics)
- [Level Lock Tailpiece Guide](https://level.co/support/articles/tailpiece-adapters/)
- [IPVM Lock Profiles Guide](https://ipvm.com/reports/different-locks-different-countries)

## See Also

- [Adapters Guide](adapters/) - Official adapters and how they work
- [3D Printable Adapters](adapters/3d/) - Community solutions when stock adapters don't fit
- [Nuki API](./nuki-api/) - Technical integration options
