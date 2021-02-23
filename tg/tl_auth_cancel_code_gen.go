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

// AuthCancelCodeRequest represents TL type `auth.cancelCode#1f040578`.
// Cancel the login verification code
//
// See https://core.telegram.org/method/auth.cancelCode for reference.
type AuthCancelCodeRequest struct {
	// Phone number
	PhoneNumber string `schemaname:"phone_number"`
	// Phone code hash from auth.sendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.sendCode
	PhoneCodeHash string `schemaname:"phone_code_hash"`
}

// AuthCancelCodeRequestTypeID is TL type id of AuthCancelCodeRequest.
const AuthCancelCodeRequestTypeID = 0x1f040578

func (c *AuthCancelCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}
	if !(c.PhoneCodeHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AuthCancelCodeRequest) String() string {
	if c == nil {
		return "AuthCancelCodeRequest(nil)"
	}
	type Alias AuthCancelCodeRequest
	return fmt.Sprintf("AuthCancelCodeRequest%+v", Alias(*c))
}

// FillFrom fills AuthCancelCodeRequest from given interface.
func (c *AuthCancelCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
}) {
	c.PhoneNumber = from.GetPhoneNumber()
	c.PhoneCodeHash = from.GetPhoneCodeHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AuthCancelCodeRequest) TypeID() uint32 {
	return AuthCancelCodeRequestTypeID
}

// SchemaName returns MTProto type name.
func (c *AuthCancelCodeRequest) SchemaName() string {
	return "auth.cancelCode"
}

// Encode implements bin.Encoder.
func (c *AuthCancelCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode auth.cancelCode#1f040578 as nil")
	}
	b.PutID(AuthCancelCodeRequestTypeID)
	b.PutString(c.PhoneNumber)
	b.PutString(c.PhoneCodeHash)
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *AuthCancelCodeRequest) GetPhoneNumber() (value string) {
	return c.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (c *AuthCancelCodeRequest) GetPhoneCodeHash() (value string) {
	return c.PhoneCodeHash
}

// Decode implements bin.Decoder.
func (c *AuthCancelCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode auth.cancelCode#1f040578 to nil")
	}
	if err := b.ConsumeID(AuthCancelCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.cancelCode#1f040578: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.cancelCode#1f040578: field phone_code_hash: %w", err)
		}
		c.PhoneCodeHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthCancelCodeRequest.
var (
	_ bin.Encoder = &AuthCancelCodeRequest{}
	_ bin.Decoder = &AuthCancelCodeRequest{}
)

// AuthCancelCode invokes method auth.cancelCode#1f040578 returning error if any.
// Cancel the login verification code
//
// Possible errors:
//  400 PHONE_CODE_EXPIRED: The phone code you provided has expired, this may happen if it was sent to any chat on telegram (if the code is sent through a telegram chat (not the official account) to avoid it append or prepend to the code some chars)
//  400 PHONE_NUMBER_INVALID: The phone number is invalid
//
// See https://core.telegram.org/method/auth.cancelCode for reference.
func (c *Client) AuthCancelCode(ctx context.Context, request *AuthCancelCodeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
