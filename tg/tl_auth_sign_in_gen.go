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

// AuthSignInRequest represents TL type `auth.signIn#bcd51581`.
// Signs in a user with a validated phone number.
//
// See https://core.telegram.org/method/auth.signIn for reference.
type AuthSignInRequest struct {
	// Phone number in the international format
	PhoneNumber string `schemaname:"phone_number"`
	// SMS-message ID, obtained from auth.sendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.sendCode
	PhoneCodeHash string `schemaname:"phone_code_hash"`
	// Valid numerical code from the SMS-message
	PhoneCode string `schemaname:"phone_code"`
}

// AuthSignInRequestTypeID is TL type id of AuthSignInRequest.
const AuthSignInRequestTypeID = 0xbcd51581

func (s *AuthSignInRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.PhoneCodeHash == "") {
		return false
	}
	if !(s.PhoneCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AuthSignInRequest) String() string {
	if s == nil {
		return "AuthSignInRequest(nil)"
	}
	type Alias AuthSignInRequest
	return fmt.Sprintf("AuthSignInRequest%+v", Alias(*s))
}

// FillFrom fills AuthSignInRequest from given interface.
func (s *AuthSignInRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetPhoneCode() (value string)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.PhoneCodeHash = from.GetPhoneCodeHash()
	s.PhoneCode = from.GetPhoneCode()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AuthSignInRequest) TypeID() uint32 {
	return AuthSignInRequestTypeID
}

// SchemaName returns MTProto type name.
func (s *AuthSignInRequest) SchemaName() string {
	return "auth.signIn"
}

// Encode implements bin.Encoder.
func (s *AuthSignInRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.signIn#bcd51581 as nil")
	}
	b.PutID(AuthSignInRequestTypeID)
	b.PutString(s.PhoneNumber)
	b.PutString(s.PhoneCodeHash)
	b.PutString(s.PhoneCode)
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *AuthSignInRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (s *AuthSignInRequest) GetPhoneCodeHash() (value string) {
	return s.PhoneCodeHash
}

// GetPhoneCode returns value of PhoneCode field.
func (s *AuthSignInRequest) GetPhoneCode() (value string) {
	return s.PhoneCode
}

// Decode implements bin.Decoder.
func (s *AuthSignInRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.signIn#bcd51581 to nil")
	}
	if err := b.ConsumeID(AuthSignInRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.signIn#bcd51581: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_code_hash: %w", err)
		}
		s.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signIn#bcd51581: field phone_code: %w", err)
		}
		s.PhoneCode = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthSignInRequest.
var (
	_ bin.Encoder = &AuthSignInRequest{}
	_ bin.Decoder = &AuthSignInRequest{}
)

// AuthSignIn invokes method auth.signIn#bcd51581 returning error if any.
// Signs in a user with a validated phone number.
//
// Possible errors:
//  400 PHONE_CODE_EMPTY: phone_code from the SMS is empty
//  400 PHONE_CODE_EXPIRED: SMS expired
//  400 PHONE_CODE_INVALID: Invalid SMS code was sent
//  400 PHONE_NUMBER_INVALID: Invalid phone number
//  400 PHONE_NUMBER_UNOCCUPIED: The code is valid but no user with the given number is registered
//
// See https://core.telegram.org/method/auth.signIn for reference.
func (c *Client) AuthSignIn(ctx context.Context, request *AuthSignInRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
