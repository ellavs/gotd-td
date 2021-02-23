// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// Theme represents TL type `theme#28f1114`.
// Theme
//
// See https://core.telegram.org/constructor/theme for reference.
type Theme struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether the current user is the creator of this theme
	Creator bool `schemaname:"creator"`
	// Whether this is the default theme
	Default bool `schemaname:"default"`
	// Theme ID
	ID int64 `schemaname:"id"`
	// Theme access hash
	AccessHash int64 `schemaname:"access_hash"`
	// Unique theme ID
	Slug string `schemaname:"slug"`
	// Theme name
	Title string `schemaname:"title"`
	// Theme
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass `schemaname:"document"`
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings ThemeSettings `schemaname:"settings"`
	// Installation count
	InstallsCount int `schemaname:"installs_count"`
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0x28f1114

func (t *Theme) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Creator == false) {
		return false
	}
	if !(t.Default == false) {
		return false
	}
	if !(t.ID == 0) {
		return false
	}
	if !(t.AccessHash == 0) {
		return false
	}
	if !(t.Slug == "") {
		return false
	}
	if !(t.Title == "") {
		return false
	}
	if !(t.Document == nil) {
		return false
	}
	if !(t.Settings.Zero()) {
		return false
	}
	if !(t.InstallsCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Theme) String() string {
	if t == nil {
		return "Theme(nil)"
	}
	type Alias Theme
	return fmt.Sprintf("Theme%+v", Alias(*t))
}

// FillFrom fills Theme from given interface.
func (t *Theme) FillFrom(from interface {
	GetCreator() (value bool)
	GetDefault() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetSlug() (value string)
	GetTitle() (value string)
	GetDocument() (value DocumentClass, ok bool)
	GetSettings() (value ThemeSettings, ok bool)
	GetInstallsCount() (value int)
}) {
	t.Creator = from.GetCreator()
	t.Default = from.GetDefault()
	t.ID = from.GetID()
	t.AccessHash = from.GetAccessHash()
	t.Slug = from.GetSlug()
	t.Title = from.GetTitle()
	if val, ok := from.GetDocument(); ok {
		t.Document = val
	}

	if val, ok := from.GetSettings(); ok {
		t.Settings = val
	}

	t.InstallsCount = from.GetInstallsCount()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *Theme) TypeID() uint32 {
	return ThemeTypeID
}

// SchemaName returns MTProto type name.
func (t *Theme) SchemaName() string {
	return "theme"
}

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutID(ThemeTypeID)
	if !(t.Creator == false) {
		t.Flags.Set(0)
	}
	if !(t.Default == false) {
		t.Flags.Set(1)
	}
	if !(t.Document == nil) {
		t.Flags.Set(2)
	}
	if !(t.Settings.Zero()) {
		t.Flags.Set(3)
	}
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode theme#28f1114: field flags: %w", err)
	}
	b.PutLong(t.ID)
	b.PutLong(t.AccessHash)
	b.PutString(t.Slug)
	b.PutString(t.Title)
	if t.Flags.Has(2) {
		if t.Document == nil {
			return fmt.Errorf("unable to encode theme#28f1114: field document is nil")
		}
		if err := t.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode theme#28f1114: field document: %w", err)
		}
	}
	if t.Flags.Has(3) {
		if err := t.Settings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode theme#28f1114: field settings: %w", err)
		}
	}
	b.PutInt(t.InstallsCount)
	return nil
}

// SetCreator sets value of Creator conditional field.
func (t *Theme) SetCreator(value bool) {
	if value {
		t.Flags.Set(0)
		t.Creator = true
	} else {
		t.Flags.Unset(0)
		t.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (t *Theme) GetCreator() (value bool) {
	return t.Flags.Has(0)
}

// SetDefault sets value of Default conditional field.
func (t *Theme) SetDefault(value bool) {
	if value {
		t.Flags.Set(1)
		t.Default = true
	} else {
		t.Flags.Unset(1)
		t.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (t *Theme) GetDefault() (value bool) {
	return t.Flags.Has(1)
}

// GetID returns value of ID field.
func (t *Theme) GetID() (value int64) {
	return t.ID
}

// GetAccessHash returns value of AccessHash field.
func (t *Theme) GetAccessHash() (value int64) {
	return t.AccessHash
}

// GetSlug returns value of Slug field.
func (t *Theme) GetSlug() (value string) {
	return t.Slug
}

// GetTitle returns value of Title field.
func (t *Theme) GetTitle() (value string) {
	return t.Title
}

// SetDocument sets value of Document conditional field.
func (t *Theme) SetDocument(value DocumentClass) {
	t.Flags.Set(2)
	t.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (t *Theme) GetDocument() (value DocumentClass, ok bool) {
	if !t.Flags.Has(2) {
		return value, false
	}
	return t.Document, true
}

// GetDocumentAsNotEmpty returns mapped value of Document conditional field and
// boolean which is true if field was set.
func (t *Theme) GetDocumentAsNotEmpty() (*Document, bool) {
	if value, ok := t.GetDocument(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// SetSettings sets value of Settings conditional field.
func (t *Theme) SetSettings(value ThemeSettings) {
	t.Flags.Set(3)
	t.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (t *Theme) GetSettings() (value ThemeSettings, ok bool) {
	if !t.Flags.Has(3) {
		return value, false
	}
	return t.Settings, true
}

// GetInstallsCount returns value of InstallsCount field.
func (t *Theme) GetInstallsCount() (value int) {
	return t.InstallsCount
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#28f1114: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field flags: %w", err)
		}
	}
	t.Creator = t.Flags.Has(0)
	t.Default = t.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field id: %w", err)
		}
		t.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field access_hash: %w", err)
		}
		t.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field slug: %w", err)
		}
		t.Slug = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field title: %w", err)
		}
		t.Title = value
	}
	if t.Flags.Has(2) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field document: %w", err)
		}
		t.Document = value
	}
	if t.Flags.Has(3) {
		if err := t.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field settings: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field installs_count: %w", err)
		}
		t.InstallsCount = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder = &Theme{}
	_ bin.Decoder = &Theme{}
)
