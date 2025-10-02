package css

import (
	"sort"
	"strings"
)

type Style map[string]string

func (s Style) Inline() string {
	if len(s) == 0 {
		return ""
	}

	// Sort keys for deterministic output
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for _, k := range keys {
		v := s[k]
		sb.WriteString(k)
		sb.WriteString(":")
		sb.WriteString(v)
		sb.WriteString(";")
	}

	return sb.String()
}
