package html

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/google/uuid"
)

// -----------------------
// Types
// -----------------------
type Node func() string

// Text helper: wrap string into Node
func Text(s string) Node {
	return func() string { return s }
}

// Indent helper for pretty printing
func indentBlock(content string, level int) string {
	prefix := strings.Repeat("  ", level)
	lines := strings.Split(strings.TrimSuffix(content, "\n"), "\n")
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

// mergeAttrs merges theme attrs with user attrs.
// - Non-class attributes: user overrides theme.
// - class attribute: merge intelligently, replacing conflicts.
func mergeAttrs(themeAttrs, userAttrs Attrs) Attrs {
	result := Attrs{}

	// Copy theme defaults first
	for k, v := range themeAttrs {
		result[k] = v
	}

	// Apply user overrides
	for k, v := range userAttrs {
		if k == "class" {
			result[k] = mergeClasses(themeAttrs["class"], v)
		} else {
			result[k] = v
		}
	}

	return result
}

// mergeClasses merges theme classes with user classes.
// User classes replace theme classes that share the same "prefix" (e.g. bg-*, text-*, font-*).
func mergeClasses(themeClass, userClass string) string {
	themeList := strings.Fields(themeClass)
	userList := strings.Fields(userClass)

	// Build a set of user class "prefixes"
	// Example: "bg-red-500" → prefix "bg-"
	userPrefixes := map[string]string{}
	for _, uc := range userList {
		prefix := classPrefix(uc)
		userPrefixes[prefix] = uc
	}

	merged := []string{}

	// Keep theme classes unless overridden
	for _, tc := range themeList {
		prefix := classPrefix(tc)
		if uc, ok := userPrefixes[prefix]; ok {
			merged = append(merged, uc) // user override
			delete(userPrefixes, prefix)
		} else {
			merged = append(merged, tc) // keep theme default
		}
	}

	// Add any remaining user classes (new, non-conflicting ones)
	for _, uc := range userList {
		prefix := classPrefix(uc)
		if _, ok := userPrefixes[prefix]; ok {
			merged = append(merged, uc)
			delete(userPrefixes, prefix)
		}
	}

	return strings.Join(merged, " ")
}

// classPrefix extracts a "group" prefix for a class.
// Example: "bg-red-500" → "bg-", "text-lg" → "text-", "rounded" → "rounded".
func classPrefix(c string) string {
	if i := strings.Index(c, "-"); i > 0 {
		return c[:i+1] // include dash, e.g. "bg-"
	}
	return c
}

// -----------------------
// Core Element Builder
// -----------------------
func Element(c context.Context, tag string, attrs Attrs, children ...Node) Node {
	return func() string {
		// Try to get theme from context (may be nil)
		theme, _ := ThemeFromContext(c)

		if len(attrs) == 0 {
			attrs = Attrs{}
		}

		// Only apply theme styles if a theme exists
		if theme != nil {
			if defaultAttrs, ok := (*theme)[tag]; ok {
				attrs = mergeAttrs(defaultAttrs, attrs) // theme first, user overrides
			}
		}

		// Auto-generate id if missing
		if _, ok := attrs["id"]; !ok {
			if uid, err := uuid.NewV7(); err == nil {
				attrs["id"] = uid.String()
			}
		}

		// Render children
		var expandedChildren []string
		for _, child := range children {
			if child != nil {
				expandedChildren = append(expandedChildren, indentBlock(child(), 1))
			}
		}
		childrenHTML := strings.Join(expandedChildren, "\n")

		// Build attributes string
		attrParts := []string{}
		for k, v := range attrs {
			if v == "" {
				continue
			}
			attrParts = append(attrParts, fmt.Sprintf(`%s="%s"`, k, v))
		}
		sort.Strings(attrParts)
		attrStr := strings.Join(attrParts, " ")

		if childrenHTML == "" {
			return fmt.Sprintf(`<%s %s />`, tag, attrStr)
		}
		return fmt.Sprintf(`<%s %s>
%s
</%s>`, tag, attrStr, childrenHTML, tag)
	}
}

// -----------------------
// HTML5 Elements Wrappers
// -----------------------

// Structural
func A(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "a", attrs, children...)
}
func Abbr(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "abbr", attrs, children...)
}
func Acronym(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "acronym", attrs, children...)
}
func Address(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "address", attrs, children...)
}
func Area(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "area", attrs, children...)
}
func Article(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "article", attrs, children...)
}
func Aside(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "aside", attrs, children...)
}
func Audio(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "audio", attrs, children...)
}
func B(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "b", attrs, children...)
}
func Base(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "base", attrs, children...)
}
func Bdi(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "bdi", attrs, children...)
}
func Bdo(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "bdo", attrs, children...)
}
func Big(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "big", attrs, children...)
}
func Blockquote(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "blockquote", attrs, children...)
}
func Body(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "body", attrs, children...)
}
func Br(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "br", attrs, children...)
}
func Button(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "button", attrs, children...)
}
func Canvas(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "canvas", attrs, children...)
}
func Caption(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "caption", attrs, children...)
}
func Cite(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "cite", attrs, children...)
}
func Code(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "code", attrs, children...)
}
func Col(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "col", attrs, children...)
}
func Colgroup(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "colgroup", attrs, children...)
}
func Data(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "data", attrs, children...)
}
func Datalist(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "datalist", attrs, children...)
}
func Dd(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dd", attrs, children...)
}
func Del(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "del", attrs, children...)
}
func Details(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "details", attrs, children...)
}
func Dfn(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dfn", attrs, children...)
}
func Dialog(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dialog", attrs, children...)
}
func Dir(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dir", attrs, children...)
}
func Div(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "div", attrs, children...)
}
func Dl(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dl", attrs, children...)
}
func Dt(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "dt", attrs, children...)
}
func Em(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "em", attrs, children...)
}
func Embed(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "embed", attrs, children...)
}
func Fieldset(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "fieldset", attrs, children...)
}
func Figcaption(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "figcaption", attrs, children...)
}
func Figure(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "figure", attrs, children...)
}
func Footer(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "footer", attrs, children...)
}
func Form(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "form", attrs, children...)
}
func H1(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h1", attrs, children...)
}
func H2(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h2", attrs, children...)
}
func H3(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h3", attrs, children...)
}
func H4(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h4", attrs, children...)
}
func H5(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h5", attrs, children...)
}
func H6(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "h6", attrs, children...)
}
func Head(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "head", attrs, children...)
}
func Header(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "header", attrs, children...)
}
func Hgroup(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "hgroup", attrs, children...)
}
func Hr(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "hr", attrs, children...)
}
func Html(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "html", attrs, children...)
}
func I(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "i", attrs, children...)
}
func Iframe(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "iframe", attrs, children...)
}
func Img(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "img", attrs, children...)
}
func Input(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "input", attrs, children...)
}
func Ins(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "ins", attrs, children...)
}
func Kbd(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "kbd", attrs, children...)
}
func Label(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "label", attrs, children...)
}
func Legend(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "legend", attrs, children...)
}
func Li(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "li", attrs, children...)
}
func Link(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "link", attrs, children...)
}
func Main(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "main", attrs, children...)
}
func Map(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "map", attrs, children...)
}
func Mark(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "mark", attrs, children...)
}
func Menu(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "menu", attrs, children...)
}
func Meta(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "meta", attrs, children...)
}
func Meter(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "meter", attrs, children...)
}
func Nav(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "nav", attrs, children...)
}
func Nobr(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "nobr", attrs, children...)
}
func Noembed(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "noembed", attrs, children...)
}
func Noframes(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "noframes", attrs, children...)
}
func Noscript(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "noscript", attrs, children...)
}
func Object(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "object", attrs, children...)
}
func Ol(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "ol", attrs, children...)
}
func Optgroup(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "optgroup", attrs, children...)
}
func Option(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "option", attrs, children...)
}
func Output(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "output", attrs, children...)
}
func P(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "p", attrs, children...)
}
func Param(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "param", attrs, children...)
}
func Picture(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "picture", attrs, children...)
}
func Pre(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "pre", attrs, children...)
}
func Progress(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "progress", attrs, children...)
}
func Q(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "q", attrs, children...)
}
func Rb(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "rb", attrs, children...)
}
func Rp(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "rp", attrs, children...)
}
func Rt(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "rt", attrs, children...)
}
func Ruby(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "ruby", attrs, children...)
}
func S(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "s", attrs, children...)
}
func Samp(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "samp", attrs, children...)
}
func Script(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "script", attrs, children...)
}
func Section(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "section", attrs, children...)
}
func Select(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "select", attrs, children...)
}
func Small(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "small", attrs, children...)
}
func Source(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "source", attrs, children...)
}
func Span(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "span", attrs, children...)
}
func Strong(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "strong", attrs, children...)
}
func Style(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "style", attrs, children...)
}
func Sub(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "sub", attrs, children...)
}
func Summary(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "summary", attrs, children...)
}
func Sup(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "sup", attrs, children...)
}
func Table(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "table", attrs, children...)
}
func Tbody(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "tbody", attrs, children...)
}
func Td(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "td", attrs, children...)
}
func Template(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "template", attrs, children...)
}
func Textarea(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "textarea", attrs, children...)
}
func Tfoot(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "tfoot", attrs, children...)
}
func Th(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "th", attrs, children...)
}
func Thead(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "thead", attrs, children...)
}
func Time(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "time", attrs, children...)
}
func Title(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "title", attrs, children...)
}
func Tr(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "tr", attrs, children...)
}
func Track(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "track", attrs, children...)
}
func U(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "u", attrs, children...)
}
func Ul(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "ul", attrs, children...)
}
func Var(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "var", attrs, children...)
}
func Video(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "video", attrs, children...)
}
func Wbr(c context.Context, attrs Attrs, children ...Node) Node {
	return Element(c, "wbr", attrs, children...)
}
