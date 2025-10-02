# 🌊 Wave
Fluid pages, powered by Go at the core.

Wave is a declarative, composable **HTML builder for Go**, designed to let you generate semantic, type-safe markup without templates.  

---

## ⚠️ Project Status

🚧 **Under Development** — Wave is experimental and evolving quickly.  
The API is not stable yet, and things may change. Contributions and feedback are welcome!

---

## ✨ Features

- **Declarative HTML**: Write Go functions instead of raw strings.
- **Composable**: Nest elements and attributes just like native HTML.
- **Typed Attributes**: Use helpers like `html.AttrID("id")`, `html.AttrClass("btn")`, `html.AttrTypeText()` to avoid typos.
- **Automatic IDs**: Elements get unique IDs if you don’t provide one.
- **Pretty Printing**: Indentation for nested elements is handled automatically.
- **Full HTML5 Coverage**: Includes wrappers for nearly all HTML5 elements and attributes.

---

## 🚀 Getting Started

Install the package:

```bash
go get github.com/GopherGhaznix/Wave
```