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

// AccountGetAccountTTLRequest represents TL type `account.getAccountTTL#8fc711d`.
// Get days to live of account
//
// See https://core.telegram.org/method/account.getAccountTTL for reference.
type AccountGetAccountTTLRequest struct {
}

// AccountGetAccountTTLRequestTypeID is TL type id of AccountGetAccountTTLRequest.
const AccountGetAccountTTLRequestTypeID = 0x8fc711d

func (g *AccountGetAccountTTLRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAccountTTLRequest) String() string {
	if g == nil {
		return "AccountGetAccountTTLRequest(nil)"
	}
	type Alias AccountGetAccountTTLRequest
	return fmt.Sprintf("AccountGetAccountTTLRequest%+v", Alias(*g))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetAccountTTLRequest) TypeID() uint32 {
	return AccountGetAccountTTLRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *AccountGetAccountTTLRequest) SchemaName() string {
	return "account.getAccountTTL"
}

// Encode implements bin.Encoder.
func (g *AccountGetAccountTTLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAccountTTL#8fc711d as nil")
	}
	b.PutID(AccountGetAccountTTLRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAccountTTLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAccountTTL#8fc711d to nil")
	}
	if err := b.ConsumeID(AccountGetAccountTTLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAccountTTL#8fc711d: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetAccountTTLRequest.
var (
	_ bin.Encoder = &AccountGetAccountTTLRequest{}
	_ bin.Decoder = &AccountGetAccountTTLRequest{}
)

// AccountGetAccountTTL invokes method account.getAccountTTL#8fc711d returning error if any.
// Get days to live of account
//
// See https://core.telegram.org/method/account.getAccountTTL for reference.
func (c *Client) AccountGetAccountTTL(ctx context.Context) (*AccountDaysTTL, error) {
	var result AccountDaysTTL

	request := &AccountGetAccountTTLRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
