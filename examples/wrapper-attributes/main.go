package main

import (
	"fmt"

	"github.com/GopherGhaznix/Wave/css"
	"github.com/GopherGhaznix/Wave/html"
)

// -----------------------
// Example Usage
// -----------------------
func main() {

	// default theme context
	// c := html.WithTheme(
	// 	context.Background(),
	// 	html.NewDefaultTheme(),
	// )

	// custom theme context
	c := MyCustomTailwindTheme()

	// HTML Stuctrue
	root := html.Div(c,
		html.Attributes(
			// Attributes
			html.AttrID("root"),
			html.AttrClass("container bg-blue-200"),
			html.AttrStyle(css.Style{
				"padding": "30px",
				"height":  "90vh",
				"border":  "1px solid #eee",
			}),
		),
		// Childerns
		html.H1(c,
			// Attributes
			html.AttrClass("title"),
			// Childerns
			html.Text("Hello Wave 🌊"),
		),
		html.P(c,
			// Attributes
			html.AttrClass("para"),
			// Childerns
			html.Text("This is a paragraph inside Wave."),
		),
		html.Div(c,
			// Attributes
			html.AttrClass("child"),
			// Childerns
			html.Span(c, nil, html.Text("Nested span text")),
			html.Em(c, nil, html.Text("and emphasized text.")),
			html.Ul(c, nil,
				// Childerns
				html.Li(c, nil, html.Text("Item 1")),
				html.Li(c, nil, html.Text("Item 2")),
				html.Li(c, nil, html.A(c, html.AttrHref("#"), html.Text("Link inside list"))),
			),
		),
		html.Input(c,
			// Attributes
			html.Attributes(html.AttrID("name"), html.AttrTypeText()),
		),
		html.Button(c,
			// Attributes
			html.Attributes(html.AttrClass("btn-primary-1")),
			// Childerns
			html.Text("Click Me"),
		),
	)

	// render final HTML
	fmt.Printf("<script src='https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4'></script>\n %s", root())
}
