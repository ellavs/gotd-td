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

// ReplyKeyboardHide represents TL type `replyKeyboardHide#a03e5b85`.
// Hide sent bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardHide for reference.
type ReplyKeyboardHide struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Use this flag if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet
	Selective bool `schemaname:"selective"`
}

// ReplyKeyboardHideTypeID is TL type id of ReplyKeyboardHide.
const ReplyKeyboardHideTypeID = 0xa03e5b85

func (r *ReplyKeyboardHide) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardHide) String() string {
	if r == nil {
		return "ReplyKeyboardHide(nil)"
	}
	type Alias ReplyKeyboardHide
	return fmt.Sprintf("ReplyKeyboardHide%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardHide from given interface.
func (r *ReplyKeyboardHide) FillFrom(from interface {
	GetSelective() (value bool)
}) {
	r.Selective = from.GetSelective()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReplyKeyboardHide) TypeID() uint32 {
	return ReplyKeyboardHideTypeID
}

// SchemaName returns MTProto type name.
func (r *ReplyKeyboardHide) SchemaName() string {
	return "replyKeyboardHide"
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardHide) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardHide#a03e5b85 as nil")
	}
	b.PutID(ReplyKeyboardHideTypeID)
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardHide#a03e5b85: field flags: %w", err)
	}
	return nil
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardHide) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardHide) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardHide) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardHide#a03e5b85 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardHideTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardHide#a03e5b85: field flags: %w", err)
		}
	}
	r.Selective = r.Flags.Has(2)
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardHide) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardHide.
var (
	_ bin.Encoder = &ReplyKeyboardHide{}
	_ bin.Decoder = &ReplyKeyboardHide{}

	_ ReplyMarkupClass = &ReplyKeyboardHide{}
)

// ReplyKeyboardForceReply represents TL type `replyKeyboardForceReply#f4108aa0`.
// Force the user to send a reply
//
// See https://core.telegram.org/constructor/replyKeyboardForceReply for reference.
type ReplyKeyboardForceReply struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again.
	SingleUse bool `schemaname:"single_use"`
	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message. Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
	Selective bool `schemaname:"selective"`
}

// ReplyKeyboardForceReplyTypeID is TL type id of ReplyKeyboardForceReply.
const ReplyKeyboardForceReplyTypeID = 0xf4108aa0

func (r *ReplyKeyboardForceReply) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardForceReply) String() string {
	if r == nil {
		return "ReplyKeyboardForceReply(nil)"
	}
	type Alias ReplyKeyboardForceReply
	return fmt.Sprintf("ReplyKeyboardForceReply%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardForceReply from given interface.
func (r *ReplyKeyboardForceReply) FillFrom(from interface {
	GetSingleUse() (value bool)
	GetSelective() (value bool)
}) {
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReplyKeyboardForceReply) TypeID() uint32 {
	return ReplyKeyboardForceReplyTypeID
}

// SchemaName returns MTProto type name.
func (r *ReplyKeyboardForceReply) SchemaName() string {
	return "replyKeyboardForceReply"
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardForceReply) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardForceReply#f4108aa0 as nil")
	}
	b.PutID(ReplyKeyboardForceReplyTypeID)
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardForceReply#f4108aa0: field flags: %w", err)
	}
	return nil
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardForceReply) GetSingleUse() (value bool) {
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardForceReply) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardForceReply) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardForceReply) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardForceReply#f4108aa0 to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardForceReplyTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardForceReply#f4108aa0: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardForceReply#f4108aa0: field flags: %w", err)
		}
	}
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardForceReply) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardForceReply.
var (
	_ bin.Encoder = &ReplyKeyboardForceReply{}
	_ bin.Decoder = &ReplyKeyboardForceReply{}

	_ ReplyMarkupClass = &ReplyKeyboardForceReply{}
)

// ReplyKeyboardMarkup represents TL type `replyKeyboardMarkup#3502758c`.
// Bot keyboard
//
// See https://core.telegram.org/constructor/replyKeyboardMarkup for reference.
type ReplyKeyboardMarkup struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). If not set, the custom keyboard is always of the same height as the app's standard keyboard.
	Resize bool `schemaname:"resize"`
	// Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again.
	SingleUse bool `schemaname:"single_use"`
	// Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
	Selective bool `schemaname:"selective"`
	// Button row
	Rows []KeyboardButtonRow `schemaname:"rows"`
}

// ReplyKeyboardMarkupTypeID is TL type id of ReplyKeyboardMarkup.
const ReplyKeyboardMarkupTypeID = 0x3502758c

func (r *ReplyKeyboardMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Resize == false) {
		return false
	}
	if !(r.SingleUse == false) {
		return false
	}
	if !(r.Selective == false) {
		return false
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyKeyboardMarkup) String() string {
	if r == nil {
		return "ReplyKeyboardMarkup(nil)"
	}
	type Alias ReplyKeyboardMarkup
	return fmt.Sprintf("ReplyKeyboardMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyKeyboardMarkup from given interface.
func (r *ReplyKeyboardMarkup) FillFrom(from interface {
	GetResize() (value bool)
	GetSingleUse() (value bool)
	GetSelective() (value bool)
	GetRows() (value []KeyboardButtonRow)
}) {
	r.Resize = from.GetResize()
	r.SingleUse = from.GetSingleUse()
	r.Selective = from.GetSelective()
	r.Rows = from.GetRows()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReplyKeyboardMarkup) TypeID() uint32 {
	return ReplyKeyboardMarkupTypeID
}

// SchemaName returns MTProto type name.
func (r *ReplyKeyboardMarkup) SchemaName() string {
	return "replyKeyboardMarkup"
}

// Encode implements bin.Encoder.
func (r *ReplyKeyboardMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyKeyboardMarkup#3502758c as nil")
	}
	b.PutID(ReplyKeyboardMarkupTypeID)
	if !(r.Resize == false) {
		r.Flags.Set(0)
	}
	if !(r.SingleUse == false) {
		r.Flags.Set(1)
	}
	if !(r.Selective == false) {
		r.Flags.Set(2)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replyKeyboardMarkup#3502758c: field flags: %w", err)
	}
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyKeyboardMarkup#3502758c: field rows element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetResize sets value of Resize conditional field.
func (r *ReplyKeyboardMarkup) SetResize(value bool) {
	if value {
		r.Flags.Set(0)
		r.Resize = true
	} else {
		r.Flags.Unset(0)
		r.Resize = false
	}
}

// GetResize returns value of Resize conditional field.
func (r *ReplyKeyboardMarkup) GetResize() (value bool) {
	return r.Flags.Has(0)
}

// SetSingleUse sets value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) SetSingleUse(value bool) {
	if value {
		r.Flags.Set(1)
		r.SingleUse = true
	} else {
		r.Flags.Unset(1)
		r.SingleUse = false
	}
}

// GetSingleUse returns value of SingleUse conditional field.
func (r *ReplyKeyboardMarkup) GetSingleUse() (value bool) {
	return r.Flags.Has(1)
}

// SetSelective sets value of Selective conditional field.
func (r *ReplyKeyboardMarkup) SetSelective(value bool) {
	if value {
		r.Flags.Set(2)
		r.Selective = true
	} else {
		r.Flags.Unset(2)
		r.Selective = false
	}
}

// GetSelective returns value of Selective conditional field.
func (r *ReplyKeyboardMarkup) GetSelective() (value bool) {
	return r.Flags.Has(2)
}

// GetRows returns value of Rows field.
func (r *ReplyKeyboardMarkup) GetRows() (value []KeyboardButtonRow) {
	return r.Rows
}

// Decode implements bin.Decoder.
func (r *ReplyKeyboardMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyKeyboardMarkup#3502758c to nil")
	}
	if err := b.ConsumeID(ReplyKeyboardMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field flags: %w", err)
		}
	}
	r.Resize = r.Flags.Has(0)
	r.SingleUse = r.Flags.Has(1)
	r.Selective = r.Flags.Has(2)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field rows: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyKeyboardMarkup#3502758c: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyKeyboardMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyKeyboardMarkup.
var (
	_ bin.Encoder = &ReplyKeyboardMarkup{}
	_ bin.Decoder = &ReplyKeyboardMarkup{}

	_ ReplyMarkupClass = &ReplyKeyboardMarkup{}
)

// ReplyInlineMarkup represents TL type `replyInlineMarkup#48a30254`.
// Bot or inline keyboard
//
// See https://core.telegram.org/constructor/replyInlineMarkup for reference.
type ReplyInlineMarkup struct {
	// Bot or inline keyboard rows
	Rows []KeyboardButtonRow `schemaname:"rows"`
}

// ReplyInlineMarkupTypeID is TL type id of ReplyInlineMarkup.
const ReplyInlineMarkupTypeID = 0x48a30254

func (r *ReplyInlineMarkup) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Rows == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplyInlineMarkup) String() string {
	if r == nil {
		return "ReplyInlineMarkup(nil)"
	}
	type Alias ReplyInlineMarkup
	return fmt.Sprintf("ReplyInlineMarkup%+v", Alias(*r))
}

// FillFrom fills ReplyInlineMarkup from given interface.
func (r *ReplyInlineMarkup) FillFrom(from interface {
	GetRows() (value []KeyboardButtonRow)
}) {
	r.Rows = from.GetRows()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ReplyInlineMarkup) TypeID() uint32 {
	return ReplyInlineMarkupTypeID
}

// SchemaName returns MTProto type name.
func (r *ReplyInlineMarkup) SchemaName() string {
	return "replyInlineMarkup"
}

// Encode implements bin.Encoder.
func (r *ReplyInlineMarkup) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replyInlineMarkup#48a30254 as nil")
	}
	b.PutID(ReplyInlineMarkupTypeID)
	b.PutVectorHeader(len(r.Rows))
	for idx, v := range r.Rows {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode replyInlineMarkup#48a30254: field rows element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetRows returns value of Rows field.
func (r *ReplyInlineMarkup) GetRows() (value []KeyboardButtonRow) {
	return r.Rows
}

// Decode implements bin.Decoder.
func (r *ReplyInlineMarkup) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replyInlineMarkup#48a30254 to nil")
	}
	if err := b.ConsumeID(ReplyInlineMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value KeyboardButtonRow
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode replyInlineMarkup#48a30254: field rows: %w", err)
			}
			r.Rows = append(r.Rows, value)
		}
	}
	return nil
}

// construct implements constructor of ReplyMarkupClass.
func (r ReplyInlineMarkup) construct() ReplyMarkupClass { return &r }

// Ensuring interfaces in compile-time for ReplyInlineMarkup.
var (
	_ bin.Encoder = &ReplyInlineMarkup{}
	_ bin.Decoder = &ReplyInlineMarkup{}

	_ ReplyMarkupClass = &ReplyInlineMarkup{}
)

// ReplyMarkupClass represents ReplyMarkup generic type.
//
// See https://core.telegram.org/type/ReplyMarkup for reference.
//
// Example:
//  g, err := tg.DecodeReplyMarkup(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ReplyKeyboardHide: // replyKeyboardHide#a03e5b85
//  case *tg.ReplyKeyboardForceReply: // replyKeyboardForceReply#f4108aa0
//  case *tg.ReplyKeyboardMarkup: // replyKeyboardMarkup#3502758c
//  case *tg.ReplyInlineMarkup: // replyInlineMarkup#48a30254
//  default: panic(v)
//  }
type ReplyMarkupClass interface {
	bin.Encoder
	bin.Decoder
	construct() ReplyMarkupClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeReplyMarkup implements binary de-serialization for ReplyMarkupClass.
func DecodeReplyMarkup(buf *bin.Buffer) (ReplyMarkupClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ReplyKeyboardHideTypeID:
		// Decoding replyKeyboardHide#a03e5b85.
		v := ReplyKeyboardHide{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardForceReplyTypeID:
		// Decoding replyKeyboardForceReply#f4108aa0.
		v := ReplyKeyboardForceReply{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyKeyboardMarkupTypeID:
		// Decoding replyKeyboardMarkup#3502758c.
		v := ReplyKeyboardMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	case ReplyInlineMarkupTypeID:
		// Decoding replyInlineMarkup#48a30254.
		v := ReplyInlineMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReplyMarkupClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReplyMarkup boxes the ReplyMarkupClass providing a helper.
type ReplyMarkupBox struct {
	ReplyMarkup ReplyMarkupClass
}

// Decode implements bin.Decoder for ReplyMarkupBox.
func (b *ReplyMarkupBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReplyMarkupBox to nil")
	}
	v, err := DecodeReplyMarkup(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReplyMarkup = v
	return nil
}

// Encode implements bin.Encode for ReplyMarkupBox.
func (b *ReplyMarkupBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReplyMarkup == nil {
		return fmt.Errorf("unable to encode ReplyMarkupClass as nil")
	}
	return b.ReplyMarkup.Encode(buf)
}

// ReplyMarkupClassSlice is adapter for slice of ReplyMarkupClass.
type ReplyMarkupClassSlice []ReplyMarkupClass

// First returns first element of slice (if exists).
func (s ReplyMarkupClassSlice) First() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ReplyMarkupClassSlice) Last() (v ReplyMarkupClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ReplyMarkupClassSlice) PopFirst() (v ReplyMarkupClass, ok bool) {
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
func (s *ReplyMarkupClassSlice) Pop() (v ReplyMarkupClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
