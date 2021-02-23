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

// NotifyPeer represents TL type `notifyPeer#9fd40bd8`.
// Notifications generated by a certain user or group.
//
// See https://core.telegram.org/constructor/notifyPeer for reference.
type NotifyPeer struct {
	// user or group
	Peer PeerClass `schemaname:"peer"`
}

// NotifyPeerTypeID is TL type id of NotifyPeer.
const NotifyPeerTypeID = 0x9fd40bd8

func (n *NotifyPeer) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotifyPeer) String() string {
	if n == nil {
		return "NotifyPeer(nil)"
	}
	type Alias NotifyPeer
	return fmt.Sprintf("NotifyPeer%+v", Alias(*n))
}

// FillFrom fills NotifyPeer from given interface.
func (n *NotifyPeer) FillFrom(from interface {
	GetPeer() (value PeerClass)
}) {
	n.Peer = from.GetPeer()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NotifyPeer) TypeID() uint32 {
	return NotifyPeerTypeID
}

// SchemaName returns MTProto type name.
func (n *NotifyPeer) SchemaName() string {
	return "notifyPeer"
}

// Encode implements bin.Encoder.
func (n *NotifyPeer) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notifyPeer#9fd40bd8 as nil")
	}
	b.PutID(NotifyPeerTypeID)
	if n.Peer == nil {
		return fmt.Errorf("unable to encode notifyPeer#9fd40bd8: field peer is nil")
	}
	if err := n.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notifyPeer#9fd40bd8: field peer: %w", err)
	}
	return nil
}

// GetPeer returns value of Peer field.
func (n *NotifyPeer) GetPeer() (value PeerClass) {
	return n.Peer
}

// Decode implements bin.Decoder.
func (n *NotifyPeer) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notifyPeer#9fd40bd8 to nil")
	}
	if err := b.ConsumeID(NotifyPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode notifyPeer#9fd40bd8: %w", err)
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode notifyPeer#9fd40bd8: field peer: %w", err)
		}
		n.Peer = value
	}
	return nil
}

// construct implements constructor of NotifyPeerClass.
func (n NotifyPeer) construct() NotifyPeerClass { return &n }

// Ensuring interfaces in compile-time for NotifyPeer.
var (
	_ bin.Encoder = &NotifyPeer{}
	_ bin.Decoder = &NotifyPeer{}

	_ NotifyPeerClass = &NotifyPeer{}
)

// NotifyUsers represents TL type `notifyUsers#b4c83b4c`.
// Notifications generated by all users.
//
// See https://core.telegram.org/constructor/notifyUsers for reference.
type NotifyUsers struct {
}

// NotifyUsersTypeID is TL type id of NotifyUsers.
const NotifyUsersTypeID = 0xb4c83b4c

func (n *NotifyUsers) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotifyUsers) String() string {
	if n == nil {
		return "NotifyUsers(nil)"
	}
	type Alias NotifyUsers
	return fmt.Sprintf("NotifyUsers%+v", Alias(*n))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NotifyUsers) TypeID() uint32 {
	return NotifyUsersTypeID
}

// SchemaName returns MTProto type name.
func (n *NotifyUsers) SchemaName() string {
	return "notifyUsers"
}

// Encode implements bin.Encoder.
func (n *NotifyUsers) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notifyUsers#b4c83b4c as nil")
	}
	b.PutID(NotifyUsersTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *NotifyUsers) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notifyUsers#b4c83b4c to nil")
	}
	if err := b.ConsumeID(NotifyUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode notifyUsers#b4c83b4c: %w", err)
	}
	return nil
}

// construct implements constructor of NotifyPeerClass.
func (n NotifyUsers) construct() NotifyPeerClass { return &n }

// Ensuring interfaces in compile-time for NotifyUsers.
var (
	_ bin.Encoder = &NotifyUsers{}
	_ bin.Decoder = &NotifyUsers{}

	_ NotifyPeerClass = &NotifyUsers{}
)

// NotifyChats represents TL type `notifyChats#c007cec3`.
// Notifications generated by all groups.
//
// See https://core.telegram.org/constructor/notifyChats for reference.
type NotifyChats struct {
}

// NotifyChatsTypeID is TL type id of NotifyChats.
const NotifyChatsTypeID = 0xc007cec3

func (n *NotifyChats) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotifyChats) String() string {
	if n == nil {
		return "NotifyChats(nil)"
	}
	type Alias NotifyChats
	return fmt.Sprintf("NotifyChats%+v", Alias(*n))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NotifyChats) TypeID() uint32 {
	return NotifyChatsTypeID
}

// SchemaName returns MTProto type name.
func (n *NotifyChats) SchemaName() string {
	return "notifyChats"
}

// Encode implements bin.Encoder.
func (n *NotifyChats) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notifyChats#c007cec3 as nil")
	}
	b.PutID(NotifyChatsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *NotifyChats) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notifyChats#c007cec3 to nil")
	}
	if err := b.ConsumeID(NotifyChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode notifyChats#c007cec3: %w", err)
	}
	return nil
}

// construct implements constructor of NotifyPeerClass.
func (n NotifyChats) construct() NotifyPeerClass { return &n }

// Ensuring interfaces in compile-time for NotifyChats.
var (
	_ bin.Encoder = &NotifyChats{}
	_ bin.Decoder = &NotifyChats{}

	_ NotifyPeerClass = &NotifyChats{}
)

// NotifyBroadcasts represents TL type `notifyBroadcasts#d612e8ef`.
// Channel notification settings
//
// See https://core.telegram.org/constructor/notifyBroadcasts for reference.
type NotifyBroadcasts struct {
}

// NotifyBroadcastsTypeID is TL type id of NotifyBroadcasts.
const NotifyBroadcastsTypeID = 0xd612e8ef

func (n *NotifyBroadcasts) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *NotifyBroadcasts) String() string {
	if n == nil {
		return "NotifyBroadcasts(nil)"
	}
	type Alias NotifyBroadcasts
	return fmt.Sprintf("NotifyBroadcasts%+v", Alias(*n))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *NotifyBroadcasts) TypeID() uint32 {
	return NotifyBroadcastsTypeID
}

// SchemaName returns MTProto type name.
func (n *NotifyBroadcasts) SchemaName() string {
	return "notifyBroadcasts"
}

// Encode implements bin.Encoder.
func (n *NotifyBroadcasts) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notifyBroadcasts#d612e8ef as nil")
	}
	b.PutID(NotifyBroadcastsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *NotifyBroadcasts) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notifyBroadcasts#d612e8ef to nil")
	}
	if err := b.ConsumeID(NotifyBroadcastsTypeID); err != nil {
		return fmt.Errorf("unable to decode notifyBroadcasts#d612e8ef: %w", err)
	}
	return nil
}

// construct implements constructor of NotifyPeerClass.
func (n NotifyBroadcasts) construct() NotifyPeerClass { return &n }

// Ensuring interfaces in compile-time for NotifyBroadcasts.
var (
	_ bin.Encoder = &NotifyBroadcasts{}
	_ bin.Decoder = &NotifyBroadcasts{}

	_ NotifyPeerClass = &NotifyBroadcasts{}
)

// NotifyPeerClass represents NotifyPeer generic type.
//
// See https://core.telegram.org/type/NotifyPeer for reference.
//
// Example:
//  g, err := tg.DecodeNotifyPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.NotifyPeer: // notifyPeer#9fd40bd8
//  case *tg.NotifyUsers: // notifyUsers#b4c83b4c
//  case *tg.NotifyChats: // notifyChats#c007cec3
//  case *tg.NotifyBroadcasts: // notifyBroadcasts#d612e8ef
//  default: panic(v)
//  }
type NotifyPeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() NotifyPeerClass

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

// DecodeNotifyPeer implements binary de-serialization for NotifyPeerClass.
func DecodeNotifyPeer(buf *bin.Buffer) (NotifyPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NotifyPeerTypeID:
		// Decoding notifyPeer#9fd40bd8.
		v := NotifyPeer{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotifyPeerClass: %w", err)
		}
		return &v, nil
	case NotifyUsersTypeID:
		// Decoding notifyUsers#b4c83b4c.
		v := NotifyUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotifyPeerClass: %w", err)
		}
		return &v, nil
	case NotifyChatsTypeID:
		// Decoding notifyChats#c007cec3.
		v := NotifyChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotifyPeerClass: %w", err)
		}
		return &v, nil
	case NotifyBroadcastsTypeID:
		// Decoding notifyBroadcasts#d612e8ef.
		v := NotifyBroadcasts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NotifyPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NotifyPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// NotifyPeer boxes the NotifyPeerClass providing a helper.
type NotifyPeerBox struct {
	NotifyPeer NotifyPeerClass
}

// Decode implements bin.Decoder for NotifyPeerBox.
func (b *NotifyPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NotifyPeerBox to nil")
	}
	v, err := DecodeNotifyPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NotifyPeer = v
	return nil
}

// Encode implements bin.Encode for NotifyPeerBox.
func (b *NotifyPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NotifyPeer == nil {
		return fmt.Errorf("unable to encode NotifyPeerClass as nil")
	}
	return b.NotifyPeer.Encode(buf)
}

// NotifyPeerClassSlice is adapter for slice of NotifyPeerClass.
type NotifyPeerClassSlice []NotifyPeerClass

// First returns first element of slice (if exists).
func (s NotifyPeerClassSlice) First() (v NotifyPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s NotifyPeerClassSlice) Last() (v NotifyPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *NotifyPeerClassSlice) PopFirst() (v NotifyPeerClass, ok bool) {
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
func (s *NotifyPeerClassSlice) Pop() (v NotifyPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
