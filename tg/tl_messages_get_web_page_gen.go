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

// MessagesGetWebPageRequest represents TL type `messages.getWebPage#32ca8f91`.
// Get instant view page
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
type MessagesGetWebPageRequest struct {
	// URL of IV page to fetch
	URL string
	// Hash for pagination, for more info click here
	Hash int
}

// MessagesGetWebPageRequestTypeID is TL type id of MessagesGetWebPageRequest.
const MessagesGetWebPageRequestTypeID = 0x32ca8f91

// Encode implements bin.Encoder.
func (g *MessagesGetWebPageRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getWebPage#32ca8f91 as nil")
	}
	b.PutID(MessagesGetWebPageRequestTypeID)
	b.PutString(g.URL)
	b.PutInt(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetWebPageRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getWebPage#32ca8f91 to nil")
	}
	if err := b.ConsumeID(MessagesGetWebPageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field url: %w", err)
		}
		g.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getWebPage#32ca8f91: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetWebPageRequest.
var (
	_ bin.Encoder = &MessagesGetWebPageRequest{}
	_ bin.Decoder = &MessagesGetWebPageRequest{}
)

// MessagesGetWebPage invokes method messages.getWebPage#32ca8f91 returning error if any.
// Get instant view page
//
// See https://core.telegram.org/method/messages.getWebPage for reference.
func (c *Client) MessagesGetWebPage(ctx context.Context, request *MessagesGetWebPageRequest) (WebPageClass, error) {
	var result WebPageBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WebPage, nil
}
