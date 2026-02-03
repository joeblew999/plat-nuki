---
title: "Country Guides"
date: 2024-01-15
draft: false
---

# Country Guides

Per-country compatibility guides for Nuki Smart Lock deployment.

---

## Official Nuki Markets

These countries have dedicated Nuki websites and official support.

### Europe (Smart Lock Pro / Go / Ultra)

| Country | Code | Lock Type | Guide | Status |
|---------|------|-----------|-------|--------|
| **Austria** | AT | Euro cylinder | [austria/](austria/) | ✅ Official |
| **Germany** | DE | Euro cylinder | [germany/](germany/) | ✅ Official |
| **Switzerland** | CH | Euro cylinder | [switzerland/](switzerland/) | ✅ Official |
| **United Kingdom** | GB | Euro cylinder + night latch | [united-kingdom/](united-kingdom/) | ✅ Official |
| **Netherlands** | NL | Euro cylinder | [netherlands/](netherlands/) | ✅ Official |
| **Belgium** | BE | Euro cylinder | [belgium/](belgium/) | ✅ Official |
| **France** | FR | Euro cylinder | [france/](france/) | ✅ Official |
| **Italy** | IT | Euro cylinder | [italy/](italy/) | ✅ Official |
| **Spain** | ES | Euro cylinder | [spain/](spain/) | ✅ Official |
| **Sweden** | SE | Scandinavian oval | [sweden/](sweden/) | ✅ Official (Ultra Nordics) |
| **Norway** | NO | Scandinavian oval | [norway/](norway/) | ✅ Official (Ultra Nordics) |
| **Denmark** | DK | Scandinavian oval | [denmark/](denmark/) | ✅ Official (Ultra Nordics) |
| **Finland** | FI | Scandinavian oval/round | [finland/](finland/) | ✅ Official (Ultra Nordics) |

### North America (Smart Lock US)

| Country | Code | Lock Type | Guide | Status |
|---------|------|-----------|-------|--------|
| **United States** | US | Deadbolt | [united-states/](united-states/) | ✅ Official (2025) |
| **Canada** | CA | Deadbolt | [canada/](canada/) | Ships from US |

### International (en-001)

Nuki's "International" site (en-001) ships to additional countries.

---

## Asia-Pacific Markets (Not Official - Our Focus)

These markets don't have dedicated Nuki support, but Nuki Pro works with euro cylinders common in these regions.

| Country | Lock Type | Compatible? | Guide | Notes |
|---------|-----------|-------------|-------|-------|
| **Singapore** | Euro cylinder | ✅ Yes | [singapore/](singapore/) | HDB + condos use euro |
| **Malaysia** | Euro cylinder | ✅ Yes | [malaysia/](malaysia/) | Similar to Singapore |
| **Hong Kong** | Euro cylinder | ✅ Yes | [hong-kong/](hong-kong/) | Commercial standard |
| **Thailand** | Euro cylinder | ✅ Yes | [thailand/](thailand/) | Condos use euro |
| **Australia** | Lockwood/Euro | ⚠️ Partial | [australia/](australia/) | **Gap**: Lockwood 001 |
| **New Zealand** | Lockwood/Euro | ⚠️ Partial | [new-zealand/](new-zealand/) | Same as Australia |
| **Philippines** | Mixed | ⚠️ Check | [philippines/](philippines/) | Deadbolt + euro mix |
| **Indonesia** | Euro/Mixed | ⚠️ Check | [indonesia/](indonesia/) | Condos euro, houses vary |

### Not Compatible

| Country | Lock Type | Why Not | Alternative |
|---------|-----------|---------|-------------|
| **Japan** | MIWA mortise | Different standard | Qrio, Sesame |
| **South Korea** | Digital mortise | Already digital | Samsung, EPIC |

---

## Which Nuki Model for Which Region?

```
┌─────────────────────────────────────────────────────────────┐
│                    NUKI MODEL SELECTOR                       │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  REGION                  LOCK TYPE           NUKI MODEL     │
│  ──────────────────────────────────────────────────────────  │
│                                                              │
│  USA, Canada             Deadbolt            Smart Lock US   │
│                                                              │
│  Western Europe          Euro cylinder       Smart Lock Pro  │
│  (DE, AT, CH, NL,                           or Go or Ultra  │
│   BE, FR, IT, ES)                                           │
│                                                              │
│  UK, Ireland             Euro cylinder       Smart Lock Pro  │
│                          + night latch       (euro doors)   │
│                                                              │
│  Scandinavia             Oval cylinder       Smart Lock     │
│  (SE, NO, DK, FI)                           Ultra Nordics   │
│                                                              │
│  Southeast Asia          Euro cylinder       Smart Lock Pro  │
│  (SG, MY, HK, TH)                                           │
│                                                              │
│  Australia (apts)        Euro cylinder       Smart Lock Pro  │
│                                                              │
│  Australia (houses)      Lockwood 001        ❌ Not compat.  │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

---

## Folder Structure

Each country guide follows this structure:

```
countries/
├── README.md              # This index
├── singapore/
│   ├── README.md          # Country guide
│   ├── images/            # Country-specific photos
│   │   └── README.md      # Image provenance
│   └── 3d/                # Country-specific 3D models
│       └── README.md      # Model sources
├── australia/
│   ├── README.md
│   ├── images/
│   │   └── lockwood-001-rimlock.jpg
│   └── 3d/
├── germany/
│   └── ...
└── ...
```

---

## Adding a New Country

1. Create folder: `countries/{country-name}/`
2. Add subfolders: `images/`, `3d/`
3. Create `README.md` with:
   - Status (Official/Unofficial)
   - Lock types found in that country
   - Adapter requirements
   - Where to buy
   - Installation considerations
   - Local regulations if any
4. Add country-specific images to `images/`
5. Link any relevant 3D models in `3d/`

---

## See Also

- [Adapters Guide](../) - Main adapters documentation
- [Visual Diagrams](../diagrams/) - Lock type photos
- [Asia/Australia Gaps](../ASIA-AUSTRALIA-GAPS/) - APAC market analysis
- [Door Lock Types](../../door-lock-types/) - Technical lock guide
