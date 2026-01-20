# 3D Models for Nuki Smart Lock

Community-created 3D printable accessories for Nuki Smart Locks.

> **Note**: These are community/third-party models, not official Nuki products.
> Nuki does not publish official CAD files for their locks.

## Model Sources (Full Provenance)

### Primary Repositories

| Repository | URL | Model Count | Formats | License |
|------------|-----|-------------|---------|---------|
| **Cults3D** | https://cults3d.com/en/tags/nuki | 33+ | STL, OBJ, STEP, 3MF, SCAD | Varies by model |
| **Printables** | https://www.printables.com/search/models?q=nuki | 50+ | STL, STEP, 3MF | Varies by model |
| **MakerWorld** | https://makerworld.com/en/search/models?keyword=nuki | 10+ | STL, STEP | Varies by model |
| **Thingiverse** | https://www.thingiverse.com/search?q=nuki | 20+ | STL | CC licenses |

### Search Aggregators

| Site | URL | Notes |
|------|-----|-------|
| **STLFinder** | https://www.stlfinder.com/3dmodels/nuki/ | Aggregates from multiple sites |
| **Yeggi** | https://www.yeggi.com/q/nuki+smart+lock/ | Another aggregator |
| **Thangs** | https://thangs.com/search/nuki | Includes AI-similar models |

---

## Recommended Models

### Adapters

| Model | Author | URL | Description | Last Verified |
|-------|--------|-----|-------------|---------------|
| **Abus Door Knob Adapter** | - | https://makerworld.com/en/models/1373160-nuki-smart-lock-pro-adapter-for-abus-door-knob | STEP file included for modification | Jan 2026 |
| **EVVA Knob Cylinder Mount** | adissu | https://www.printables.com/model/1264912-nuki-smart-lock-pro-gen-5-mounting-kit-for-knob-cy | For Gen 5 Pro | Jan 2026 |
| **Multi-size Adapter Set** | - | https://cults3d.com/en/tags/nuki | Various inner diameters 7.5-9mm | Jan 2026 |

### Spacers & Mounting

| Model | Author | URL | Description | Last Verified |
|-------|--------|-----|-------------|---------------|
| **Smart Lock Spacer** | vmLOGIC | https://www.printables.com/model/379910-nuki-smart-lock-spacer | Larger adhesive surface | Jan 2026 |
| **Universal Spacer Plate** | - | Search on Cults3D | For rosettes 32-56mm | - |

### Wall Protectors

| Model | Author | URL | Description | Last Verified |
|-------|--------|-----|-------------|---------------|
| **Wall Bumper v2/v3** | - | https://cults3d.com/en/tags/nuki | Prevents wall damage | - |
| **Door Stop** | - | Search Printables | Various sizes | - |

### Child Safety

| Model | Author | URL | Description | Last Verified |
|-------|--------|-----|-------------|---------------|
| **Child Saver 3.0/4.0** | - | https://cults3d.com/en/3d-model/home/nuki-smart-lock-3-0-4-0-child-saver-parental-controls | Prevents knob turning | Jan 2026 |

---

## Gaps - Models Needed

These adapters are **not currently available** and need to be designed:

| Gap | Market | Priority | Notes |
|-----|--------|----------|-------|
| **Lockwood oval (AU)** | Australia | 🔴 High | Most common AU residential lock |
| **MIWA LA-type** | Japan | ⚫ Very High (Hard) | Would require major engineering |
| **Local SEA brands** | Singapore, Malaysia | 🟡 Medium | Test and document what's needed |
| **UK night latch bypass** | UK | 🟡 Medium | For older British doors |
| **Corona cylinder** | Philippines | 🟢 Low | Local brand |

If you design one of these, please share it!

---

## File Formats Explained

| Format | Extension | Use Case | Editable? |
|--------|-----------|----------|-----------|
| **STL** | `.stl` | 3D printing (most common) | Difficult |
| **OBJ** | `.obj` | 3D printing, rendering | Difficult |
| **STEP** | `.step`, `.stp` | CAD modification, CNC | ✅ Yes |
| **3MF** | `.3mf` | 3D printing (includes metadata) | Limited |
| **SCAD** | `.scad` | OpenSCAD parametric source | ✅ Yes |
| **F3D** | `.f3d` | Fusion 360 native | ✅ Yes |

**Recommendation**: Download STEP files when available for easier modification.

---

## Printing Guidelines

### Material Selection

| Material | Pros | Cons | Use For |
|----------|------|------|---------|
| **PLA** | Easy to print, cheap | Brittle, heat-sensitive | Prototypes, spacers |
| **PETG** | Durable, heat-resistant | Stringing | Final adapters |
| **ABS** | Strong, heat-resistant | Warping, fumes | High-stress parts |
| **TPU** | Flexible | Hard to print | Bumpers |

### Print Settings

```
Adapters (need precision):
- Layer height: 0.1mm - 0.15mm
- Infill: 40-60%
- Walls: 3-4
- Material: PETG

Spacers (need strength):
- Layer height: 0.2mm
- Infill: 40%
- Walls: 3
- Material: PLA or PETG

Wall bumpers (need flex):
- Layer height: 0.2mm
- Infill: 20%
- Walls: 2
- Material: TPU or PETG
```

---

## Design Your Own Adapter

### Step 1: Measure Your Cylinder

```
┌─────────────────────────────────────────┐
│                                         │
│         WHAT TO MEASURE                 │
│                                         │
│         ┌─────────────┐                 │
│         │   KNOB      │                 │
│         │  ┌─────┐    │                 │
│         │  │     │◄───┼── Inner Ø      │
│         │  └─────┘    │                 │
│         │◄───────────►│                 │
│         │   Outer Ø   │                 │
│         └──────┬──────┘                 │
│                │                         │
│         Height │                         │
│                ▼                         │
│                                         │
└─────────────────────────────────────────┘
```

Measure with calipers:
- **Inner diameter**: Where the key goes (typically 7-9mm)
- **Outer diameter**: Body of the knob (typically 22-28mm)
- **Height**: How tall the knob is

### Step 2: Get Nuki Interface Dimensions

The adapter must interface with the Nuki's mounting socket. Download an existing adapter STEP file and measure the Nuki-side interface.

### Step 3: CAD Software Options

| Software | Cost | Difficulty | Best For |
|----------|------|------------|----------|
| **TinkerCAD** | Free | Easy | Simple shapes |
| **Fusion 360** | Free (hobby) | Medium | Precise modeling |
| **OpenSCAD** | Free | Medium | Parametric (adjustable) |
| **FreeCAD** | Free | Medium | Open source |
| **SolidWorks** | $$$ | Hard | Professional |

### Step 4: Test and Iterate

1. Print at 0.2mm layer height first (faster)
2. Test fit on cylinder
3. Adjust dimensions
4. Print final at 0.1mm layer height
5. Test with Nuki

### Step 5: Share Your Design!

Upload to:
- Printables.com (recommended)
- Cults3D
- Thingiverse
- MakerWorld

Include:
- Photos of it installed
- Measurements
- Print settings
- Which cylinder brand it fits

---

## Folder Structure

```
docs/adapters/3d/
├── README.md           # This file
├── adapters/           # Downloaded adapter STL/STEP files
│   └── .gitkeep
├── spacers/            # Downloaded spacer files
│   └── .gitkeep
└── accessories/        # Wall bumpers, child locks, etc.
    └── .gitkeep
```

Place downloaded files in the appropriate subfolder.

---

## See Also

- [Adapters Guide](../README.md) - Official adapters and visual guide
- [Door Lock Types](../../door-lock-types.md) - Understanding different locks
- [Nuki API](../../nuki-api.md) - Software integration
