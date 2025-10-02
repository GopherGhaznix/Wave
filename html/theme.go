package html

import (
	"context"
)

// unexported key type ensures uniqueness
type themeContextKey struct{}

// Theme maps an HTML tag name (e.g. "p", "h1", "button")
// to a set of default attributes (Attrs) that define its
// styling and behavior. These defaults can be applied
// automatically when rendering elements, while still
// allowing user-specified attributes to override them.
//
// Example:
//
//	Theme{
//	  "p": AttrClass("mb-4 text-base text-gray-800"),
//	  "a": AttrClass("text-blue-600 hover:underline"),
//	}
type Theme map[string]Attrs

// WithTheme returns a new context carrying the given theme.
func WithTheme(ctx context.Context, theme *Theme) context.Context {
	return context.WithValue(ctx, themeContextKey{}, theme)
}

// ThemeFromContext retrieves the theme from context, if set.
func ThemeFromContext(ctx context.Context) (*Theme, bool) {
	t, ok := ctx.Value(themeContextKey{}).(*Theme)
	return t, ok
}

// NewTheme creates a Theme from a given map of element styles.
func NewTheme(elementStyle map[string]Attrs) *Theme {
	t := Theme(elementStyle)
	return &t
}

// NewDefaultTheme returns a Theme with sensible defaults for major HTML elements.
func NewDefaultTheme() *Theme {
	t := map[string]Attrs{
		// Text
		"p":     AttrClass("mb-4 text-base leading-relaxed text-gray-800"),
		"span":  AttrClass(""),
		"small": AttrClass("text-sm text-gray-500"),
		"code":  AttrClass("font-mono text-sm bg-gray-100 px-1 py-0.5 rounded"),

		// Headings
		"h1": AttrClass("text-4xl font-bold mb-4 text-gray-900"),
		"h2": AttrClass("text-3xl font-semibold mb-3 text-gray-900"),
		"h3": AttrClass("text-2xl font-semibold mb-2 text-gray-900"),
		"h4": AttrClass("text-xl font-semibold mb-2 text-gray-900"),
		"h5": AttrClass("text-lg font-semibold mb-1 text-gray-900"),
		"h6": AttrClass("text-base font-semibold text-gray-800"),

		// Links
		"a": AttrClass("text-blue-500 hover:underline"),

		// Buttons
		"button": AttrClass("px-4 py-2 rounded bg-blue-500 text-white hover:bg-blue-600 disabled:opacity-50 disabled:cursor-not-allowed"),
		"submit": AttrClass("px-4 py-2 rounded bg-green-500 text-white hover:bg-green-600"),

		// Forms
		"form":     AttrClass("space-y-4"),
		"label":    AttrClass("mb-2 block text-sm font-medium text-gray-700"),
		"input":    AttrClass("border border-gray-300 bg-gray-100 mb-2 px-3 py-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-500"),
		"select":   AttrClass("border border-gray-300 mb-2 px-3 py-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-500"),
		"textarea": AttrClass("border border-gray-300 mb-2 px-3 py-2 rounded w-full focus:outline-none focus:ring-2 focus:ring-blue-500"),

		// Lists
		"ul": AttrClass("list-disc pl-5 mb-4"),
		"ol": AttrClass("list-decimal pl-5 mb-4"),
		"li": AttrClass("mb-1"),

		// Tables
		"table": AttrClass("min-w-full border-collapse border border-gray-200"),
		"thead": AttrClass("bg-gray-100"),
		"tbody": AttrClass("px-4"),
		"tr":    AttrClass("border-b border-gray-200"),
		"th":    AttrClass("px-4 py-2 text-left font-medium text-gray-600"),
		"td":    AttrClass("px-4 py-2 text-gray-700"),

		// Images & Media
		"img":   AttrClass("max-w-full h-auto rounded"),
		"video": AttrClass("max-w-full rounded"),

		// Containers
		"div":     AttrClass(""),
		"section": AttrClass("mb-8"),
		"article": AttrClass("mb-8"),
		"header":  AttrClass("mb-6"),
		"footer":  AttrClass("mt-6 text-sm text-gray-500"),
		"main":    AttrClass("flex-1"),
		"nav":     AttrClass("flex space-x-4"),
	}

	theme := Theme(t)
	return &theme
}
