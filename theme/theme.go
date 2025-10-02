package theme

import "github.com/GopherGhaznix/Wave/html"

// Theme defines reusable style tokens and class utilities
type Theme struct {
	Colors struct {
		Primary   string
		Secondary string
		Success   string
		Danger    string
		Warning   string
		Info      string
		Light     string
		Dark      string
	}

	Typography struct {
		Base     string
		Heading  string
		Subtitle string
		Muted    string
		Code     string
	}

	Layout struct {
		Container string
		Section   string
		Card      string
		Button    string
		Input     string
	}

	Variants map[string]html.Attrs
}
