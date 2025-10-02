package html

import (
	"maps"

	"github.com/GopherGhaznix/Wave/css"
)

// -----------------------
// Types
// -----------------------
type Attrs map[string]string

// Attributes returns a new Attrs map that combines all key-value pairs from the provided maps.
func Attributes(mapsList ...Attrs) Attrs {
	merged := Attrs{}
	for _, m := range mapsList {
		maps.Copy(merged, m)
	}
	return merged
}

// -----------------------
// Attribute Wrappers
// -----------------------

// Common form attributes
func AttrAccept(value string) Attrs       { return Attrs{"accept": value} }
func AttrHref(value string) Attrs         { return Attrs{"href": value} }
func AttrAutocomplete(value string) Attrs { return Attrs{"autocomplete": value} }
func AttrCapture(value string) Attrs      { return Attrs{"capture": value} }
func AttrDisabled() Attrs                 { return Attrs{"disabled": "disabled"} }
func AttrDirname(value string) Attrs      { return Attrs{"dirname": value} }
func AttrFor(value string) Attrs          { return Attrs{"for": value} }
func AttrForm(value string) Attrs         { return Attrs{"form": value} }
func AttrMax(value string) Attrs          { return Attrs{"max": value} }
func AttrMaxLength(value string) Attrs    { return Attrs{"maxlength": value} }
func AttrMin(value string) Attrs          { return Attrs{"min": value} }
func AttrMinLength(value string) Attrs    { return Attrs{"minlength": value} }
func AttrMultiple() Attrs                 { return Attrs{"multiple": "multiple"} }
func AttrPattern(value string) Attrs      { return Attrs{"pattern": value} }
func AttrPlaceholder(value string) Attrs  { return Attrs{"placeholder": value} }
func AttrReadonly() Attrs                 { return Attrs{"readonly": "readonly"} }
func AttrRel(value string) Attrs          { return Attrs{"rel": value} }
func AttrRequired() Attrs                 { return Attrs{"required": "required"} }
func AttrSize(value string) Attrs         { return Attrs{"size": value} }
func AttrStep(value string) Attrs         { return Attrs{"step": value} }

// Input type attributes
func AttrTypeText() Attrs          { return Attrs{"type": "text"} }
func AttrTypeEmail() Attrs         { return Attrs{"type": "email"} }
func AttrTypePassword() Attrs      { return Attrs{"type": "password"} }
func AttrTypeCheckbox() Attrs      { return Attrs{"type": "checkbox"} }
func AttrTypeRadio() Attrs         { return Attrs{"type": "radio"} }
func AttrTypeFile() Attrs          { return Attrs{"type": "file"} }
func AttrTypeNumber() Attrs        { return Attrs{"type": "number"} }
func AttrTypeRange() Attrs         { return Attrs{"type": "range"} }
func AttrTypeDate() Attrs          { return Attrs{"type": "date"} }
func AttrTypeTime() Attrs          { return Attrs{"type": "time"} }
func AttrTypeDatetimeLocal() Attrs { return Attrs{"type": "datetime-local"} }
func AttrTypeMonth() Attrs         { return Attrs{"type": "month"} }
func AttrTypeWeek() Attrs          { return Attrs{"type": "week"} }
func AttrTypeColor() Attrs         { return Attrs{"type": "color"} }
func AttrTypeHidden() Attrs        { return Attrs{"type": "hidden"} }
func AttrTypeSearch() Attrs        { return Attrs{"type": "search"} }
func AttrTypeTel() Attrs           { return Attrs{"type": "tel"} }
func AttrTypeUrl() Attrs           { return Attrs{"type": "url"} }
func AttrTypeButton() Attrs        { return Attrs{"type": "button"} }
func AttrTypeSubmit() Attrs        { return Attrs{"type": "submit"} }
func AttrTypeReset() Attrs         { return Attrs{"type": "reset"} }
func AttrTypeImage() Attrs         { return Attrs{"type": "image"} }

// Media/timing
func AttrElementTiming(value string) Attrs { return Attrs{"elementtiming": value} }
func AttrCrossOrigin(value string) Attrs   { return Attrs{"crossorigin": value} }

// Global attributes
func AttrAccessKey(value string) Attrs             { return Attrs{"accesskey": value} }
func AttrAnchor(value string) Attrs                { return Attrs{"anchor": value} }
func AttrAutoCapitalize(value string) Attrs        { return Attrs{"autocapitalize": value} }
func AttrAutoCorrect(value string) Attrs           { return Attrs{"autocorrect": value} }
func AttrAutoFocus() Attrs                         { return Attrs{"autofocus": "autofocus"} }
func AttrClass(value string) Attrs                 { return Attrs{"class": value} }
func AttrContentEditable(value string) Attrs       { return Attrs{"contenteditable": value} }
func AttrData(name, value string) Attrs            { return Attrs{"data-" + name: value} }
func AttrDir(value string) Attrs                   { return Attrs{"dir": value} }
func AttrDraggable(value string) Attrs             { return Attrs{"draggable": value} }
func AttrEnterKeyHint(value string) Attrs          { return Attrs{"enterkeyhint": value} }
func AttrExportParts(value string) Attrs           { return Attrs{"exportparts": value} }
func AttrHidden() Attrs                            { return Attrs{"hidden": "hidden"} }
func AttrID(value string) Attrs                    { return Attrs{"id": value} }
func AttrInert() Attrs                             { return Attrs{"inert": "inert"} }
func AttrInputMode(value string) Attrs             { return Attrs{"inputmode": value} }
func AttrIs(value string) Attrs                    { return Attrs{"is": value} }
func AttrItemID(value string) Attrs                { return Attrs{"itemid": value} }
func AttrItemProp(value string) Attrs              { return Attrs{"itemprop": value} }
func AttrItemRef(value string) Attrs               { return Attrs{"itemref": value} }
func AttrItemScope() Attrs                         { return Attrs{"itemscope": "itemscope"} }
func AttrItemType(value string) Attrs              { return Attrs{"itemtype": value} }
func AttrLang(value string) Attrs                  { return Attrs{"lang": value} }
func AttrNonce(value string) Attrs                 { return Attrs{"nonce": value} }
func AttrPart(value string) Attrs                  { return Attrs{"part": value} }
func AttrPopover(value string) Attrs               { return Attrs{"popover": value} }
func AttrSlot(value string) Attrs                  { return Attrs{"slot": value} }
func AttrSpellCheck(value string) Attrs            { return Attrs{"spellcheck": value} }
func AttrTabIndex(value string) Attrs              { return Attrs{"tabindex": value} }
func AttrTitle(value string) Attrs                 { return Attrs{"title": value} }
func AttrTranslate(value string) Attrs             { return Attrs{"translate": value} }
func AttrVirtualKeyboardPolicy(value string) Attrs { return Attrs{"virtualkeyboardpolicy": value} }
func AttrWritingSuggestions(value string) Attrs    { return Attrs{"writingsuggestions": value} }

// AttrStyle converts a style map into an inline "style" attribute.
func AttrStyle(styles css.Style) Attrs { return Attrs{"style": styles.Inline()} }

// <script> specific
func AttrImportMap(value string) Attrs        { return Attrs{"importmap": value} }
func AttrSpeculationRules(value string) Attrs { return Attrs{"speculationrules": value} }

// <meta> specific
func AttrMetaName(value string) Attrs    { return Attrs{"name": value} }
func AttrColorScheme(value string) Attrs { return Attrs{"color-scheme": value} }
func AttrReferrer(value string) Attrs    { return Attrs{"referrer": value} }
func AttrRobots(value string) Attrs      { return Attrs{"robots": value} }
func AttrThemeColor(value string) Attrs  { return Attrs{"theme-color": value} }
func AttrViewport(value string) Attrs    { return Attrs{"viewport": value} }
func AttrHttpEquiv(value string) Attrs   { return Attrs{"http-equiv": value} }
