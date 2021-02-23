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

// PhoneRequestCallRequest represents TL type `phone.requestCall#42ff96ed`.
// Start a telegram phone call
//
// See https://core.telegram.org/method/phone.requestCall for reference.
type PhoneRequestCallRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// Whether to start a video call
	Video bool `schemaname:"video"`
	// Destination of the phone call
	UserID InputUserClass `schemaname:"user_id"`
	// Random ID to avoid resending the same object
	RandomID int `schemaname:"random_id"`
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GAHash []byte `schemaname:"g_a_hash"`
	// Phone call settings
	Protocol PhoneCallProtocol `schemaname:"protocol"`
}

// PhoneRequestCallRequestTypeID is TL type id of PhoneRequestCallRequest.
const PhoneRequestCallRequestTypeID = 0x42ff96ed

func (r *PhoneRequestCallRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Video == false) {
		return false
	}
	if !(r.UserID == nil) {
		return false
	}
	if !(r.RandomID == 0) {
		return false
	}
	if !(r.GAHash == nil) {
		return false
	}
	if !(r.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *PhoneRequestCallRequest) String() string {
	if r == nil {
		return "PhoneRequestCallRequest(nil)"
	}
	type Alias PhoneRequestCallRequest
	return fmt.Sprintf("PhoneRequestCallRequest%+v", Alias(*r))
}

// FillFrom fills PhoneRequestCallRequest from given interface.
func (r *PhoneRequestCallRequest) FillFrom(from interface {
	GetVideo() (value bool)
	GetUserID() (value InputUserClass)
	GetRandomID() (value int)
	GetGAHash() (value []byte)
	GetProtocol() (value PhoneCallProtocol)
}) {
	r.Video = from.GetVideo()
	r.UserID = from.GetUserID()
	r.RandomID = from.GetRandomID()
	r.GAHash = from.GetGAHash()
	r.Protocol = from.GetProtocol()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *PhoneRequestCallRequest) TypeID() uint32 {
	return PhoneRequestCallRequestTypeID
}

// SchemaName returns MTProto type name.
func (r *PhoneRequestCallRequest) SchemaName() string {
	return "phone.requestCall"
}

// Encode implements bin.Encoder.
func (r *PhoneRequestCallRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode phone.requestCall#42ff96ed as nil")
	}
	b.PutID(PhoneRequestCallRequestTypeID)
	if !(r.Video == false) {
		r.Flags.Set(0)
	}
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field flags: %w", err)
	}
	if r.UserID == nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field user_id: %w", err)
	}
	b.PutInt(r.RandomID)
	b.PutBytes(r.GAHash)
	if err := r.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.requestCall#42ff96ed: field protocol: %w", err)
	}
	return nil
}

// SetVideo sets value of Video conditional field.
func (r *PhoneRequestCallRequest) SetVideo(value bool) {
	if value {
		r.Flags.Set(0)
		r.Video = true
	} else {
		r.Flags.Unset(0)
		r.Video = false
	}
}

// GetVideo returns value of Video conditional field.
func (r *PhoneRequestCallRequest) GetVideo() (value bool) {
	return r.Flags.Has(0)
}

// GetUserID returns value of UserID field.
func (r *PhoneRequestCallRequest) GetUserID() (value InputUserClass) {
	return r.UserID
}

// GetRandomID returns value of RandomID field.
func (r *PhoneRequestCallRequest) GetRandomID() (value int) {
	return r.RandomID
}

// GetGAHash returns value of GAHash field.
func (r *PhoneRequestCallRequest) GetGAHash() (value []byte) {
	return r.GAHash
}

// GetProtocol returns value of Protocol field.
func (r *PhoneRequestCallRequest) GetProtocol() (value PhoneCallProtocol) {
	return r.Protocol
}

// Decode implements bin.Decoder.
func (r *PhoneRequestCallRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode phone.requestCall#42ff96ed to nil")
	}
	if err := b.ConsumeID(PhoneRequestCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: %w", err)
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field flags: %w", err)
		}
	}
	r.Video = r.Flags.Has(0)
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field random_id: %w", err)
		}
		r.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field g_a_hash: %w", err)
		}
		r.GAHash = value
	}
	{
		if err := r.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.requestCall#42ff96ed: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneRequestCallRequest.
var (
	_ bin.Encoder = &PhoneRequestCallRequest{}
	_ bin.Decoder = &PhoneRequestCallRequest{}
)

// PhoneRequestCall invokes method phone.requestCall#42ff96ed returning error if any.
// Start a telegram phone call
//
// Possible errors:
//  400 CALL_PROTOCOL_FLAGS_INVALID: Call protocol flags invalid
//  400 PARTICIPANT_VERSION_OUTDATED: The other participant does not use an up to date telegram client with support for calls
//  400 USER_ID_INVALID: The provided user ID is invalid
//  403 USER_IS_BLOCKED: You were blocked by this user
//  403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this
//
// See https://core.telegram.org/method/phone.requestCall for reference.
func (c *Client) PhoneRequestCall(ctx context.Context, request *PhoneRequestCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
