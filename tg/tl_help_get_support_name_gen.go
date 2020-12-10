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

// HelpGetSupportNameRequest represents TL type `help.getSupportName#d360e72c`.
// Get localized name of the telegram support user
//
// See https://core.telegram.org/method/help.getSupportName for reference.
type HelpGetSupportNameRequest struct {
}

// HelpGetSupportNameRequestTypeID is TL type id of HelpGetSupportNameRequest.
const HelpGetSupportNameRequestTypeID = 0xd360e72c

// Encode implements bin.Encoder.
func (g *HelpGetSupportNameRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getSupportName#d360e72c as nil")
	}
	b.PutID(HelpGetSupportNameRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetSupportNameRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getSupportName#d360e72c to nil")
	}
	if err := b.ConsumeID(HelpGetSupportNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getSupportName#d360e72c: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetSupportNameRequest.
var (
	_ bin.Encoder = &HelpGetSupportNameRequest{}
	_ bin.Decoder = &HelpGetSupportNameRequest{}
)

// HelpGetSupportName invokes method help.getSupportName#d360e72c returning error if any.
// Get localized name of the telegram support user
//
// See https://core.telegram.org/method/help.getSupportName for reference.
func (c *Client) HelpGetSupportName(ctx context.Context, request *HelpGetSupportNameRequest) (*HelpSupportName, error) {
	var result HelpSupportName
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
