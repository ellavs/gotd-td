// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// AccountSaveAutoSaveSettingsRequest represents TL type `account.saveAutoSaveSettings#d69b8361`.
//
// See https://core.telegram.org/method/account.saveAutoSaveSettings for reference.
type AccountSaveAutoSaveSettingsRequest struct {
	// Flags field of AccountSaveAutoSaveSettingsRequest.
	Flags bin.Fields
	// Users field of AccountSaveAutoSaveSettingsRequest.
	Users bool
	// Chats field of AccountSaveAutoSaveSettingsRequest.
	Chats bool
	// Broadcasts field of AccountSaveAutoSaveSettingsRequest.
	Broadcasts bool
	// Peer field of AccountSaveAutoSaveSettingsRequest.
	//
	// Use SetPeer and GetPeer helpers.
	Peer InputPeerClass
	// Settings field of AccountSaveAutoSaveSettingsRequest.
	Settings AutoSaveSettings
}

// AccountSaveAutoSaveSettingsRequestTypeID is TL type id of AccountSaveAutoSaveSettingsRequest.
const AccountSaveAutoSaveSettingsRequestTypeID = 0xd69b8361

// Ensuring interfaces in compile-time for AccountSaveAutoSaveSettingsRequest.
var (
	_ bin.Encoder     = &AccountSaveAutoSaveSettingsRequest{}
	_ bin.Decoder     = &AccountSaveAutoSaveSettingsRequest{}
	_ bin.BareEncoder = &AccountSaveAutoSaveSettingsRequest{}
	_ bin.BareDecoder = &AccountSaveAutoSaveSettingsRequest{}
)

func (s *AccountSaveAutoSaveSettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Users == false) {
		return false
	}
	if !(s.Chats == false) {
		return false
	}
	if !(s.Broadcasts == false) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSaveAutoSaveSettingsRequest) String() string {
	if s == nil {
		return "AccountSaveAutoSaveSettingsRequest(nil)"
	}
	type Alias AccountSaveAutoSaveSettingsRequest
	return fmt.Sprintf("AccountSaveAutoSaveSettingsRequest%+v", Alias(*s))
}

// FillFrom fills AccountSaveAutoSaveSettingsRequest from given interface.
func (s *AccountSaveAutoSaveSettingsRequest) FillFrom(from interface {
	GetUsers() (value bool)
	GetChats() (value bool)
	GetBroadcasts() (value bool)
	GetPeer() (value InputPeerClass, ok bool)
	GetSettings() (value AutoSaveSettings)
}) {
	s.Users = from.GetUsers()
	s.Chats = from.GetChats()
	s.Broadcasts = from.GetBroadcasts()
	if val, ok := from.GetPeer(); ok {
		s.Peer = val
	}

	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSaveAutoSaveSettingsRequest) TypeID() uint32 {
	return AccountSaveAutoSaveSettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSaveAutoSaveSettingsRequest) TypeName() string {
	return "account.saveAutoSaveSettings"
}

// TypeInfo returns info about TL type.
func (s *AccountSaveAutoSaveSettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.saveAutoSaveSettings",
		ID:   AccountSaveAutoSaveSettingsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Users",
			SchemaName: "users",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Broadcasts",
			SchemaName: "broadcasts",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *AccountSaveAutoSaveSettingsRequest) SetFlags() {
	if !(s.Users == false) {
		s.Flags.Set(0)
	}
	if !(s.Chats == false) {
		s.Flags.Set(1)
	}
	if !(s.Broadcasts == false) {
		s.Flags.Set(2)
	}
	if !(s.Peer == nil) {
		s.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (s *AccountSaveAutoSaveSettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveAutoSaveSettings#d69b8361 as nil")
	}
	b.PutID(AccountSaveAutoSaveSettingsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSaveAutoSaveSettingsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveAutoSaveSettings#d69b8361 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoSaveSettings#d69b8361: field flags: %w", err)
	}
	if s.Flags.Has(3) {
		if s.Peer == nil {
			return fmt.Errorf("unable to encode account.saveAutoSaveSettings#d69b8361: field peer is nil")
		}
		if err := s.Peer.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.saveAutoSaveSettings#d69b8361: field peer: %w", err)
		}
	}
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveAutoSaveSettings#d69b8361: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveAutoSaveSettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveAutoSaveSettings#d69b8361 to nil")
	}
	if err := b.ConsumeID(AccountSaveAutoSaveSettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveAutoSaveSettings#d69b8361: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSaveAutoSaveSettingsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveAutoSaveSettings#d69b8361 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoSaveSettings#d69b8361: field flags: %w", err)
		}
	}
	s.Users = s.Flags.Has(0)
	s.Chats = s.Flags.Has(1)
	s.Broadcasts = s.Flags.Has(2)
	if s.Flags.Has(3) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveAutoSaveSettings#d69b8361: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveAutoSaveSettings#d69b8361: field settings: %w", err)
		}
	}
	return nil
}

// SetUsers sets value of Users conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) SetUsers(value bool) {
	if value {
		s.Flags.Set(0)
		s.Users = true
	} else {
		s.Flags.Unset(0)
		s.Users = false
	}
}

// GetUsers returns value of Users conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) GetUsers() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(0)
}

// SetChats sets value of Chats conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) SetChats(value bool) {
	if value {
		s.Flags.Set(1)
		s.Chats = true
	} else {
		s.Flags.Unset(1)
		s.Chats = false
	}
}

// GetChats returns value of Chats conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) GetChats() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetBroadcasts sets value of Broadcasts conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) SetBroadcasts(value bool) {
	if value {
		s.Flags.Set(2)
		s.Broadcasts = true
	} else {
		s.Flags.Unset(2)
		s.Broadcasts = false
	}
}

// GetBroadcasts returns value of Broadcasts conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) GetBroadcasts() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(2)
}

// SetPeer sets value of Peer conditional field.
func (s *AccountSaveAutoSaveSettingsRequest) SetPeer(value InputPeerClass) {
	s.Flags.Set(3)
	s.Peer = value
}

// GetPeer returns value of Peer conditional field and
// boolean which is true if field was set.
func (s *AccountSaveAutoSaveSettingsRequest) GetPeer() (value InputPeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Peer, true
}

// GetSettings returns value of Settings field.
func (s *AccountSaveAutoSaveSettingsRequest) GetSettings() (value AutoSaveSettings) {
	if s == nil {
		return
	}
	return s.Settings
}

// AccountSaveAutoSaveSettings invokes method account.saveAutoSaveSettings#d69b8361 returning error if any.
//
// See https://core.telegram.org/method/account.saveAutoSaveSettings for reference.
func (c *Client) AccountSaveAutoSaveSettings(ctx context.Context, request *AccountSaveAutoSaveSettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
