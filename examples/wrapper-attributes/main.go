package main

import (
	"fmt"

	"github.com/GopherGhaznix/Wave/html"
)

// -----------------------
// Example Usage
// -----------------------
func main() {

	root := html.Div(
		html.Attributes(
			// Attributes
			html.AttrID("root"),
			html.AttrClass("container"),
		),
		// Childerns
		html.H1(
			// Attributes
			html.AttrClass("title"),
			// Childerns
			html.Text("Hello Wave 🌊"),
		),
		html.P(
			// Attributes
			html.AttrClass("text-gray-600"),
			// Childerns
			html.Text("This is a paragraph inside Wave."),
		),
		html.Div(
			// Attributes
			html.AttrClass("child"),
			// Childerns
			html.Span(nil, html.Text("Nested span text")),
			html.Em(nil, html.Text("and emphasized text.")),
			html.Ul(nil,
				// Childerns
				html.Li(nil, html.Text("Item 1")),
				html.Li(nil, html.Text("Item 2")),
				html.Li(nil, html.A(html.AttrHref("#"), html.Text("Link inside list"))),
			),
		),
		html.Input(
			// Attributes
			html.Attributes(html.AttrID("name"), html.AttrTypeText()),
		),
		html.Button(
			// Attributes
			html.Attributes(html.AttrClass("btn-primary-1")),
			// Childerns
			html.Text("Click Me"),
		),
	)

	fmt.Println(root()) // render final HTML
}
