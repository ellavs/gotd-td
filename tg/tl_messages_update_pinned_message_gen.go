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

// MessagesUpdatePinnedMessageRequest represents TL type `messages.updatePinnedMessage#d2aaf7ec`.
// Pin a message
//
// See https://core.telegram.org/method/messages.updatePinnedMessage for reference.
type MessagesUpdatePinnedMessageRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Pin the message silently, without triggering a notification
	Silent bool
	// Whether the message should unpinned or pinned
	Unpin bool
	// Whether the message should only be pinned on the local side of a one-to-one chat
	PmOneside bool
	// The peer where to pin the message
	Peer InputPeerClass
	// The message to pin or unpin
	ID int
}

// MessagesUpdatePinnedMessageRequestTypeID is TL type id of MessagesUpdatePinnedMessageRequest.
const MessagesUpdatePinnedMessageRequestTypeID = 0xd2aaf7ec

// Encode implements bin.Encoder.
func (u *MessagesUpdatePinnedMessageRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.updatePinnedMessage#d2aaf7ec as nil")
	}
	b.PutID(MessagesUpdatePinnedMessageRequestTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
	}
	b.PutInt(u.ID)
	return nil
}

// SetSilent sets value of Silent conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetSilent(value bool) {
	if value {
		u.Flags.Set(0)
	} else {
		u.Flags.Unset(0)
	}
}

// SetUnpin sets value of Unpin conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetUnpin(value bool) {
	if value {
		u.Flags.Set(1)
	} else {
		u.Flags.Unset(1)
	}
}

// SetPmOneside sets value of PmOneside conditional field.
func (u *MessagesUpdatePinnedMessageRequest) SetPmOneside(value bool) {
	if value {
		u.Flags.Set(2)
	} else {
		u.Flags.Unset(2)
	}
}

// Decode implements bin.Decoder.
func (u *MessagesUpdatePinnedMessageRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.updatePinnedMessage#d2aaf7ec to nil")
	}
	if err := b.ConsumeID(MessagesUpdatePinnedMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field flags: %w", err)
		}
	}
	u.Silent = u.Flags.Has(0)
	u.Unpin = u.Flags.Has(1)
	u.PmOneside = u.Flags.Has(2)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.updatePinnedMessage#d2aaf7ec: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUpdatePinnedMessageRequest.
var (
	_ bin.Encoder = &MessagesUpdatePinnedMessageRequest{}
	_ bin.Decoder = &MessagesUpdatePinnedMessageRequest{}
)

// MessagesUpdatePinnedMessage invokes method messages.updatePinnedMessage#d2aaf7ec returning error if any.
// Pin a message
//
// See https://core.telegram.org/method/messages.updatePinnedMessage for reference.
func (c *Client) MessagesUpdatePinnedMessage(ctx context.Context, request *MessagesUpdatePinnedMessageRequest) (UpdatesClass, error) {
	var result UpdatesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
