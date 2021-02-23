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

// MessagesGetRecentLocationsRequest represents TL type `messages.getRecentLocations#bbc45b09`.
// Get live location history of a certain user
//
// See https://core.telegram.org/method/messages.getRecentLocations for reference.
type MessagesGetRecentLocationsRequest struct {
	// User
	Peer InputPeerClass `schemaname:"peer"`
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int `schemaname:"limit"`
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int `schemaname:"hash"`
}

// MessagesGetRecentLocationsRequestTypeID is TL type id of MessagesGetRecentLocationsRequest.
const MessagesGetRecentLocationsRequestTypeID = 0xbbc45b09

func (g *MessagesGetRecentLocationsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetRecentLocationsRequest) String() string {
	if g == nil {
		return "MessagesGetRecentLocationsRequest(nil)"
	}
	type Alias MessagesGetRecentLocationsRequest
	return fmt.Sprintf("MessagesGetRecentLocationsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetRecentLocationsRequest from given interface.
func (g *MessagesGetRecentLocationsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetLimit() (value int)
	GetHash() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetRecentLocationsRequest) TypeID() uint32 {
	return MessagesGetRecentLocationsRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *MessagesGetRecentLocationsRequest) SchemaName() string {
	return "messages.getRecentLocations"
}

// Encode implements bin.Encoder.
func (g *MessagesGetRecentLocationsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getRecentLocations#bbc45b09 as nil")
	}
	b.PutID(MessagesGetRecentLocationsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getRecentLocations#bbc45b09: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getRecentLocations#bbc45b09: field peer: %w", err)
	}
	b.PutInt(g.Limit)
	b.PutInt(g.Hash)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetRecentLocationsRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetLimit returns value of Limit field.
func (g *MessagesGetRecentLocationsRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetRecentLocationsRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetRecentLocationsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getRecentLocations#bbc45b09 to nil")
	}
	if err := b.ConsumeID(MessagesGetRecentLocationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getRecentLocations#bbc45b09: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#bbc45b09: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#bbc45b09: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getRecentLocations#bbc45b09: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetRecentLocationsRequest.
var (
	_ bin.Encoder = &MessagesGetRecentLocationsRequest{}
	_ bin.Decoder = &MessagesGetRecentLocationsRequest{}
)

// MessagesGetRecentLocations invokes method messages.getRecentLocations#bbc45b09 returning error if any.
// Get live location history of a certain user
//
// See https://core.telegram.org/method/messages.getRecentLocations for reference.
func (c *Client) MessagesGetRecentLocations(ctx context.Context, request *MessagesGetRecentLocationsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
