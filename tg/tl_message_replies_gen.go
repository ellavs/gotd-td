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

// MessageReplies represents TL type `messageReplies#4128faac`.
// Info about the comment section of a channel post, or a simple message thread
//
// See https://core.telegram.org/constructor/messageReplies for reference.
type MessageReplies struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether this constructor contains information about the comment section of a channel post, or a simple message thread
	Comments bool
	// Contains the total number of replies in this thread or comment section.
	Replies int
	// PTS of the message that started this thread.
	RepliesPts int
	// For channel post comments, contains information about the last few comment posters for a specific thread, to show a small list of commenter profile pictures in client previews.
	//
	// Use SetRecentRepliers and GetRecentRepliers helpers.
	RecentRepliers []PeerClass
	// For channel post comments, contains the ID of the associated discussion supergroup
	//
	// Use SetChannelID and GetChannelID helpers.
	ChannelID int
	// ID of the latest message in this thread or comment section.
	//
	// Use SetMaxID and GetMaxID helpers.
	MaxID int
	// Contains the ID of the latest read message in this thread or comment section.
	//
	// Use SetReadMaxID and GetReadMaxID helpers.
	ReadMaxID int
}

// MessageRepliesTypeID is TL type id of MessageReplies.
const MessageRepliesTypeID = 0x4128faac

// Encode implements bin.Encoder.
func (m *MessageReplies) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplies#4128faac as nil")
	}
	b.PutID(MessageRepliesTypeID)
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messageReplies#4128faac: field flags: %w", err)
	}
	b.PutInt(m.Replies)
	b.PutInt(m.RepliesPts)
	if m.Flags.Has(1) {
		b.PutVectorHeader(len(m.RecentRepliers))
		for idx, v := range m.RecentRepliers {
			if v == nil {
				return fmt.Errorf("unable to encode messageReplies#4128faac: field recent_repliers element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messageReplies#4128faac: field recent_repliers element with index %d: %w", idx, err)
			}
		}
	}
	if m.Flags.Has(0) {
		b.PutInt(m.ChannelID)
	}
	if m.Flags.Has(2) {
		b.PutInt(m.MaxID)
	}
	if m.Flags.Has(3) {
		b.PutInt(m.ReadMaxID)
	}
	return nil
}

// SetComments sets value of Comments conditional field.
func (m *MessageReplies) SetComments(value bool) {
	if value {
		m.Flags.Set(0)
	} else {
		m.Flags.Unset(0)
	}
}

// SetRecentRepliers sets value of RecentRepliers conditional field.
func (m *MessageReplies) SetRecentRepliers(value []PeerClass) {
	m.Flags.Set(1)
	m.RecentRepliers = value
}

// GetRecentRepliers returns value of RecentRepliers conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetRecentRepliers() (value []PeerClass, ok bool) {
	if !m.Flags.Has(1) {
		return value, false
	}
	return m.RecentRepliers, true
}

// SetChannelID sets value of ChannelID conditional field.
func (m *MessageReplies) SetChannelID(value int) {
	m.Flags.Set(0)
	m.ChannelID = value
}

// GetChannelID returns value of ChannelID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetChannelID() (value int, ok bool) {
	if !m.Flags.Has(0) {
		return value, false
	}
	return m.ChannelID, true
}

// SetMaxID sets value of MaxID conditional field.
func (m *MessageReplies) SetMaxID(value int) {
	m.Flags.Set(2)
	m.MaxID = value
}

// GetMaxID returns value of MaxID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetMaxID() (value int, ok bool) {
	if !m.Flags.Has(2) {
		return value, false
	}
	return m.MaxID, true
}

// SetReadMaxID sets value of ReadMaxID conditional field.
func (m *MessageReplies) SetReadMaxID(value int) {
	m.Flags.Set(3)
	m.ReadMaxID = value
}

// GetReadMaxID returns value of ReadMaxID conditional field and
// boolean which is true if field was set.
func (m *MessageReplies) GetReadMaxID() (value int, ok bool) {
	if !m.Flags.Has(3) {
		return value, false
	}
	return m.ReadMaxID, true
}

// Decode implements bin.Decoder.
func (m *MessageReplies) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplies#4128faac to nil")
	}
	if err := b.ConsumeID(MessageRepliesTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplies#4128faac: %w", err)
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field flags: %w", err)
		}
	}
	m.Comments = m.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field replies: %w", err)
		}
		m.Replies = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field replies_pts: %w", err)
		}
		m.RepliesPts = value
	}
	if m.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field recent_repliers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplies#4128faac: field recent_repliers: %w", err)
			}
			m.RecentRepliers = append(m.RecentRepliers, value)
		}
	}
	if m.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field channel_id: %w", err)
		}
		m.ChannelID = value
	}
	if m.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field max_id: %w", err)
		}
		m.MaxID = value
	}
	if m.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplies#4128faac: field read_max_id: %w", err)
		}
		m.ReadMaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessageReplies.
var (
	_ bin.Encoder = &MessageReplies{}
	_ bin.Decoder = &MessageReplies{}
)
