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

// AuthImportBotAuthorizationRequest represents TL type `auth.importBotAuthorization#67a3ff2c`.
// Login as a bot
//
// See https://core.telegram.org/method/auth.importBotAuthorization for reference.
type AuthImportBotAuthorizationRequest struct {
	// Reserved for future use
	Flags int
	// Application identifier (see. App configuration)
	APIID int
	// Application identifier hash (see. App configuration)
	APIHash string
	// Bot token (see bots)
	BotAuthToken string
}

// AuthImportBotAuthorizationRequestTypeID is TL type id of AuthImportBotAuthorizationRequest.
const AuthImportBotAuthorizationRequestTypeID = 0x67a3ff2c

// Encode implements bin.Encoder.
func (i *AuthImportBotAuthorizationRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode auth.importBotAuthorization#67a3ff2c as nil")
	}
	b.PutID(AuthImportBotAuthorizationRequestTypeID)
	b.PutInt(i.Flags)
	b.PutInt(i.APIID)
	b.PutString(i.APIHash)
	b.PutString(i.BotAuthToken)
	return nil
}

// Decode implements bin.Decoder.
func (i *AuthImportBotAuthorizationRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode auth.importBotAuthorization#67a3ff2c to nil")
	}
	if err := b.ConsumeID(AuthImportBotAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field flags: %w", err)
		}
		i.Flags = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field api_id: %w", err)
		}
		i.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field api_hash: %w", err)
		}
		i.APIHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.importBotAuthorization#67a3ff2c: field bot_auth_token: %w", err)
		}
		i.BotAuthToken = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthImportBotAuthorizationRequest.
var (
	_ bin.Encoder = &AuthImportBotAuthorizationRequest{}
	_ bin.Decoder = &AuthImportBotAuthorizationRequest{}
)

// AuthImportBotAuthorization invokes method auth.importBotAuthorization#67a3ff2c returning error if any.
// Login as a bot
//
// See https://core.telegram.org/method/auth.importBotAuthorization for reference.
func (c *Client) AuthImportBotAuthorization(ctx context.Context, request *AuthImportBotAuthorizationRequest) (AuthAuthorizationClass, error) {
	var result AuthAuthorizationBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Authorization, nil
}
