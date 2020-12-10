// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// HelpPromoDataEmpty represents TL type `help.promoDataEmpty#98f6ac75`.
// No PSA/MTProxy info is available
//
// See https://core.telegram.org/constructor/help.promoDataEmpty for reference.
type HelpPromoDataEmpty struct {
	// Re-fetch PSA/MTProxy info after the specified number of seconds
	Expires int
}

// HelpPromoDataEmptyTypeID is TL type id of HelpPromoDataEmpty.
const HelpPromoDataEmptyTypeID = 0x98f6ac75

// Encode implements bin.Encoder.
func (p *HelpPromoDataEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.promoDataEmpty#98f6ac75 as nil")
	}
	b.PutID(HelpPromoDataEmptyTypeID)
	b.PutInt(p.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (p *HelpPromoDataEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.promoDataEmpty#98f6ac75 to nil")
	}
	if err := b.ConsumeID(HelpPromoDataEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.promoDataEmpty#98f6ac75: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoDataEmpty#98f6ac75: field expires: %w", err)
		}
		p.Expires = value
	}
	return nil
}

// construct implements constructor of HelpPromoDataClass.
func (p HelpPromoDataEmpty) construct() HelpPromoDataClass { return &p }

// Ensuring interfaces in compile-time for HelpPromoDataEmpty.
var (
	_ bin.Encoder = &HelpPromoDataEmpty{}
	_ bin.Decoder = &HelpPromoDataEmpty{}

	_ HelpPromoDataClass = &HelpPromoDataEmpty{}
)

// HelpPromoData represents TL type `help.promoData#8c39793f`.
// MTProxy/Public Service Announcement information
//
// See https://core.telegram.org/constructor/help.promoData for reference.
type HelpPromoData struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// MTProxy-related channel
	Proxy bool
	// Expiry of PSA/MTProxy info
	Expires int
	// MTProxy/PSA peer
	Peer PeerClass
	// Chat info
	Chats []ChatClass
	// User info
	Users []UserClass
	// PSA type
	//
	// Use SetPsaType and GetPsaType helpers.
	PsaType string
	// PSA message
	//
	// Use SetPsaMessage and GetPsaMessage helpers.
	PsaMessage string
}

// HelpPromoDataTypeID is TL type id of HelpPromoData.
const HelpPromoDataTypeID = 0x8c39793f

// Encode implements bin.Encoder.
func (p *HelpPromoData) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode help.promoData#8c39793f as nil")
	}
	b.PutID(HelpPromoDataTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field flags: %w", err)
	}
	b.PutInt(p.Expires)
	if p.Peer == nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field peer is nil")
	}
	if err := p.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.promoData#8c39793f: field peer: %w", err)
	}
	b.PutVectorHeader(len(p.Chats))
	for idx, v := range p.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode help.promoData#8c39793f: field users element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(1) {
		b.PutString(p.PsaType)
	}
	if p.Flags.Has(2) {
		b.PutString(p.PsaMessage)
	}
	return nil
}

// SetProxy sets value of Proxy conditional field.
func (p *HelpPromoData) SetProxy(value bool) {
	if value {
		p.Flags.Set(0)
	} else {
		p.Flags.Unset(0)
	}
}

// SetPsaType sets value of PsaType conditional field.
func (p *HelpPromoData) SetPsaType(value string) {
	p.Flags.Set(1)
	p.PsaType = value
}

// GetPsaType returns value of PsaType conditional field and
// boolean which is true if field was set.
func (p *HelpPromoData) GetPsaType() (value string, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.PsaType, true
}

// SetPsaMessage sets value of PsaMessage conditional field.
func (p *HelpPromoData) SetPsaMessage(value string) {
	p.Flags.Set(2)
	p.PsaMessage = value
}

// GetPsaMessage returns value of PsaMessage conditional field and
// boolean which is true if field was set.
func (p *HelpPromoData) GetPsaMessage() (value string, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.PsaMessage, true
}

// Decode implements bin.Decoder.
func (p *HelpPromoData) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode help.promoData#8c39793f to nil")
	}
	if err := b.ConsumeID(HelpPromoDataTypeID); err != nil {
		return fmt.Errorf("unable to decode help.promoData#8c39793f: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field flags: %w", err)
		}
	}
	p.Proxy = p.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field expires: %w", err)
		}
		p.Expires = value
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field peer: %w", err)
		}
		p.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.promoData#8c39793f: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.promoData#8c39793f: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field psa_type: %w", err)
		}
		p.PsaType = value
	}
	if p.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.promoData#8c39793f: field psa_message: %w", err)
		}
		p.PsaMessage = value
	}
	return nil
}

// construct implements constructor of HelpPromoDataClass.
func (p HelpPromoData) construct() HelpPromoDataClass { return &p }

// Ensuring interfaces in compile-time for HelpPromoData.
var (
	_ bin.Encoder = &HelpPromoData{}
	_ bin.Decoder = &HelpPromoData{}

	_ HelpPromoDataClass = &HelpPromoData{}
)

// HelpPromoDataClass represents help.PromoData generic type.
//
// See https://core.telegram.org/type/help.PromoData for reference.
//
// Example:
//  g, err := DecodeHelpPromoData(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpPromoDataEmpty: // help.promoDataEmpty#98f6ac75
//  case *HelpPromoData: // help.promoData#8c39793f
//  default: panic(v)
//  }
type HelpPromoDataClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpPromoDataClass
}

// DecodeHelpPromoData implements binary de-serialization for HelpPromoDataClass.
func DecodeHelpPromoData(buf *bin.Buffer) (HelpPromoDataClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpPromoDataEmptyTypeID:
		// Decoding help.promoDataEmpty#98f6ac75.
		v := HelpPromoDataEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", err)
		}
		return &v, nil
	case HelpPromoDataTypeID:
		// Decoding help.promoData#8c39793f.
		v := HelpPromoData{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpPromoDataClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpPromoData boxes the HelpPromoDataClass providing a helper.
type HelpPromoDataBox struct {
	PromoData HelpPromoDataClass
}

// Decode implements bin.Decoder for HelpPromoDataBox.
func (b *HelpPromoDataBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpPromoDataBox to nil")
	}
	v, err := DecodeHelpPromoData(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PromoData = v
	return nil
}

// Encode implements bin.Encode for HelpPromoDataBox.
func (b *HelpPromoDataBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PromoData == nil {
		return fmt.Errorf("unable to encode HelpPromoDataClass as nil")
	}
	return b.PromoData.Encode(buf)
}
