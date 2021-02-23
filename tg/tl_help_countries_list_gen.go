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

// HelpCountriesListNotModified represents TL type `help.countriesListNotModified#93cc1f32`.
// The country list has not changed
//
// See https://core.telegram.org/constructor/help.countriesListNotModified for reference.
type HelpCountriesListNotModified struct {
}

// HelpCountriesListNotModifiedTypeID is TL type id of HelpCountriesListNotModified.
const HelpCountriesListNotModifiedTypeID = 0x93cc1f32

func (c *HelpCountriesListNotModified) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *HelpCountriesListNotModified) String() string {
	if c == nil {
		return "HelpCountriesListNotModified(nil)"
	}
	type Alias HelpCountriesListNotModified
	return fmt.Sprintf("HelpCountriesListNotModified%+v", Alias(*c))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *HelpCountriesListNotModified) TypeID() uint32 {
	return HelpCountriesListNotModifiedTypeID
}

// SchemaName returns MTProto type name.
func (c *HelpCountriesListNotModified) SchemaName() string {
	return "help.countriesListNotModified"
}

// Encode implements bin.Encoder.
func (c *HelpCountriesListNotModified) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.countriesListNotModified#93cc1f32 as nil")
	}
	b.PutID(HelpCountriesListNotModifiedTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *HelpCountriesListNotModified) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.countriesListNotModified#93cc1f32 to nil")
	}
	if err := b.ConsumeID(HelpCountriesListNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode help.countriesListNotModified#93cc1f32: %w", err)
	}
	return nil
}

// construct implements constructor of HelpCountriesListClass.
func (c HelpCountriesListNotModified) construct() HelpCountriesListClass { return &c }

// Ensuring interfaces in compile-time for HelpCountriesListNotModified.
var (
	_ bin.Encoder = &HelpCountriesListNotModified{}
	_ bin.Decoder = &HelpCountriesListNotModified{}

	_ HelpCountriesListClass = &HelpCountriesListNotModified{}
)

// HelpCountriesList represents TL type `help.countriesList#87d0759e`.
// Name, ISO code, localized name and phone codes/patterns of all available countries
//
// See https://core.telegram.org/constructor/help.countriesList for reference.
type HelpCountriesList struct {
	// Name, ISO code, localized name and phone codes/patterns of all available countries
	Countries []HelpCountry `schemaname:"countries"`
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int `schemaname:"hash"`
}

// HelpCountriesListTypeID is TL type id of HelpCountriesList.
const HelpCountriesListTypeID = 0x87d0759e

func (c *HelpCountriesList) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Countries == nil) {
		return false
	}
	if !(c.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *HelpCountriesList) String() string {
	if c == nil {
		return "HelpCountriesList(nil)"
	}
	type Alias HelpCountriesList
	return fmt.Sprintf("HelpCountriesList%+v", Alias(*c))
}

// FillFrom fills HelpCountriesList from given interface.
func (c *HelpCountriesList) FillFrom(from interface {
	GetCountries() (value []HelpCountry)
	GetHash() (value int)
}) {
	c.Countries = from.GetCountries()
	c.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *HelpCountriesList) TypeID() uint32 {
	return HelpCountriesListTypeID
}

// SchemaName returns MTProto type name.
func (c *HelpCountriesList) SchemaName() string {
	return "help.countriesList"
}

// Encode implements bin.Encoder.
func (c *HelpCountriesList) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode help.countriesList#87d0759e as nil")
	}
	b.PutID(HelpCountriesListTypeID)
	b.PutVectorHeader(len(c.Countries))
	for idx, v := range c.Countries {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.countriesList#87d0759e: field countries element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.Hash)
	return nil
}

// GetCountries returns value of Countries field.
func (c *HelpCountriesList) GetCountries() (value []HelpCountry) {
	return c.Countries
}

// GetHash returns value of Hash field.
func (c *HelpCountriesList) GetHash() (value int) {
	return c.Hash
}

// Decode implements bin.Decoder.
func (c *HelpCountriesList) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode help.countriesList#87d0759e to nil")
	}
	if err := b.ConsumeID(HelpCountriesListTypeID); err != nil {
		return fmt.Errorf("unable to decode help.countriesList#87d0759e: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.countriesList#87d0759e: field countries: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value HelpCountry
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode help.countriesList#87d0759e: field countries: %w", err)
			}
			c.Countries = append(c.Countries, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.countriesList#87d0759e: field hash: %w", err)
		}
		c.Hash = value
	}
	return nil
}

// construct implements constructor of HelpCountriesListClass.
func (c HelpCountriesList) construct() HelpCountriesListClass { return &c }

// Ensuring interfaces in compile-time for HelpCountriesList.
var (
	_ bin.Encoder = &HelpCountriesList{}
	_ bin.Decoder = &HelpCountriesList{}

	_ HelpCountriesListClass = &HelpCountriesList{}
)

// HelpCountriesListClass represents help.CountriesList generic type.
//
// See https://core.telegram.org/type/help.CountriesList for reference.
//
// Example:
//  g, err := tg.DecodeHelpCountriesList(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.HelpCountriesListNotModified: // help.countriesListNotModified#93cc1f32
//  case *tg.HelpCountriesList: // help.countriesList#87d0759e
//  default: panic(v)
//  }
type HelpCountriesListClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpCountriesListClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map HelpCountriesListClass to HelpCountriesList.
	AsModified() (*HelpCountriesList, bool)
}

// AsModified tries to map HelpCountriesListNotModified to HelpCountriesList.
func (c *HelpCountriesListNotModified) AsModified() (*HelpCountriesList, bool) {
	return nil, false
}

// AsModified tries to map HelpCountriesList to HelpCountriesList.
func (c *HelpCountriesList) AsModified() (*HelpCountriesList, bool) {
	return c, true
}

// DecodeHelpCountriesList implements binary de-serialization for HelpCountriesListClass.
func DecodeHelpCountriesList(buf *bin.Buffer) (HelpCountriesListClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpCountriesListNotModifiedTypeID:
		// Decoding help.countriesListNotModified#93cc1f32.
		v := HelpCountriesListNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", err)
		}
		return &v, nil
	case HelpCountriesListTypeID:
		// Decoding help.countriesList#87d0759e.
		v := HelpCountriesList{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpCountriesListClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpCountriesList boxes the HelpCountriesListClass providing a helper.
type HelpCountriesListBox struct {
	CountriesList HelpCountriesListClass
}

// Decode implements bin.Decoder for HelpCountriesListBox.
func (b *HelpCountriesListBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpCountriesListBox to nil")
	}
	v, err := DecodeHelpCountriesList(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CountriesList = v
	return nil
}

// Encode implements bin.Encode for HelpCountriesListBox.
func (b *HelpCountriesListBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CountriesList == nil {
		return fmt.Errorf("unable to encode HelpCountriesListClass as nil")
	}
	return b.CountriesList.Encode(buf)
}

// HelpCountriesListClassSlice is adapter for slice of HelpCountriesListClass.
type HelpCountriesListClassSlice []HelpCountriesListClass

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s HelpCountriesListClassSlice) AppendOnlyModified(to []*HelpCountriesList) []*HelpCountriesList {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s HelpCountriesListClassSlice) AsModified() (to []*HelpCountriesList) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s HelpCountriesListClassSlice) FirstAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s HelpCountriesListClassSlice) LastAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *HelpCountriesListClassSlice) PopFirstAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *HelpCountriesListClassSlice) PopAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// First returns first element of slice (if exists).
func (s HelpCountriesListClassSlice) First() (v HelpCountriesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpCountriesListClassSlice) Last() (v HelpCountriesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpCountriesListClassSlice) PopFirst() (v HelpCountriesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpCountriesListClassSlice) Pop() (v HelpCountriesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
