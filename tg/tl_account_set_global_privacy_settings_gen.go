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

// AccountSetGlobalPrivacySettingsRequest represents TL type `account.setGlobalPrivacySettings#1edaaac2`.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
type AccountSetGlobalPrivacySettingsRequest struct {
	// Global privacy settings
	Settings GlobalPrivacySettings
}

// AccountSetGlobalPrivacySettingsRequestTypeID is TL type id of AccountSetGlobalPrivacySettingsRequest.
const AccountSetGlobalPrivacySettingsRequestTypeID = 0x1edaaac2

// Encode implements bin.Encoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setGlobalPrivacySettings#1edaaac2 as nil")
	}
	b.PutID(AccountSetGlobalPrivacySettingsRequestTypeID)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setGlobalPrivacySettings#1edaaac2 to nil")
	}
	if err := b.ConsumeID(AccountSetGlobalPrivacySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: %w", err)
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetGlobalPrivacySettingsRequest.
var (
	_ bin.Encoder = &AccountSetGlobalPrivacySettingsRequest{}
	_ bin.Decoder = &AccountSetGlobalPrivacySettingsRequest{}
)

// AccountSetGlobalPrivacySettings invokes method account.setGlobalPrivacySettings#1edaaac2 returning error if any.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
func (c *Client) AccountSetGlobalPrivacySettings(ctx context.Context, request *AccountSetGlobalPrivacySettingsRequest) (*GlobalPrivacySettings, error) {
	var result GlobalPrivacySettings
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
