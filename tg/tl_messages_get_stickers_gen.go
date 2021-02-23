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

// MessagesGetStickersRequest represents TL type `messages.getStickers#43d4f2c`.
// Get stickers by emoji
//
// See https://core.telegram.org/method/messages.getStickers for reference.
type MessagesGetStickersRequest struct {
	// The emoji
	Emoticon string `schemaname:"emoticon"`
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int `schemaname:"hash"`
}

// MessagesGetStickersRequestTypeID is TL type id of MessagesGetStickersRequest.
const MessagesGetStickersRequestTypeID = 0x43d4f2c

func (g *MessagesGetStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Emoticon == "") {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetStickersRequest) String() string {
	if g == nil {
		return "MessagesGetStickersRequest(nil)"
	}
	type Alias MessagesGetStickersRequest
	return fmt.Sprintf("MessagesGetStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetStickersRequest from given interface.
func (g *MessagesGetStickersRequest) FillFrom(from interface {
	GetEmoticon() (value string)
	GetHash() (value int)
}) {
	g.Emoticon = from.GetEmoticon()
	g.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetStickersRequest) TypeID() uint32 {
	return MessagesGetStickersRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *MessagesGetStickersRequest) SchemaName() string {
	return "messages.getStickers"
}

// Encode implements bin.Encoder.
func (g *MessagesGetStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStickers#43d4f2c as nil")
	}
	b.PutID(MessagesGetStickersRequestTypeID)
	b.PutString(g.Emoticon)
	b.PutInt(g.Hash)
	return nil
}

// GetEmoticon returns value of Emoticon field.
func (g *MessagesGetStickersRequest) GetEmoticon() (value string) {
	return g.Emoticon
}

// GetHash returns value of Hash field.
func (g *MessagesGetStickersRequest) GetHash() (value int) {
	return g.Hash
}

// Decode implements bin.Decoder.
func (g *MessagesGetStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStickers#43d4f2c to nil")
	}
	if err := b.ConsumeID(MessagesGetStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStickers#43d4f2c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickers#43d4f2c: field emoticon: %w", err)
		}
		g.Emoticon = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStickers#43d4f2c: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStickersRequest.
var (
	_ bin.Encoder = &MessagesGetStickersRequest{}
	_ bin.Decoder = &MessagesGetStickersRequest{}
)

// MessagesGetStickers invokes method messages.getStickers#43d4f2c returning error if any.
// Get stickers by emoji
//
// See https://core.telegram.org/method/messages.getStickers for reference.
func (c *Client) MessagesGetStickers(ctx context.Context, request *MessagesGetStickersRequest) (MessagesStickersClass, error) {
	var result MessagesStickersBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Stickers, nil
}
