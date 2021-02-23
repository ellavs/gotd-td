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

// InputPeerEmpty represents TL type `inputPeerEmpty#7f3b18ea`.
// An empty constructor, no user or chat is defined.
//
// See https://core.telegram.org/constructor/inputPeerEmpty for reference.
type InputPeerEmpty struct {
}

// InputPeerEmptyTypeID is TL type id of InputPeerEmpty.
const InputPeerEmptyTypeID = 0x7f3b18ea

func (i *InputPeerEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerEmpty) String() string {
	if i == nil {
		return "InputPeerEmpty(nil)"
	}
	type Alias InputPeerEmpty
	return fmt.Sprintf("InputPeerEmpty%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerEmpty) TypeID() uint32 {
	return InputPeerEmptyTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerEmpty) SchemaName() string {
	return "inputPeerEmpty"
}

// Encode implements bin.Encoder.
func (i *InputPeerEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerEmpty#7f3b18ea as nil")
	}
	b.PutID(InputPeerEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerEmpty#7f3b18ea to nil")
	}
	if err := b.ConsumeID(InputPeerEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerEmpty#7f3b18ea: %w", err)
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerEmpty) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerEmpty.
var (
	_ bin.Encoder = &InputPeerEmpty{}
	_ bin.Decoder = &InputPeerEmpty{}

	_ InputPeerClass = &InputPeerEmpty{}
)

// InputPeerSelf represents TL type `inputPeerSelf#7da07ec9`.
// Defines the current user.
//
// See https://core.telegram.org/constructor/inputPeerSelf for reference.
type InputPeerSelf struct {
}

// InputPeerSelfTypeID is TL type id of InputPeerSelf.
const InputPeerSelfTypeID = 0x7da07ec9

func (i *InputPeerSelf) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerSelf) String() string {
	if i == nil {
		return "InputPeerSelf(nil)"
	}
	type Alias InputPeerSelf
	return fmt.Sprintf("InputPeerSelf%+v", Alias(*i))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerSelf) TypeID() uint32 {
	return InputPeerSelfTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerSelf) SchemaName() string {
	return "inputPeerSelf"
}

// Encode implements bin.Encoder.
func (i *InputPeerSelf) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerSelf#7da07ec9 as nil")
	}
	b.PutID(InputPeerSelfTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPeerSelf) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerSelf#7da07ec9 to nil")
	}
	if err := b.ConsumeID(InputPeerSelfTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerSelf#7da07ec9: %w", err)
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerSelf) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerSelf.
var (
	_ bin.Encoder = &InputPeerSelf{}
	_ bin.Decoder = &InputPeerSelf{}

	_ InputPeerClass = &InputPeerSelf{}
)

// InputPeerChat represents TL type `inputPeerChat#179be863`.
// Defines a chat for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerChat for reference.
type InputPeerChat struct {
	// Chat idientifier
	ChatID int `schemaname:"chat_id"`
}

// InputPeerChatTypeID is TL type id of InputPeerChat.
const InputPeerChatTypeID = 0x179be863

func (i *InputPeerChat) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChat) String() string {
	if i == nil {
		return "InputPeerChat(nil)"
	}
	type Alias InputPeerChat
	return fmt.Sprintf("InputPeerChat%+v", Alias(*i))
}

// FillFrom fills InputPeerChat from given interface.
func (i *InputPeerChat) FillFrom(from interface {
	GetChatID() (value int)
}) {
	i.ChatID = from.GetChatID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerChat) TypeID() uint32 {
	return InputPeerChatTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerChat) SchemaName() string {
	return "inputPeerChat"
}

// Encode implements bin.Encoder.
func (i *InputPeerChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChat#179be863 as nil")
	}
	b.PutID(InputPeerChatTypeID)
	b.PutInt(i.ChatID)
	return nil
}

// GetChatID returns value of ChatID field.
func (i *InputPeerChat) GetChatID() (value int) {
	return i.ChatID
}

// Decode implements bin.Decoder.
func (i *InputPeerChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChat#179be863 to nil")
	}
	if err := b.ConsumeID(InputPeerChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChat#179be863: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChat#179be863: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerChat) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChat.
var (
	_ bin.Encoder = &InputPeerChat{}
	_ bin.Decoder = &InputPeerChat{}

	_ InputPeerClass = &InputPeerChat{}
)

// InputPeerUser represents TL type `inputPeerUser#7b8e7de6`.
// Defines a user for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerUser for reference.
type InputPeerUser struct {
	// User identifier
	UserID int `schemaname:"user_id"`
	// access_hash value from the user¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/user
	AccessHash int64 `schemaname:"access_hash"`
}

// InputPeerUserTypeID is TL type id of InputPeerUser.
const InputPeerUserTypeID = 0x7b8e7de6

func (i *InputPeerUser) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.UserID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerUser) String() string {
	if i == nil {
		return "InputPeerUser(nil)"
	}
	type Alias InputPeerUser
	return fmt.Sprintf("InputPeerUser%+v", Alias(*i))
}

// FillFrom fills InputPeerUser from given interface.
func (i *InputPeerUser) FillFrom(from interface {
	GetUserID() (value int)
	GetAccessHash() (value int64)
}) {
	i.UserID = from.GetUserID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerUser) TypeID() uint32 {
	return InputPeerUserTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerUser) SchemaName() string {
	return "inputPeerUser"
}

// Encode implements bin.Encoder.
func (i *InputPeerUser) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUser#7b8e7de6 as nil")
	}
	b.PutID(InputPeerUserTypeID)
	b.PutInt(i.UserID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetUserID returns value of UserID field.
func (i *InputPeerUser) GetUserID() (value int) {
	return i.UserID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPeerUser) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputPeerUser) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUser#7b8e7de6 to nil")
	}
	if err := b.ConsumeID(InputPeerUserTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerUser#7b8e7de6: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUser#7b8e7de6: field user_id: %w", err)
		}
		i.UserID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUser#7b8e7de6: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerUser) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerUser.
var (
	_ bin.Encoder = &InputPeerUser{}
	_ bin.Decoder = &InputPeerUser{}

	_ InputPeerClass = &InputPeerUser{}
)

// InputPeerChannel represents TL type `inputPeerChannel#20adaef8`.
// Defines a channel for further interaction.
//
// See https://core.telegram.org/constructor/inputPeerChannel for reference.
type InputPeerChannel struct {
	// Channel identifier
	ChannelID int `schemaname:"channel_id"`
	// access_hash value from the channel¹ constructor
	//
	// Links:
	//  1) https://core.telegram.org/constructor/channel
	AccessHash int64 `schemaname:"access_hash"`
}

// InputPeerChannelTypeID is TL type id of InputPeerChannel.
const InputPeerChannelTypeID = 0x20adaef8

func (i *InputPeerChannel) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChannelID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChannel) String() string {
	if i == nil {
		return "InputPeerChannel(nil)"
	}
	type Alias InputPeerChannel
	return fmt.Sprintf("InputPeerChannel%+v", Alias(*i))
}

// FillFrom fills InputPeerChannel from given interface.
func (i *InputPeerChannel) FillFrom(from interface {
	GetChannelID() (value int)
	GetAccessHash() (value int64)
}) {
	i.ChannelID = from.GetChannelID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerChannel) TypeID() uint32 {
	return InputPeerChannelTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerChannel) SchemaName() string {
	return "inputPeerChannel"
}

// Encode implements bin.Encoder.
func (i *InputPeerChannel) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannel#20adaef8 as nil")
	}
	b.PutID(InputPeerChannelTypeID)
	b.PutInt(i.ChannelID)
	b.PutLong(i.AccessHash)
	return nil
}

// GetChannelID returns value of ChannelID field.
func (i *InputPeerChannel) GetChannelID() (value int) {
	return i.ChannelID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputPeerChannel) GetAccessHash() (value int64) {
	return i.AccessHash
}

// Decode implements bin.Decoder.
func (i *InputPeerChannel) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannel#20adaef8 to nil")
	}
	if err := b.ConsumeID(InputPeerChannelTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChannel#20adaef8: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannel#20adaef8: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannel#20adaef8: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerChannel) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChannel.
var (
	_ bin.Encoder = &InputPeerChannel{}
	_ bin.Decoder = &InputPeerChannel{}

	_ InputPeerClass = &InputPeerChannel{}
)

// InputPeerUserFromMessage represents TL type `inputPeerUserFromMessage#17bae2e6`.
// Defines a min¹ user that was seen in a certain message of a certain chat.
//
// Links:
//  1) https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputPeerUserFromMessage for reference.
type InputPeerUserFromMessage struct {
	// The chat where the user was seen
	Peer InputPeerClass `schemaname:"peer"`
	// The message ID
	MsgID int `schemaname:"msg_id"`
	// The identifier of the user that was seen
	UserID int `schemaname:"user_id"`
}

// InputPeerUserFromMessageTypeID is TL type id of InputPeerUserFromMessage.
const InputPeerUserFromMessageTypeID = 0x17bae2e6

func (i *InputPeerUserFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerUserFromMessage) String() string {
	if i == nil {
		return "InputPeerUserFromMessage(nil)"
	}
	type Alias InputPeerUserFromMessage
	return fmt.Sprintf("InputPeerUserFromMessage%+v", Alias(*i))
}

// FillFrom fills InputPeerUserFromMessage from given interface.
func (i *InputPeerUserFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetUserID() (value int)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.UserID = from.GetUserID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerUserFromMessage) TypeID() uint32 {
	return InputPeerUserFromMessageTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerUserFromMessage) SchemaName() string {
	return "inputPeerUserFromMessage"
}

// Encode implements bin.Encoder.
func (i *InputPeerUserFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerUserFromMessage#17bae2e6 as nil")
	}
	b.PutID(InputPeerUserFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputPeerUserFromMessage#17bae2e6: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerUserFromMessage#17bae2e6: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.UserID)
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputPeerUserFromMessage) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputPeerUserFromMessage) GetMsgID() (value int) {
	return i.MsgID
}

// GetUserID returns value of UserID field.
func (i *InputPeerUserFromMessage) GetUserID() (value int) {
	return i.UserID
}

// Decode implements bin.Decoder.
func (i *InputPeerUserFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerUserFromMessage#17bae2e6 to nil")
	}
	if err := b.ConsumeID(InputPeerUserFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerUserFromMessage#17bae2e6: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#17bae2e6: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#17bae2e6: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerUserFromMessage#17bae2e6: field user_id: %w", err)
		}
		i.UserID = value
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerUserFromMessage) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerUserFromMessage.
var (
	_ bin.Encoder = &InputPeerUserFromMessage{}
	_ bin.Decoder = &InputPeerUserFromMessage{}

	_ InputPeerClass = &InputPeerUserFromMessage{}
)

// InputPeerChannelFromMessage represents TL type `inputPeerChannelFromMessage#9c95f7bb`.
// Defines a min¹ channel that was seen in a certain message of a certain chat.
//
// Links:
//  1) https://core.telegram.org/api/min
//
// See https://core.telegram.org/constructor/inputPeerChannelFromMessage for reference.
type InputPeerChannelFromMessage struct {
	// The chat where the channel's message was seen
	Peer InputPeerClass `schemaname:"peer"`
	// The message ID
	MsgID int `schemaname:"msg_id"`
	// The identifier of the channel that was seen
	ChannelID int `schemaname:"channel_id"`
}

// InputPeerChannelFromMessageTypeID is TL type id of InputPeerChannelFromMessage.
const InputPeerChannelFromMessageTypeID = 0x9c95f7bb

func (i *InputPeerChannelFromMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.MsgID == 0) {
		return false
	}
	if !(i.ChannelID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPeerChannelFromMessage) String() string {
	if i == nil {
		return "InputPeerChannelFromMessage(nil)"
	}
	type Alias InputPeerChannelFromMessage
	return fmt.Sprintf("InputPeerChannelFromMessage%+v", Alias(*i))
}

// FillFrom fills InputPeerChannelFromMessage from given interface.
func (i *InputPeerChannelFromMessage) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMsgID() (value int)
	GetChannelID() (value int)
}) {
	i.Peer = from.GetPeer()
	i.MsgID = from.GetMsgID()
	i.ChannelID = from.GetChannelID()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPeerChannelFromMessage) TypeID() uint32 {
	return InputPeerChannelFromMessageTypeID
}

// SchemaName returns MTProto type name.
func (i *InputPeerChannelFromMessage) SchemaName() string {
	return "inputPeerChannelFromMessage"
}

// Encode implements bin.Encoder.
func (i *InputPeerChannelFromMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPeerChannelFromMessage#9c95f7bb as nil")
	}
	b.PutID(InputPeerChannelFromMessageTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputPeerChannelFromMessage#9c95f7bb: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPeerChannelFromMessage#9c95f7bb: field peer: %w", err)
	}
	b.PutInt(i.MsgID)
	b.PutInt(i.ChannelID)
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputPeerChannelFromMessage) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetMsgID returns value of MsgID field.
func (i *InputPeerChannelFromMessage) GetMsgID() (value int) {
	return i.MsgID
}

// GetChannelID returns value of ChannelID field.
func (i *InputPeerChannelFromMessage) GetChannelID() (value int) {
	return i.ChannelID
}

// Decode implements bin.Decoder.
func (i *InputPeerChannelFromMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPeerChannelFromMessage#9c95f7bb to nil")
	}
	if err := b.ConsumeID(InputPeerChannelFromMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPeerChannelFromMessage#9c95f7bb: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#9c95f7bb: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#9c95f7bb: field msg_id: %w", err)
		}
		i.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputPeerChannelFromMessage#9c95f7bb: field channel_id: %w", err)
		}
		i.ChannelID = value
	}
	return nil
}

// construct implements constructor of InputPeerClass.
func (i InputPeerChannelFromMessage) construct() InputPeerClass { return &i }

// Ensuring interfaces in compile-time for InputPeerChannelFromMessage.
var (
	_ bin.Encoder = &InputPeerChannelFromMessage{}
	_ bin.Decoder = &InputPeerChannelFromMessage{}

	_ InputPeerClass = &InputPeerChannelFromMessage{}
)

// InputPeerClass represents InputPeer generic type.
//
// See https://core.telegram.org/type/InputPeer for reference.
//
// Example:
//  g, err := tg.DecodeInputPeer(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputPeerEmpty: // inputPeerEmpty#7f3b18ea
//  case *tg.InputPeerSelf: // inputPeerSelf#7da07ec9
//  case *tg.InputPeerChat: // inputPeerChat#179be863
//  case *tg.InputPeerUser: // inputPeerUser#7b8e7de6
//  case *tg.InputPeerChannel: // inputPeerChannel#20adaef8
//  case *tg.InputPeerUserFromMessage: // inputPeerUserFromMessage#17bae2e6
//  case *tg.InputPeerChannelFromMessage: // inputPeerChannelFromMessage#9c95f7bb
//  default: panic(v)
//  }
type InputPeerClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPeerClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// SchemaName returns MTProto type name.
	SchemaName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInputPeer implements binary de-serialization for InputPeerClass.
func DecodeInputPeer(buf *bin.Buffer) (InputPeerClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPeerEmptyTypeID:
		// Decoding inputPeerEmpty#7f3b18ea.
		v := InputPeerEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerSelfTypeID:
		// Decoding inputPeerSelf#7da07ec9.
		v := InputPeerSelf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChatTypeID:
		// Decoding inputPeerChat#179be863.
		v := InputPeerChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerUserTypeID:
		// Decoding inputPeerUser#7b8e7de6.
		v := InputPeerUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChannelTypeID:
		// Decoding inputPeerChannel#20adaef8.
		v := InputPeerChannel{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerUserFromMessageTypeID:
		// Decoding inputPeerUserFromMessage#17bae2e6.
		v := InputPeerUserFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	case InputPeerChannelFromMessageTypeID:
		// Decoding inputPeerChannelFromMessage#9c95f7bb.
		v := InputPeerChannelFromMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPeerClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPeerClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPeer boxes the InputPeerClass providing a helper.
type InputPeerBox struct {
	InputPeer InputPeerClass
}

// Decode implements bin.Decoder for InputPeerBox.
func (b *InputPeerBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPeerBox to nil")
	}
	v, err := DecodeInputPeer(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPeer = v
	return nil
}

// Encode implements bin.Encode for InputPeerBox.
func (b *InputPeerBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPeer == nil {
		return fmt.Errorf("unable to encode InputPeerClass as nil")
	}
	return b.InputPeer.Encode(buf)
}

// InputPeerClassSlice is adapter for slice of InputPeerClass.
type InputPeerClassSlice []InputPeerClass

// First returns first element of slice (if exists).
func (s InputPeerClassSlice) First() (v InputPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputPeerClassSlice) Last() (v InputPeerClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputPeerClassSlice) PopFirst() (v InputPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	a[len(a)-1] = nil
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputPeerClassSlice) Pop() (v InputPeerClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
