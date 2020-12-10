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

// MessagesGetScheduledMessagesRequest represents TL type `messages.getScheduledMessages#bdbb0464`.
// Get scheduled messages
//
// See https://core.telegram.org/method/messages.getScheduledMessages for reference.
type MessagesGetScheduledMessagesRequest struct {
	// Peer
	Peer InputPeerClass
	// IDs of scheduled messages
	ID []int
}

// MessagesGetScheduledMessagesRequestTypeID is TL type id of MessagesGetScheduledMessagesRequest.
const MessagesGetScheduledMessagesRequestTypeID = 0xbdbb0464

// Encode implements bin.Encoder.
func (g *MessagesGetScheduledMessagesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getScheduledMessages#bdbb0464 as nil")
	}
	b.PutID(MessagesGetScheduledMessagesRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getScheduledMessages#bdbb0464: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getScheduledMessages#bdbb0464: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetScheduledMessagesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getScheduledMessages#bdbb0464 to nil")
	}
	if err := b.ConsumeID(MessagesGetScheduledMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getScheduledMessages#bdbb0464: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetScheduledMessagesRequest.
var (
	_ bin.Encoder = &MessagesGetScheduledMessagesRequest{}
	_ bin.Decoder = &MessagesGetScheduledMessagesRequest{}
)

// MessagesGetScheduledMessages invokes method messages.getScheduledMessages#bdbb0464 returning error if any.
// Get scheduled messages
//
// See https://core.telegram.org/method/messages.getScheduledMessages for reference.
func (c *Client) MessagesGetScheduledMessages(ctx context.Context, request *MessagesGetScheduledMessagesRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
