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

// ChannelAdminLogEvent represents TL type `channelAdminLogEvent#3b5a3e40`.
// Admin log event
//
// See https://core.telegram.org/constructor/channelAdminLogEvent for reference.
type ChannelAdminLogEvent struct {
	// Event ID
	ID int64 `schemaname:"id"`
	// Date
	Date int `schemaname:"date"`
	// User ID
	UserID int `schemaname:"user_id"`
	// Action
	Action ChannelAdminLogEventActionClass `schemaname:"action"`
}

// ChannelAdminLogEventTypeID is TL type id of ChannelAdminLogEvent.
const ChannelAdminLogEventTypeID = 0x3b5a3e40

func (c *ChannelAdminLogEvent) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Action == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelAdminLogEvent) String() string {
	if c == nil {
		return "ChannelAdminLogEvent(nil)"
	}
	type Alias ChannelAdminLogEvent
	return fmt.Sprintf("ChannelAdminLogEvent%+v", Alias(*c))
}

// FillFrom fills ChannelAdminLogEvent from given interface.
func (c *ChannelAdminLogEvent) FillFrom(from interface {
	GetID() (value int64)
	GetDate() (value int)
	GetUserID() (value int)
	GetAction() (value ChannelAdminLogEventActionClass)
}) {
	c.ID = from.GetID()
	c.Date = from.GetDate()
	c.UserID = from.GetUserID()
	c.Action = from.GetAction()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelAdminLogEvent) TypeID() uint32 {
	return ChannelAdminLogEventTypeID
}

// SchemaName returns MTProto type name.
func (c *ChannelAdminLogEvent) SchemaName() string {
	return "channelAdminLogEvent"
}

// Encode implements bin.Encoder.
func (c *ChannelAdminLogEvent) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelAdminLogEvent#3b5a3e40 as nil")
	}
	b.PutID(ChannelAdminLogEventTypeID)
	b.PutLong(c.ID)
	b.PutInt(c.Date)
	b.PutInt(c.UserID)
	if c.Action == nil {
		return fmt.Errorf("unable to encode channelAdminLogEvent#3b5a3e40: field action is nil")
	}
	if err := c.Action.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelAdminLogEvent#3b5a3e40: field action: %w", err)
	}
	return nil
}

// GetID returns value of ID field.
func (c *ChannelAdminLogEvent) GetID() (value int64) {
	return c.ID
}

// GetDate returns value of Date field.
func (c *ChannelAdminLogEvent) GetDate() (value int) {
	return c.Date
}

// GetUserID returns value of UserID field.
func (c *ChannelAdminLogEvent) GetUserID() (value int) {
	return c.UserID
}

// GetAction returns value of Action field.
func (c *ChannelAdminLogEvent) GetAction() (value ChannelAdminLogEventActionClass) {
	return c.Action
}

// Decode implements bin.Decoder.
func (c *ChannelAdminLogEvent) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelAdminLogEvent#3b5a3e40 to nil")
	}
	if err := b.ConsumeID(ChannelAdminLogEventTypeID); err != nil {
		return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field date: %w", err)
		}
		c.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := DecodeChannelAdminLogEventAction(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelAdminLogEvent#3b5a3e40: field action: %w", err)
		}
		c.Action = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelAdminLogEvent.
var (
	_ bin.Encoder = &ChannelAdminLogEvent{}
	_ bin.Decoder = &ChannelAdminLogEvent{}
)
