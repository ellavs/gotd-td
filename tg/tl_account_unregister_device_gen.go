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

// AccountUnregisterDeviceRequest represents TL type `account.unregisterDevice#3076c4bf`.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
type AccountUnregisterDeviceRequest struct {
	// Device token type.Possible values:1 - APNS (device token for apple push)2 - FCM (firebase token for google firebase)3 - MPNS (channel URI for microsoft push)4 - Simple push (endpoint for firefox's simple push API)5 - Ubuntu phone (token for ubuntu push)6 - Blackberry (token for blackberry push)7 - Unused8 - WNS (windows push)9 - APNS VoIP (token for apple push VoIP)10 - Web push (web push, see below)11 - MPNS VoIP (token for microsoft push VoIP)12 - Tizen (token for tizen push)For 10 web push, the token must be a JSON-encoded object containing the keys described in PUSH updates
	TokenType int
	// Device token
	Token string
	// List of user identifiers of other users currently using the client
	OtherUids []int
}

// AccountUnregisterDeviceRequestTypeID is TL type id of AccountUnregisterDeviceRequest.
const AccountUnregisterDeviceRequestTypeID = 0x3076c4bf

// Encode implements bin.Encoder.
func (u *AccountUnregisterDeviceRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.unregisterDevice#3076c4bf as nil")
	}
	b.PutID(AccountUnregisterDeviceRequestTypeID)
	b.PutInt(u.TokenType)
	b.PutString(u.Token)
	b.PutVectorHeader(len(u.OtherUids))
	for _, v := range u.OtherUids {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUnregisterDeviceRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.unregisterDevice#3076c4bf to nil")
	}
	if err := b.ConsumeID(AccountUnregisterDeviceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field token_type: %w", err)
		}
		u.TokenType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field token: %w", err)
		}
		u.Token = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field other_uids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode account.unregisterDevice#3076c4bf: field other_uids: %w", err)
			}
			u.OtherUids = append(u.OtherUids, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUnregisterDeviceRequest.
var (
	_ bin.Encoder = &AccountUnregisterDeviceRequest{}
	_ bin.Decoder = &AccountUnregisterDeviceRequest{}
)

// AccountUnregisterDevice invokes method account.unregisterDevice#3076c4bf returning error if any.
// Deletes a device by its token, stops sending PUSH-notifications to it.
//
// See https://core.telegram.org/method/account.unregisterDevice for reference.
func (c *Client) AccountUnregisterDevice(ctx context.Context, request *AccountUnregisterDeviceRequest) (BoolClass, error) {
	var result BoolBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Bool, nil
}
