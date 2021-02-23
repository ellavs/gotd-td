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

// AccountResetWebAuthorizationRequest represents TL type `account.resetWebAuthorization#2d01b9ef`.
// Log out an active web telegram login¹ session
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorization for reference.
type AccountResetWebAuthorizationRequest struct {
	// Session¹ hash
	//
	// Links:
	//  1) https://core.telegram.org/constructor/webAuthorization
	Hash int64 `schemaname:"hash"`
}

// AccountResetWebAuthorizationRequestTypeID is TL type id of AccountResetWebAuthorizationRequest.
const AccountResetWebAuthorizationRequestTypeID = 0x2d01b9ef

func (r *AccountResetWebAuthorizationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetWebAuthorizationRequest) String() string {
	if r == nil {
		return "AccountResetWebAuthorizationRequest(nil)"
	}
	type Alias AccountResetWebAuthorizationRequest
	return fmt.Sprintf("AccountResetWebAuthorizationRequest%+v", Alias(*r))
}

// FillFrom fills AccountResetWebAuthorizationRequest from given interface.
func (r *AccountResetWebAuthorizationRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	r.Hash = from.GetHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AccountResetWebAuthorizationRequest) TypeID() uint32 {
	return AccountResetWebAuthorizationRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *AccountResetWebAuthorizationRequest) SchemaName() string {
	return "account.resetWebAuthorization"
}

// Encode implements bin.Encoder.
func (r *AccountResetWebAuthorizationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorization#2d01b9ef as nil")
	}
	b.PutID(AccountResetWebAuthorizationRequestTypeID)
	b.PutLong(r.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (r *AccountResetWebAuthorizationRequest) GetHash() (value int64) {
	return r.Hash
}

// Decode implements bin.Decoder.
func (r *AccountResetWebAuthorizationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorization#2d01b9ef to nil")
	}
	if err := b.ConsumeID(AccountResetWebAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWebAuthorization#2d01b9ef: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.resetWebAuthorization#2d01b9ef: field hash: %w", err)
		}
		r.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetWebAuthorizationRequest.
var (
	_ bin.Encoder = &AccountResetWebAuthorizationRequest{}
	_ bin.Decoder = &AccountResetWebAuthorizationRequest{}
)

// AccountResetWebAuthorization invokes method account.resetWebAuthorization#2d01b9ef returning error if any.
// Log out an active web telegram login¹ session
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorization for reference.
func (c *Client) AccountResetWebAuthorization(ctx context.Context, hash int64) (bool, error) {
	var result BoolBox

	request := &AccountResetWebAuthorizationRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
