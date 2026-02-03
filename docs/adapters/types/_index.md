---
title: "Adapter Types"
date: 2024-01-15
draft: false
---

# Adapter Types

Adapters organized by type, with country mappings. This avoids duplication - each adapter is documented once, countries reference the adapter they need.

---

## Adapter Index

### European Knob Cylinder Adapters (8 colors)

For **Smart Lock Pro / Go / Ultra** with euro profile knob cylinders.

| Adapter | Color | Folder | Fits Brands | Countries |
|---------|-------|--------|-------------|-----------|
| **Abus/CISA** | Black | [black-abus/](black-abus/) | Abus, CISA, Bricard, AXA, ISEO | DE, AT, CH, FR, IT, NL, BE, SG, MY, HK, TH |
| **CES/EVVA** | Green | [green-evva/](green-evva/) | CES, EVVA | DE, AT, CH |
| **DOM/Winkhaus** | Blue | [blue-dom/](blue-dom/) | DOM, Winkhaus | DE, NL, BE |
| **Dormakaba** | Yellow | [yellow-dormakaba/](yellow-dormakaba/) | Dormakaba, Kaba | CH, DE, SG, MY, HK, TH, Commercial worldwide |
| **KESO** | Red | [red-keso/](red-keso/) | KESO | CH |
| **Nemef** | White | [white-nemef/](white-nemef/) | Nemef | NL, BE |
| **Yale/M&C** | Magenta | [magenta-yale/](magenta-yale/) | Yale, M&C | UK, SG, MY, TH, Worldwide |
| **Mauer** | Orange | [orange-mauer/](orange-mauer/) | Mauer | DE |

### US Deadbolt Adapters (3 included)

For **Smart Lock (US)** - all 3 included in box.

| Adapter | Shape | Folder | Fits Brands | Countries |
|---------|-------|--------|-------------|-----------|
| **Square** | ▢ | [us-square/](us-square/) | Schlage, some Kwikset | US, CA |
| **D-Shape** | ◗◖ | [us-dshape/](us-dshape/) | Most US brands | US, CA |
| **Slotted** | ═ | [us-slotted/](us-slotted/) | Older locks | US, CA |

### Scandinavian (Universal Cylinder)

For **Smart Lock Ultra Nordics** - replaces cylinder entirely.

| Adapter | Type | Folder | Fits | Countries |
|---------|------|--------|------|-----------|
| **Universal Cylinder** | Oval/Round | [nordic-universal/](nordic-universal/) | ASSA, Abloy, Ruko, Trioving | SE, NO, DK, FI |

### Custom / 3D Printed

For gaps not covered by official adapters.

| Adapter | Type | Folder | Fits | Countries |
|---------|------|--------|------|-----------|
| **Lockwood Oval** | 3D Print | [custom-lockwood-oval/](custom-lockwood-oval/) | Lockwood oval cylinders | AU, NZ |
| **Generic Spacer** | 3D Print | [custom-spacer/](custom-spacer/) | Better adhesion | All |

---

## Country → Adapter Mapping

Quick reference: which adapter(s) for which country.

### Official Markets

| Country | Primary Adapter | Secondary | Notes |
|---------|-----------------|-----------|-------|
| **Germany** | Black, Blue | Green, Yellow | DOM/Winkhaus common |
| **Austria** | Green, Black | Yellow | EVVA is Austrian |
| **Switzerland** | Yellow, Red | Black, Green | Kaba/KESO are Swiss |
| **Netherlands** | White | Magenta, Black | Nemef is Dutch |
| **Belgium** | White, Black | Magenta | Mix of NL/FR brands |
| **France** | Black | Blue | Bricard common |
| **Italy** | Black | - | ISEO/CISA Italian |
| **Spain** | Black, Magenta | - | Mix |
| **UK** | Magenta | Black | Yale dominant |
| **Sweden** | Nordic Universal | - | Scandinavian oval |
| **Norway** | Nordic Universal | - | Scandinavian oval |
| **Denmark** | Nordic Universal | - | Scandinavian oval |
| **Finland** | Nordic Universal | - | Scandinavian oval |
| **USA** | US (all 3 included) | - | Deadbolt |
| **Canada** | US (all 3 included) | - | Deadbolt |

### APAC Markets

| Country | Primary Adapter | Secondary | Notes |
|---------|-----------------|-----------|-------|
| **Singapore** | Yellow, Magenta | Black | Dormakaba, Yale common |
| **Malaysia** | Magenta, Yellow | Black | Yale, Häfele common |
| **Hong Kong** | Yellow, Magenta | Black | Commercial euro |
| **Thailand** | Magenta, Yellow | Black | Yale common |
| **Australia** | Magenta, Black | Custom Lockwood | Euro apts only |
| **New Zealand** | Magenta, Black | Custom Lockwood | Same as AU |

---

## Folder Structure

```
adapters/
├── types/
│   ├── README.md                 # This file (index)
│   │
│   ├── black-abus/               # EU Knob Adapter - Black
│   │   ├── README.md
│   │   ├── images/
│   │   └── 3d/
│   │
│   ├── green-evva/               # EU Knob Adapter - Green
│   │   └── ...
│   │
│   ├── blue-dom/                 # EU Knob Adapter - Blue
│   ├── yellow-dormakaba/         # EU Knob Adapter - Yellow
│   ├── red-keso/                 # EU Knob Adapter - Red
│   ├── white-nemef/              # EU Knob Adapter - White
│   ├── magenta-yale/             # EU Knob Adapter - Magenta
│   ├── orange-mauer/             # EU Knob Adapter - Orange
│   │
│   ├── us-square/                # US Deadbolt - Square
│   ├── us-dshape/                # US Deadbolt - D-shape
│   ├── us-slotted/               # US Deadbolt - Slotted
│   │
│   ├── nordic-universal/         # Scandinavian Universal Cylinder
│   │
│   └── custom-lockwood-oval/     # 3D Printable - AU/NZ gap
│
└── countries/
    ├── singapore/
    │   └── README.md             # Links to types/yellow-dormakaba, types/magenta-yale
    └── ...
```

---

## How This Works

1. **Adapter documentation** lives in `types/{adapter-name}/`
   - Images of the adapter
   - 3D print files if available
   - Which cylinder brands it fits
   - Installation notes

2. **Country guides** in `countries/{country}/` link to the relevant adapter types
   - No duplicate adapter info
   - Country-specific buying info, regulations, etc.
   - References shared adapter documentation

3. **Adding a new country**:
   - Create country folder
   - Link to existing adapter types
   - Only add country-specific content

4. **Adding a new adapter**:
   - Create adapter type folder
   - Document once
   - Update country mappings

---

## See Also

- [Country Guides](../countries/)
- [Visual Diagrams](../diagrams/)
- [3D Models](../3d/)
