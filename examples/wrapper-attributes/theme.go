package main

import (
	"context"

	"github.com/GopherGhaznix/Wave/html"
)

func MyCustomTailwindTheme() context.Context {
	return html.WithTheme(
		context.Background(),
		html.NewTheme(
			map[string]html.Attrs{
				// ----------------
				// Text & Inline
				// ----------------
				"p":      html.AttrClass("mb-4 text-base leading-relaxed text-gray-800"),
				"span":   html.AttrClass(""),
				"small":  html.AttrClass("text-sm text-gray-500"),
				"strong": html.AttrClass("font-semibold text-gray-900"),
				"em":     html.AttrClass("italic"),
				"mark":   html.AttrClass("bg-yellow-200 text-gray-900 px-1 rounded"),
				"code":   html.AttrClass("font-mono text-sm bg-gray-100 px-1 py-0.5 rounded text-red-600"),

				// ----------------
				// Headings
				// ----------------
				"h1": html.AttrClass("text-4xl font-extrabold mb-6 tracking-tight text-gray-900"),
				"h2": html.AttrClass("text-3xl font-bold mb-5 tracking-tight text-gray-900"),
				"h3": html.AttrClass("text-2xl font-semibold mb-4 tracking-tight text-gray-900"),
				"h4": html.AttrClass("text-xl font-semibold mb-3 text-gray-900"),
				"h5": html.AttrClass("text-lg font-medium mb-2 text-gray-800"),
				"h6": html.AttrClass("text-base font-medium text-gray-700"),

				// ----------------
				// Links
				// ----------------
				"a": html.AttrClass("text-blue-600 hover:text-blue-700 hover:underline transition-colors"),

				// ----------------
				// Buttons
				// ----------------
				"button": html.AttrClass("inline-flex items-center px-4 py-2 rounded-lg bg-blue-600 text-white font-medium hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"),
				"submit": html.AttrClass("inline-flex items-center px-4 py-2 rounded-lg bg-green-600 text-white font-medium hover:bg-green-700 transition-colors"),

				// ----------------
				// Forms
				// ----------------
				"form":     html.AttrClass("space-y-4"),
				"label":    html.AttrClass("block mb-1 text-sm font-medium text-gray-700"),
				"input":    html.AttrClass("border border-gray-300 bg-white mb-2 px-3 py-2 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"),
				"select":   html.AttrClass("border border-gray-300 bg-white mb-2 px-3 py-2 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900"),
				"textarea": html.AttrClass("border border-gray-300 bg-white mb-2 px-3 py-2 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 text-gray-900 resize-y"),

				// ----------------
				// Lists
				// ----------------
				"ul": html.AttrClass("list-disc pl-6 mb-4 space-y-1 text-gray-800"),
				"ol": html.AttrClass("list-decimal pl-6 mb-4 space-y-1 text-gray-800"),
				"li": html.AttrClass("leading-relaxed"),

				// ----------------
				// Tables
				// ----------------
				"table": html.AttrClass("min-w-full border-collapse border border-gray-200 rounded-lg overflow-hidden"),
				"thead": html.AttrClass("bg-gray-50"),
				"tbody": html.AttrClass("divide-y divide-gray-200"),
				"tr":    html.AttrClass("hover:bg-gray-50 transition-colors"),
				"th":    html.AttrClass("px-4 py-2 text-left font-medium text-gray-700"),
				"td":    html.AttrClass("px-4 py-2 text-gray-700"),

				// ----------------
				// Media
				// ----------------
				"img":   html.AttrClass("max-w-full h-auto rounded-lg shadow-sm"),
				"video": html.AttrClass("max-w-full rounded-lg"),

				// ----------------
				// Structural Containers
				// ----------------
				"div":     html.AttrClass(""),
				"section": html.AttrClass("mb-8"),
				"article": html.AttrClass("mb-8"),
				"header":  html.AttrClass("mb-6"),
				"footer":  html.AttrClass("mt-6 text-sm text-gray-500"),
				"main":    html.AttrClass("flex-1"),
				"nav":     html.AttrClass("flex space-x-4"),

				// ----------------
				// Semantic Elements
				// ----------------
				"aside":      html.AttrClass("text-sm text-gray-600 p-4 bg-gray-50 rounded-lg"),
				"figure":     html.AttrClass("my-6"),
				"figcaption": html.AttrClass("mt-2 text-sm text-gray-500 text-center"),
				"blockquote": html.AttrClass("border-l-4 border-blue-500 pl-4 italic text-gray-700 my-4"),
				"hr":         html.AttrClass("my-6 border-gray-300"),
				"pre":        html.AttrClass("bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto font-mono text-sm"),
			},
		),
	)
}
