package main

import (
	"context"
	"fmt"

	"github.com/GopherGhaznix/Wave/html"
)

// -----------------------
// Example Usage
// -----------------------
func main() {

	c := context.Background()

	root := html.Div(c,
		// Attributes
		html.Attrs{"id": "root", "class": "container"},
		// Childerns
		html.H1(c,
			// Attributes
			html.Attrs{"class": "title"},
			// Childerns
			html.Text("Hello Wave 🌊"),
		),

		html.P(c,
			// Attributes
			html.Attrs{
				"class": "para",
			},
			// Childerns
			html.Text("This is a paragraph inside Wave."),
		),

		html.Div(c,
			// Attributes
			html.Attrs{
				"class": "child",
			},
			// Childerns
			html.Span(c, nil, html.Text("Nested span text")),
			html.Em(c, nil, html.Text("and emphasized text.")),
			html.Ul(c,
				nil,
				// Childerns
				html.Li(c, nil, html.Text("Item 1")),
				html.Li(c, nil, html.Text("Item 2")),
				html.Li(c, nil, html.A(c, html.Attrs{"href": "#"}, html.Text("Link inside list"))),
			),
		),
		html.Input(c,
			// Attributes
			html.Attributes(html.Attrs{"id": "name", "type": "text"}),
		),
		html.Button(c,
			// Attributes
			html.Attributes(html.Attrs{"class": "btn-primary-1"}),
			// Childerns
			html.Text("Click Me"),
		),
	)

	// render final HTML
	fmt.Printf(`<script src="https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"></script>\n %s`, root())
}
