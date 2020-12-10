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

// AuthSignUpRequest represents TL type `auth.signUp#80eee427`.
// Registers a validated phone number in the system.
//
// See https://core.telegram.org/method/auth.signUp for reference.
type AuthSignUpRequest struct {
	// Phone number in the international format
	PhoneNumber string
	// SMS-message ID
	PhoneCodeHash string
	// New user first name
	FirstName string
	// New user last name
	LastName string
}

// AuthSignUpRequestTypeID is TL type id of AuthSignUpRequest.
const AuthSignUpRequestTypeID = 0x80eee427

// Encode implements bin.Encoder.
func (s *AuthSignUpRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode auth.signUp#80eee427 as nil")
	}
	b.PutID(AuthSignUpRequestTypeID)
	b.PutString(s.PhoneNumber)
	b.PutString(s.PhoneCodeHash)
	b.PutString(s.FirstName)
	b.PutString(s.LastName)
	return nil
}

// Decode implements bin.Decoder.
func (s *AuthSignUpRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode auth.signUp#80eee427 to nil")
	}
	if err := b.ConsumeID(AuthSignUpRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.signUp#80eee427: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signUp#80eee427: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signUp#80eee427: field phone_code_hash: %w", err)
		}
		s.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signUp#80eee427: field first_name: %w", err)
		}
		s.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.signUp#80eee427: field last_name: %w", err)
		}
		s.LastName = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthSignUpRequest.
var (
	_ bin.Encoder = &AuthSignUpRequest{}
	_ bin.Decoder = &AuthSignUpRequest{}
)

// AuthSignUp invokes method auth.signUp#80eee427 returning error if any.
// Registers a validated phone number in the system.
//
// See https://core.telegram.org/method/auth.signUp for reference.
func (c *Client) AuthSignUp(ctx context.Context, request *AuthSignUpRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
