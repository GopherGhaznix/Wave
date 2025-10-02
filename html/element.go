package html

import (
	"fmt"
	"maps"
	"sort"
	"strings"

	"github.com/google/uuid"
)

// -----------------------
// Types
// -----------------------
type Attrs map[string]string
type Node func() string

// Attributes returns a new Attrs map that combines all key-value pairs from the provided maps.
func Attributes(mapsList ...Attrs) Attrs {
	merged := Attrs{}
	for _, m := range mapsList {
		maps.Copy(merged, m)
	}
	return merged
}

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

// -----------------------
// Core Element Builder
// -----------------------
func Element(tag string, attrs Attrs, children ...Node) Node {
	return func() string {
		if attrs == nil {
			attrs = Attrs{}
		}

		// auto-generate id if missing
		if _, ok := attrs["id"]; !ok {
			if uid, err := uuid.NewV7(); err == nil {
				attrs["id"] = uid.String()
			}
		}

		// render children
		var expandedChildren []string
		for _, child := range children {
			if child != nil {
				expandedChildren = append(expandedChildren, indentBlock(child(), 1))
			}
		}
		childrenHTML := strings.Join(expandedChildren, "\n")

		// build attributes string
		attrParts := []string{}
		for k, v := range attrs {
			if v != "" {
				attrParts = append(attrParts, fmt.Sprintf(`%s="%s"`, k, v))
			}
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
func A(attrs Attrs, children ...Node) Node          { return Element("a", attrs, children...) }
func Abbr(attrs Attrs, children ...Node) Node       { return Element("abbr", attrs, children...) }
func Acronym(attrs Attrs, children ...Node) Node    { return Element("acronym", attrs, children...) }
func Address(attrs Attrs, children ...Node) Node    { return Element("address", attrs, children...) }
func Area(attrs Attrs, children ...Node) Node       { return Element("area", attrs, children...) }
func Article(attrs Attrs, children ...Node) Node    { return Element("article", attrs, children...) }
func Aside(attrs Attrs, children ...Node) Node      { return Element("aside", attrs, children...) }
func Audio(attrs Attrs, children ...Node) Node      { return Element("audio", attrs, children...) }
func B(attrs Attrs, children ...Node) Node          { return Element("b", attrs, children...) }
func Base(attrs Attrs, children ...Node) Node       { return Element("base", attrs, children...) }
func Bdi(attrs Attrs, children ...Node) Node        { return Element("bdi", attrs, children...) }
func Bdo(attrs Attrs, children ...Node) Node        { return Element("bdo", attrs, children...) }
func Big(attrs Attrs, children ...Node) Node        { return Element("big", attrs, children...) }
func Blockquote(attrs Attrs, children ...Node) Node { return Element("blockquote", attrs, children...) }
func Body(attrs Attrs, children ...Node) Node       { return Element("body", attrs, children...) }
func Br(attrs Attrs, children ...Node) Node         { return Element("br", attrs, children...) }
func Button(attrs Attrs, children ...Node) Node     { return Element("button", attrs, children...) }
func Canvas(attrs Attrs, children ...Node) Node     { return Element("canvas", attrs, children...) }
func Caption(attrs Attrs, children ...Node) Node    { return Element("caption", attrs, children...) }
func Cite(attrs Attrs, children ...Node) Node       { return Element("cite", attrs, children...) }
func Code(attrs Attrs, children ...Node) Node       { return Element("code", attrs, children...) }
func Col(attrs Attrs, children ...Node) Node        { return Element("col", attrs, children...) }
func Colgroup(attrs Attrs, children ...Node) Node   { return Element("colgroup", attrs, children...) }
func Data(attrs Attrs, children ...Node) Node       { return Element("data", attrs, children...) }
func Datalist(attrs Attrs, children ...Node) Node   { return Element("datalist", attrs, children...) }
func Dd(attrs Attrs, children ...Node) Node         { return Element("dd", attrs, children...) }
func Del(attrs Attrs, children ...Node) Node        { return Element("del", attrs, children...) }
func Details(attrs Attrs, children ...Node) Node    { return Element("details", attrs, children...) }
func Dfn(attrs Attrs, children ...Node) Node        { return Element("dfn", attrs, children...) }
func Dialog(attrs Attrs, children ...Node) Node     { return Element("dialog", attrs, children...) }
func Dir(attrs Attrs, children ...Node) Node        { return Element("dir", attrs, children...) }
func Div(attrs Attrs, children ...Node) Node        { return Element("div", attrs, children...) }
func Dl(attrs Attrs, children ...Node) Node         { return Element("dl", attrs, children...) }
func Dt(attrs Attrs, children ...Node) Node         { return Element("dt", attrs, children...) }
func Em(attrs Attrs, children ...Node) Node         { return Element("em", attrs, children...) }
func Embed(attrs Attrs, children ...Node) Node      { return Element("embed", attrs, children...) }
func Fieldset(attrs Attrs, children ...Node) Node   { return Element("fieldset", attrs, children...) }
func Figcaption(attrs Attrs, children ...Node) Node { return Element("figcaption", attrs, children...) }
func Figure(attrs Attrs, children ...Node) Node     { return Element("figure", attrs, children...) }
func Footer(attrs Attrs, children ...Node) Node     { return Element("footer", attrs, children...) }
func Form(attrs Attrs, children ...Node) Node       { return Element("form", attrs, children...) }
func H1(attrs Attrs, children ...Node) Node         { return Element("h1", attrs, children...) }
func H2(attrs Attrs, children ...Node) Node         { return Element("h2", attrs, children...) }
func H3(attrs Attrs, children ...Node) Node         { return Element("h3", attrs, children...) }
func H4(attrs Attrs, children ...Node) Node         { return Element("h4", attrs, children...) }
func H5(attrs Attrs, children ...Node) Node         { return Element("h5", attrs, children...) }
func H6(attrs Attrs, children ...Node) Node         { return Element("h6", attrs, children...) }
func Head(attrs Attrs, children ...Node) Node       { return Element("head", attrs, children...) }
func Header(attrs Attrs, children ...Node) Node     { return Element("header", attrs, children...) }
func Hgroup(attrs Attrs, children ...Node) Node     { return Element("hgroup", attrs, children...) }
func Hr(attrs Attrs, children ...Node) Node         { return Element("hr", attrs, children...) }
func Html(attrs Attrs, children ...Node) Node       { return Element("html", attrs, children...) }
func I(attrs Attrs, children ...Node) Node          { return Element("i", attrs, children...) }
func Iframe(attrs Attrs, children ...Node) Node     { return Element("iframe", attrs, children...) }
func Img(attrs Attrs, children ...Node) Node        { return Element("img", attrs, children...) }
func Input(attrs Attrs, children ...Node) Node      { return Element("input", attrs, children...) }
func Ins(attrs Attrs, children ...Node) Node        { return Element("ins", attrs, children...) }
func Kbd(attrs Attrs, children ...Node) Node        { return Element("kbd", attrs, children...) }
func Label(attrs Attrs, children ...Node) Node      { return Element("label", attrs, children...) }
func Legend(attrs Attrs, children ...Node) Node     { return Element("legend", attrs, children...) }
func Li(attrs Attrs, children ...Node) Node         { return Element("li", attrs, children...) }
func Link(attrs Attrs, children ...Node) Node       { return Element("link", attrs, children...) }
func Main(attrs Attrs, children ...Node) Node       { return Element("main", attrs, children...) }
func Map(attrs Attrs, children ...Node) Node        { return Element("map", attrs, children...) }
func Mark(attrs Attrs, children ...Node) Node       { return Element("mark", attrs, children...) }
func Menu(attrs Attrs, children ...Node) Node       { return Element("menu", attrs, children...) }
func Meta(attrs Attrs, children ...Node) Node       { return Element("meta", attrs, children...) }
func Meter(attrs Attrs, children ...Node) Node      { return Element("meter", attrs, children...) }
func Nav(attrs Attrs, children ...Node) Node        { return Element("nav", attrs, children...) }
func Nobr(attrs Attrs, children ...Node) Node       { return Element("nobr", attrs, children...) }
func Noembed(attrs Attrs, children ...Node) Node    { return Element("noembed", attrs, children...) }
func Noframes(attrs Attrs, children ...Node) Node   { return Element("noframes", attrs, children...) }
func Noscript(attrs Attrs, children ...Node) Node   { return Element("noscript", attrs, children...) }
func Object(attrs Attrs, children ...Node) Node     { return Element("object", attrs, children...) }
func Ol(attrs Attrs, children ...Node) Node         { return Element("ol", attrs, children...) }
func Optgroup(attrs Attrs, children ...Node) Node   { return Element("optgroup", attrs, children...) }
func Option(attrs Attrs, children ...Node) Node     { return Element("option", attrs, children...) }
func Output(attrs Attrs, children ...Node) Node     { return Element("output", attrs, children...) }
func P(attrs Attrs, children ...Node) Node          { return Element("p", attrs, children...) }
func Param(attrs Attrs, children ...Node) Node      { return Element("param", attrs, children...) }
func Picture(attrs Attrs, children ...Node) Node    { return Element("picture", attrs, children...) }
func Pre(attrs Attrs, children ...Node) Node        { return Element("pre", attrs, children...) }
func Progress(attrs Attrs, children ...Node) Node   { return Element("progress", attrs, children...) }
func Q(attrs Attrs, children ...Node) Node          { return Element("q", attrs, children...) }
func Rb(attrs Attrs, children ...Node) Node         { return Element("rb", attrs, children...) }
func Rp(attrs Attrs, children ...Node) Node         { return Element("rp", attrs, children...) }
func Rt(attrs Attrs, children ...Node) Node         { return Element("rt", attrs, children...) }
func Ruby(attrs Attrs, children ...Node) Node       { return Element("ruby", attrs, children...) }
func S(attrs Attrs, children ...Node) Node          { return Element("s", attrs, children...) }
func Samp(attrs Attrs, children ...Node) Node       { return Element("samp", attrs, children...) }
func Script(attrs Attrs, children ...Node) Node     { return Element("script", attrs, children...) }
func Section(attrs Attrs, children ...Node) Node    { return Element("section", attrs, children...) }
func Select(attrs Attrs, children ...Node) Node     { return Element("select", attrs, children...) }
func Small(attrs Attrs, children ...Node) Node      { return Element("small", attrs, children...) }
func Source(attrs Attrs, children ...Node) Node     { return Element("source", attrs, children...) }
func Span(attrs Attrs, children ...Node) Node       { return Element("span", attrs, children...) }
func Strong(attrs Attrs, children ...Node) Node     { return Element("strong", attrs, children...) }
func Style(attrs Attrs, children ...Node) Node      { return Element("style", attrs, children...) }
func Sub(attrs Attrs, children ...Node) Node        { return Element("sub", attrs, children...) }
func Summary(attrs Attrs, children ...Node) Node    { return Element("summary", attrs, children...) }
func Sup(attrs Attrs, children ...Node) Node        { return Element("sup", attrs, children...) }
func Table(attrs Attrs, children ...Node) Node      { return Element("table", attrs, children...) }
func Tbody(attrs Attrs, children ...Node) Node      { return Element("tbody", attrs, children...) }
func Td(attrs Attrs, children ...Node) Node         { return Element("td", attrs, children...) }
func Template(attrs Attrs, children ...Node) Node   { return Element("template", attrs, children...) }
func Textarea(attrs Attrs, children ...Node) Node   { return Element("textarea", attrs, children...) }
func Tfoot(attrs Attrs, children ...Node) Node      { return Element("tfoot", attrs, children...) }
func Th(attrs Attrs, children ...Node) Node         { return Element("th", attrs, children...) }
func Thead(attrs Attrs, children ...Node) Node      { return Element("thead", attrs, children...) }
func Time(attrs Attrs, children ...Node) Node       { return Element("time", attrs, children...) }
func Title(attrs Attrs, children ...Node) Node      { return Element("title", attrs, children...) }
func Tr(attrs Attrs, children ...Node) Node         { return Element("tr", attrs, children...) }
func Track(attrs Attrs, children ...Node) Node      { return Element("track", attrs, children...) }
func U(attrs Attrs, children ...Node) Node          { return Element("u", attrs, children...) }
func Ul(attrs Attrs, children ...Node) Node         { return Element("ul", attrs, children...) }
func Var(attrs Attrs, children ...Node) Node        { return Element("var", attrs, children...) }
func Video(attrs Attrs, children ...Node) Node      { return Element("video", attrs, children...) }
func Wbr(attrs Attrs, children ...Node) Node        { return Element("wbr", attrs, children...) }
