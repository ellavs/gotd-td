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

// UpdatesChannelDifferenceEmpty represents TL type `updates.channelDifferenceEmpty#3e11affb`.
// There are no new updates
//
// See https://core.telegram.org/constructor/updates.channelDifferenceEmpty for reference.
type UpdatesChannelDifferenceEmpty struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether there are more updates that must be fetched (always false)
	Final bool `schemaname:"final"`
	// The latest PTS¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int `schemaname:"pts"`
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int `schemaname:"timeout"`
}

// UpdatesChannelDifferenceEmptyTypeID is TL type id of UpdatesChannelDifferenceEmpty.
const UpdatesChannelDifferenceEmptyTypeID = 0x3e11affb

func (c *UpdatesChannelDifferenceEmpty) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceEmpty) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceEmpty(nil)"
	}
	type Alias UpdatesChannelDifferenceEmpty
	return fmt.Sprintf("UpdatesChannelDifferenceEmpty%+v", Alias(*c))
}

// FillFrom fills UpdatesChannelDifferenceEmpty from given interface.
func (c *UpdatesChannelDifferenceEmpty) FillFrom(from interface {
	GetFinal() (value bool)
	GetPts() (value int)
	GetTimeout() (value int, ok bool)
}) {
	c.Final = from.GetFinal()
	c.Pts = from.GetPts()
	if val, ok := from.GetTimeout(); ok {
		c.Timeout = val
	}

}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifferenceEmpty) TypeID() uint32 {
	return UpdatesChannelDifferenceEmptyTypeID
}

// SchemaName returns MTProto type name.
func (c *UpdatesChannelDifferenceEmpty) SchemaName() string {
	return "updates.channelDifferenceEmpty"
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceEmpty#3e11affb as nil")
	}
	b.PutID(UpdatesChannelDifferenceEmptyTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifferenceEmpty) GetPts() (value int) {
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceEmpty) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceEmpty#3e11affb to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field timeout: %w", err)
		}
		c.Timeout = value
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceEmpty) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceEmpty.
var (
	_ bin.Encoder = &UpdatesChannelDifferenceEmpty{}
	_ bin.Decoder = &UpdatesChannelDifferenceEmpty{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceEmpty{}
)

// UpdatesChannelDifferenceTooLong represents TL type `updates.channelDifferenceTooLong#a4bcc6fe`.
// The provided pts + limit < remote pts. Simply, there are too many updates to be fetched (more than limit), the client has to resolve the update gap in one of the following ways:
//
// See https://core.telegram.org/constructor/updates.channelDifferenceTooLong for reference.
type UpdatesChannelDifferenceTooLong struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether there are more updates that must be fetched (always false)
	Final bool `schemaname:"final"`
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int `schemaname:"timeout"`
	// Dialog containing the latest PTS¹ that can be used to reset the channel state
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Dialog DialogClass `schemaname:"dialog"`
	// The latest messages
	Messages []MessageClass `schemaname:"messages"`
	// Chats from messages
	Chats []ChatClass `schemaname:"chats"`
	// Users from messages
	Users []UserClass `schemaname:"users"`
}

// UpdatesChannelDifferenceTooLongTypeID is TL type id of UpdatesChannelDifferenceTooLong.
const UpdatesChannelDifferenceTooLongTypeID = 0xa4bcc6fe

func (c *UpdatesChannelDifferenceTooLong) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.Dialog == nil) {
		return false
	}
	if !(c.Messages == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceTooLong) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceTooLong(nil)"
	}
	type Alias UpdatesChannelDifferenceTooLong
	return fmt.Sprintf("UpdatesChannelDifferenceTooLong%+v", Alias(*c))
}

// FillFrom fills UpdatesChannelDifferenceTooLong from given interface.
func (c *UpdatesChannelDifferenceTooLong) FillFrom(from interface {
	GetFinal() (value bool)
	GetTimeout() (value int, ok bool)
	GetDialog() (value DialogClass)
	GetMessages() (value []MessageClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.Final = from.GetFinal()
	if val, ok := from.GetTimeout(); ok {
		c.Timeout = val
	}

	c.Dialog = from.GetDialog()
	c.Messages = from.GetMessages()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifferenceTooLong) TypeID() uint32 {
	return UpdatesChannelDifferenceTooLongTypeID
}

// SchemaName returns MTProto type name.
func (c *UpdatesChannelDifferenceTooLong) SchemaName() string {
	return "updates.channelDifferenceTooLong"
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceTooLong) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceTooLong#a4bcc6fe as nil")
	}
	b.PutID(UpdatesChannelDifferenceTooLongTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	if c.Dialog == nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog is nil")
	}
	if err := c.Dialog.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
	}
	b.PutVectorHeader(len(c.Messages))
	for idx, v := range c.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceTooLong) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetDialog returns value of Dialog field.
func (c *UpdatesChannelDifferenceTooLong) GetDialog() (value DialogClass) {
	return c.Dialog
}

// GetMessages returns value of Messages field.
func (c *UpdatesChannelDifferenceTooLong) GetMessages() (value []MessageClass) {
	return c.Messages
}

// MapMessages returns field Messages wrapped in MessageClassSlice helper.
func (c *UpdatesChannelDifferenceTooLong) MapMessages() (value MessageClassSlice) {
	return MessageClassSlice(c.Messages)
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifferenceTooLong) GetChats() (value []ChatClass) {
	return c.Chats
}

// MapChats returns field Chats wrapped in ChatClassSlice helper.
func (c *UpdatesChannelDifferenceTooLong) MapChats() (value ChatClassSlice) {
	return ChatClassSlice(c.Chats)
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifferenceTooLong) GetUsers() (value []UserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (c *UpdatesChannelDifferenceTooLong) MapUsers() (value UserClassSlice) {
	return UserClassSlice(c.Users)
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceTooLong) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceTooLong#a4bcc6fe to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTooLongTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		value, err := DecodeDialog(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
		}
		c.Dialog = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
			}
			c.Messages = append(c.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceTooLong) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceTooLong.
var (
	_ bin.Encoder = &UpdatesChannelDifferenceTooLong{}
	_ bin.Decoder = &UpdatesChannelDifferenceTooLong{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceTooLong{}
)

// UpdatesChannelDifference represents TL type `updates.channelDifference#2064674e`.
// The new updates
//
// See https://core.telegram.org/constructor/updates.channelDifference for reference.
type UpdatesChannelDifference struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether there are more updates to be fetched using getDifference, starting from the provided pts
	Final bool `schemaname:"final"`
	// The PTS¹ from which to start getting updates the next time
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int `schemaname:"pts"`
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int `schemaname:"timeout"`
	// New messages
	NewMessages []MessageClass `schemaname:"new_messages"`
	// Other updates
	OtherUpdates []UpdateClass `schemaname:"other_updates"`
	// Chats
	Chats []ChatClass `schemaname:"chats"`
	// Users
	Users []UserClass `schemaname:"users"`
}

// UpdatesChannelDifferenceTypeID is TL type id of UpdatesChannelDifference.
const UpdatesChannelDifferenceTypeID = 0x2064674e

func (c *UpdatesChannelDifference) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.NewMessages == nil) {
		return false
	}
	if !(c.OtherUpdates == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifference) String() string {
	if c == nil {
		return "UpdatesChannelDifference(nil)"
	}
	type Alias UpdatesChannelDifference
	return fmt.Sprintf("UpdatesChannelDifference%+v", Alias(*c))
}

// FillFrom fills UpdatesChannelDifference from given interface.
func (c *UpdatesChannelDifference) FillFrom(from interface {
	GetFinal() (value bool)
	GetPts() (value int)
	GetTimeout() (value int, ok bool)
	GetNewMessages() (value []MessageClass)
	GetOtherUpdates() (value []UpdateClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.Final = from.GetFinal()
	c.Pts = from.GetPts()
	if val, ok := from.GetTimeout(); ok {
		c.Timeout = val
	}

	c.NewMessages = from.GetNewMessages()
	c.OtherUpdates = from.GetOtherUpdates()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifference) TypeID() uint32 {
	return UpdatesChannelDifferenceTypeID
}

// SchemaName returns MTProto type name.
func (c *UpdatesChannelDifference) SchemaName() string {
	return "updates.channelDifference"
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifference) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifference#2064674e as nil")
	}
	b.PutID(UpdatesChannelDifferenceTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	b.PutVectorHeader(len(c.NewMessages))
	for idx, v := range c.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.OtherUpdates))
	for idx, v := range c.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifference) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifference) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifference) GetPts() (value int) {
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifference) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifference) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetNewMessages returns value of NewMessages field.
func (c *UpdatesChannelDifference) GetNewMessages() (value []MessageClass) {
	return c.NewMessages
}

// MapNewMessages returns field NewMessages wrapped in MessageClassSlice helper.
func (c *UpdatesChannelDifference) MapNewMessages() (value MessageClassSlice) {
	return MessageClassSlice(c.NewMessages)
}

// GetOtherUpdates returns value of OtherUpdates field.
func (c *UpdatesChannelDifference) GetOtherUpdates() (value []UpdateClass) {
	return c.OtherUpdates
}

// MapOtherUpdates returns field OtherUpdates wrapped in UpdateClassSlice helper.
func (c *UpdatesChannelDifference) MapOtherUpdates() (value UpdateClassSlice) {
	return UpdateClassSlice(c.OtherUpdates)
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifference) GetChats() (value []ChatClass) {
	return c.Chats
}

// MapChats returns field Chats wrapped in ChatClassSlice helper.
func (c *UpdatesChannelDifference) MapChats() (value ChatClassSlice) {
	return ChatClassSlice(c.Chats)
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifference) GetUsers() (value []UserClass) {
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (c *UpdatesChannelDifference) MapUsers() (value UserClassSlice) {
	return UserClassSlice(c.Users)
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifference) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifference#2064674e to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifference#2064674e: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
			}
			c.NewMessages = append(c.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
			}
			c.OtherUpdates = append(c.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifference) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifference.
var (
	_ bin.Encoder = &UpdatesChannelDifference{}
	_ bin.Decoder = &UpdatesChannelDifference{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifference{}
)

// UpdatesChannelDifferenceClass represents updates.ChannelDifference generic type.
//
// See https://core.telegram.org/type/updates.ChannelDifference for reference.
//
// Example:
//  g, err := tg.DecodeUpdatesChannelDifference(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UpdatesChannelDifferenceEmpty: // updates.channelDifferenceEmpty#3e11affb
//  case *tg.UpdatesChannelDifferenceTooLong: // updates.channelDifferenceTooLong#a4bcc6fe
//  case *tg.UpdatesChannelDifference: // updates.channelDifference#2064674e
//  default: panic(v)
//  }
type UpdatesChannelDifferenceClass interface {
	bin.Encoder
	bin.Decoder
	construct() UpdatesChannelDifferenceClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Whether there are more updates that must be fetched (always false)
	GetFinal() (value bool)
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	GetTimeout() (value int, ok bool)
	// AsNotEmpty tries to map UpdatesChannelDifferenceClass to NotEmptyUpdatesChannelDifference.
	AsNotEmpty() (NotEmptyUpdatesChannelDifference, bool)
}

// NotEmptyUpdatesChannelDifference represents NotEmpty subset of UpdatesChannelDifferenceClass.
type NotEmptyUpdatesChannelDifference interface {
	bin.Encoder
	bin.Decoder
	construct() UpdatesChannelDifferenceClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Whether there are more updates that must be fetched (always false)
	GetFinal() (value bool)
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	GetTimeout() (value int, ok bool)
	// Chats from messages
	GetChats() (value []ChatClass)
	// Users from messages
	GetUsers() (value []UserClass)
}

// AsNotEmpty tries to map UpdatesChannelDifferenceEmpty to NotEmptyUpdatesChannelDifference.
func (c *UpdatesChannelDifferenceEmpty) AsNotEmpty() (NotEmptyUpdatesChannelDifference, bool) {
	value, ok := (UpdatesChannelDifferenceClass(c)).(NotEmptyUpdatesChannelDifference)
	return value, ok
}

// AsNotEmpty tries to map UpdatesChannelDifferenceTooLong to NotEmptyUpdatesChannelDifference.
func (c *UpdatesChannelDifferenceTooLong) AsNotEmpty() (NotEmptyUpdatesChannelDifference, bool) {
	value, ok := (UpdatesChannelDifferenceClass(c)).(NotEmptyUpdatesChannelDifference)
	return value, ok
}

// AsNotEmpty tries to map UpdatesChannelDifference to NotEmptyUpdatesChannelDifference.
func (c *UpdatesChannelDifference) AsNotEmpty() (NotEmptyUpdatesChannelDifference, bool) {
	value, ok := (UpdatesChannelDifferenceClass(c)).(NotEmptyUpdatesChannelDifference)
	return value, ok
}

// DecodeUpdatesChannelDifference implements binary de-serialization for UpdatesChannelDifferenceClass.
func DecodeUpdatesChannelDifference(buf *bin.Buffer) (UpdatesChannelDifferenceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UpdatesChannelDifferenceEmptyTypeID:
		// Decoding updates.channelDifferenceEmpty#3e11affb.
		v := UpdatesChannelDifferenceEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTooLongTypeID:
		// Decoding updates.channelDifferenceTooLong#a4bcc6fe.
		v := UpdatesChannelDifferenceTooLong{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTypeID:
		// Decoding updates.channelDifference#2064674e.
		v := UpdatesChannelDifference{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", bin.NewUnexpectedID(id))
	}
}

// UpdatesChannelDifference boxes the UpdatesChannelDifferenceClass providing a helper.
type UpdatesChannelDifferenceBox struct {
	ChannelDifference UpdatesChannelDifferenceClass
}

// Decode implements bin.Decoder for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UpdatesChannelDifferenceBox to nil")
	}
	v, err := DecodeUpdatesChannelDifference(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelDifference = v
	return nil
}

// Encode implements bin.Encode for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelDifference == nil {
		return fmt.Errorf("unable to encode UpdatesChannelDifferenceClass as nil")
	}
	return b.ChannelDifference.Encode(buf)
}

// UpdatesChannelDifferenceClassSlice is adapter for slice of UpdatesChannelDifferenceClass.
type UpdatesChannelDifferenceClassSlice []UpdatesChannelDifferenceClass

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s UpdatesChannelDifferenceClassSlice) AppendOnlyNotEmpty(to []NotEmptyUpdatesChannelDifference) []NotEmptyUpdatesChannelDifference {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s UpdatesChannelDifferenceClassSlice) AsNotEmpty() (to []NotEmptyUpdatesChannelDifference) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s UpdatesChannelDifferenceClassSlice) FirstAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s UpdatesChannelDifferenceClassSlice) LastAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *UpdatesChannelDifferenceClassSlice) PopFirstAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *UpdatesChannelDifferenceClassSlice) PopAsNotEmpty() (v NotEmptyUpdatesChannelDifference, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// First returns first element of slice (if exists).
func (s UpdatesChannelDifferenceClassSlice) First() (v UpdatesChannelDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s UpdatesChannelDifferenceClassSlice) Last() (v UpdatesChannelDifferenceClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *UpdatesChannelDifferenceClassSlice) PopFirst() (v UpdatesChannelDifferenceClass, ok bool) {
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
func (s *UpdatesChannelDifferenceClassSlice) Pop() (v UpdatesChannelDifferenceClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
