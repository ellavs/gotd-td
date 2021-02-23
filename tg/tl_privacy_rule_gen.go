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

// PrivacyValueAllowContacts represents TL type `privacyValueAllowContacts#fffe1bac`.
// Allow all contacts
//
// See https://core.telegram.org/constructor/privacyValueAllowContacts for reference.
type PrivacyValueAllowContacts struct {
}

// PrivacyValueAllowContactsTypeID is TL type id of PrivacyValueAllowContacts.
const PrivacyValueAllowContactsTypeID = 0xfffe1bac

func (p *PrivacyValueAllowContacts) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueAllowContacts) String() string {
	if p == nil {
		return "PrivacyValueAllowContacts(nil)"
	}
	type Alias PrivacyValueAllowContacts
	return fmt.Sprintf("PrivacyValueAllowContacts%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueAllowContacts) TypeID() uint32 {
	return PrivacyValueAllowContactsTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueAllowContacts) SchemaName() string {
	return "privacyValueAllowContacts"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueAllowContacts) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueAllowContacts#fffe1bac as nil")
	}
	b.PutID(PrivacyValueAllowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyValueAllowContacts) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueAllowContacts#fffe1bac to nil")
	}
	if err := b.ConsumeID(PrivacyValueAllowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueAllowContacts#fffe1bac: %w", err)
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueAllowContacts) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueAllowContacts.
var (
	_ bin.Encoder = &PrivacyValueAllowContacts{}
	_ bin.Decoder = &PrivacyValueAllowContacts{}

	_ PrivacyRuleClass = &PrivacyValueAllowContacts{}
)

// PrivacyValueAllowAll represents TL type `privacyValueAllowAll#65427b82`.
// Allow all users
//
// See https://core.telegram.org/constructor/privacyValueAllowAll for reference.
type PrivacyValueAllowAll struct {
}

// PrivacyValueAllowAllTypeID is TL type id of PrivacyValueAllowAll.
const PrivacyValueAllowAllTypeID = 0x65427b82

func (p *PrivacyValueAllowAll) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueAllowAll) String() string {
	if p == nil {
		return "PrivacyValueAllowAll(nil)"
	}
	type Alias PrivacyValueAllowAll
	return fmt.Sprintf("PrivacyValueAllowAll%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueAllowAll) TypeID() uint32 {
	return PrivacyValueAllowAllTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueAllowAll) SchemaName() string {
	return "privacyValueAllowAll"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueAllowAll) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueAllowAll#65427b82 as nil")
	}
	b.PutID(PrivacyValueAllowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyValueAllowAll) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueAllowAll#65427b82 to nil")
	}
	if err := b.ConsumeID(PrivacyValueAllowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueAllowAll#65427b82: %w", err)
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueAllowAll) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueAllowAll.
var (
	_ bin.Encoder = &PrivacyValueAllowAll{}
	_ bin.Decoder = &PrivacyValueAllowAll{}

	_ PrivacyRuleClass = &PrivacyValueAllowAll{}
)

// PrivacyValueAllowUsers represents TL type `privacyValueAllowUsers#4d5bbe0c`.
// Allow only certain users
//
// See https://core.telegram.org/constructor/privacyValueAllowUsers for reference.
type PrivacyValueAllowUsers struct {
	// Allowed users
	Users []int `schemaname:"users"`
}

// PrivacyValueAllowUsersTypeID is TL type id of PrivacyValueAllowUsers.
const PrivacyValueAllowUsersTypeID = 0x4d5bbe0c

func (p *PrivacyValueAllowUsers) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueAllowUsers) String() string {
	if p == nil {
		return "PrivacyValueAllowUsers(nil)"
	}
	type Alias PrivacyValueAllowUsers
	return fmt.Sprintf("PrivacyValueAllowUsers%+v", Alias(*p))
}

// FillFrom fills PrivacyValueAllowUsers from given interface.
func (p *PrivacyValueAllowUsers) FillFrom(from interface {
	GetUsers() (value []int)
}) {
	p.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueAllowUsers) TypeID() uint32 {
	return PrivacyValueAllowUsersTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueAllowUsers) SchemaName() string {
	return "privacyValueAllowUsers"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueAllowUsers) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueAllowUsers#4d5bbe0c as nil")
	}
	b.PutID(PrivacyValueAllowUsersTypeID)
	b.PutVectorHeader(len(p.Users))
	for _, v := range p.Users {
		b.PutInt(v)
	}
	return nil
}

// GetUsers returns value of Users field.
func (p *PrivacyValueAllowUsers) GetUsers() (value []int) {
	return p.Users
}

// Decode implements bin.Decoder.
func (p *PrivacyValueAllowUsers) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueAllowUsers#4d5bbe0c to nil")
	}
	if err := b.ConsumeID(PrivacyValueAllowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueAllowUsers#4d5bbe0c: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode privacyValueAllowUsers#4d5bbe0c: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode privacyValueAllowUsers#4d5bbe0c: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueAllowUsers) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueAllowUsers.
var (
	_ bin.Encoder = &PrivacyValueAllowUsers{}
	_ bin.Decoder = &PrivacyValueAllowUsers{}

	_ PrivacyRuleClass = &PrivacyValueAllowUsers{}
)

// PrivacyValueDisallowContacts represents TL type `privacyValueDisallowContacts#f888fa1a`.
// Disallow only contacts
//
// See https://core.telegram.org/constructor/privacyValueDisallowContacts for reference.
type PrivacyValueDisallowContacts struct {
}

// PrivacyValueDisallowContactsTypeID is TL type id of PrivacyValueDisallowContacts.
const PrivacyValueDisallowContactsTypeID = 0xf888fa1a

func (p *PrivacyValueDisallowContacts) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueDisallowContacts) String() string {
	if p == nil {
		return "PrivacyValueDisallowContacts(nil)"
	}
	type Alias PrivacyValueDisallowContacts
	return fmt.Sprintf("PrivacyValueDisallowContacts%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueDisallowContacts) TypeID() uint32 {
	return PrivacyValueDisallowContactsTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueDisallowContacts) SchemaName() string {
	return "privacyValueDisallowContacts"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueDisallowContacts) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueDisallowContacts#f888fa1a as nil")
	}
	b.PutID(PrivacyValueDisallowContactsTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyValueDisallowContacts) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueDisallowContacts#f888fa1a to nil")
	}
	if err := b.ConsumeID(PrivacyValueDisallowContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueDisallowContacts#f888fa1a: %w", err)
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueDisallowContacts) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueDisallowContacts.
var (
	_ bin.Encoder = &PrivacyValueDisallowContacts{}
	_ bin.Decoder = &PrivacyValueDisallowContacts{}

	_ PrivacyRuleClass = &PrivacyValueDisallowContacts{}
)

// PrivacyValueDisallowAll represents TL type `privacyValueDisallowAll#8b73e763`.
// Disallow all users
//
// See https://core.telegram.org/constructor/privacyValueDisallowAll for reference.
type PrivacyValueDisallowAll struct {
}

// PrivacyValueDisallowAllTypeID is TL type id of PrivacyValueDisallowAll.
const PrivacyValueDisallowAllTypeID = 0x8b73e763

func (p *PrivacyValueDisallowAll) Zero() bool {
	if p == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueDisallowAll) String() string {
	if p == nil {
		return "PrivacyValueDisallowAll(nil)"
	}
	type Alias PrivacyValueDisallowAll
	return fmt.Sprintf("PrivacyValueDisallowAll%+v", Alias(*p))
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueDisallowAll) TypeID() uint32 {
	return PrivacyValueDisallowAllTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueDisallowAll) SchemaName() string {
	return "privacyValueDisallowAll"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueDisallowAll) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueDisallowAll#8b73e763 as nil")
	}
	b.PutID(PrivacyValueDisallowAllTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PrivacyValueDisallowAll) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueDisallowAll#8b73e763 to nil")
	}
	if err := b.ConsumeID(PrivacyValueDisallowAllTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueDisallowAll#8b73e763: %w", err)
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueDisallowAll) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueDisallowAll.
var (
	_ bin.Encoder = &PrivacyValueDisallowAll{}
	_ bin.Decoder = &PrivacyValueDisallowAll{}

	_ PrivacyRuleClass = &PrivacyValueDisallowAll{}
)

// PrivacyValueDisallowUsers represents TL type `privacyValueDisallowUsers#c7f49b7`.
// Disallow only certain users
//
// See https://core.telegram.org/constructor/privacyValueDisallowUsers for reference.
type PrivacyValueDisallowUsers struct {
	// Disallowed users
	Users []int `schemaname:"users"`
}

// PrivacyValueDisallowUsersTypeID is TL type id of PrivacyValueDisallowUsers.
const PrivacyValueDisallowUsersTypeID = 0xc7f49b7

func (p *PrivacyValueDisallowUsers) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueDisallowUsers) String() string {
	if p == nil {
		return "PrivacyValueDisallowUsers(nil)"
	}
	type Alias PrivacyValueDisallowUsers
	return fmt.Sprintf("PrivacyValueDisallowUsers%+v", Alias(*p))
}

// FillFrom fills PrivacyValueDisallowUsers from given interface.
func (p *PrivacyValueDisallowUsers) FillFrom(from interface {
	GetUsers() (value []int)
}) {
	p.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueDisallowUsers) TypeID() uint32 {
	return PrivacyValueDisallowUsersTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueDisallowUsers) SchemaName() string {
	return "privacyValueDisallowUsers"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueDisallowUsers) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueDisallowUsers#c7f49b7 as nil")
	}
	b.PutID(PrivacyValueDisallowUsersTypeID)
	b.PutVectorHeader(len(p.Users))
	for _, v := range p.Users {
		b.PutInt(v)
	}
	return nil
}

// GetUsers returns value of Users field.
func (p *PrivacyValueDisallowUsers) GetUsers() (value []int) {
	return p.Users
}

// Decode implements bin.Decoder.
func (p *PrivacyValueDisallowUsers) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueDisallowUsers#c7f49b7 to nil")
	}
	if err := b.ConsumeID(PrivacyValueDisallowUsersTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueDisallowUsers#c7f49b7: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode privacyValueDisallowUsers#c7f49b7: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode privacyValueDisallowUsers#c7f49b7: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueDisallowUsers) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueDisallowUsers.
var (
	_ bin.Encoder = &PrivacyValueDisallowUsers{}
	_ bin.Decoder = &PrivacyValueDisallowUsers{}

	_ PrivacyRuleClass = &PrivacyValueDisallowUsers{}
)

// PrivacyValueAllowChatParticipants represents TL type `privacyValueAllowChatParticipants#18be796b`.
// Allow all participants of certain chats
//
// See https://core.telegram.org/constructor/privacyValueAllowChatParticipants for reference.
type PrivacyValueAllowChatParticipants struct {
	// Allowed chats
	Chats []int `schemaname:"chats"`
}

// PrivacyValueAllowChatParticipantsTypeID is TL type id of PrivacyValueAllowChatParticipants.
const PrivacyValueAllowChatParticipantsTypeID = 0x18be796b

func (p *PrivacyValueAllowChatParticipants) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueAllowChatParticipants) String() string {
	if p == nil {
		return "PrivacyValueAllowChatParticipants(nil)"
	}
	type Alias PrivacyValueAllowChatParticipants
	return fmt.Sprintf("PrivacyValueAllowChatParticipants%+v", Alias(*p))
}

// FillFrom fills PrivacyValueAllowChatParticipants from given interface.
func (p *PrivacyValueAllowChatParticipants) FillFrom(from interface {
	GetChats() (value []int)
}) {
	p.Chats = from.GetChats()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueAllowChatParticipants) TypeID() uint32 {
	return PrivacyValueAllowChatParticipantsTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueAllowChatParticipants) SchemaName() string {
	return "privacyValueAllowChatParticipants"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueAllowChatParticipants) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueAllowChatParticipants#18be796b as nil")
	}
	b.PutID(PrivacyValueAllowChatParticipantsTypeID)
	b.PutVectorHeader(len(p.Chats))
	for _, v := range p.Chats {
		b.PutInt(v)
	}
	return nil
}

// GetChats returns value of Chats field.
func (p *PrivacyValueAllowChatParticipants) GetChats() (value []int) {
	return p.Chats
}

// Decode implements bin.Decoder.
func (p *PrivacyValueAllowChatParticipants) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueAllowChatParticipants#18be796b to nil")
	}
	if err := b.ConsumeID(PrivacyValueAllowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueAllowChatParticipants#18be796b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode privacyValueAllowChatParticipants#18be796b: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode privacyValueAllowChatParticipants#18be796b: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueAllowChatParticipants) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueAllowChatParticipants.
var (
	_ bin.Encoder = &PrivacyValueAllowChatParticipants{}
	_ bin.Decoder = &PrivacyValueAllowChatParticipants{}

	_ PrivacyRuleClass = &PrivacyValueAllowChatParticipants{}
)

// PrivacyValueDisallowChatParticipants represents TL type `privacyValueDisallowChatParticipants#acae0690`.
// Disallow only participants of certain chats
//
// See https://core.telegram.org/constructor/privacyValueDisallowChatParticipants for reference.
type PrivacyValueDisallowChatParticipants struct {
	// Disallowed chats
	Chats []int `schemaname:"chats"`
}

// PrivacyValueDisallowChatParticipantsTypeID is TL type id of PrivacyValueDisallowChatParticipants.
const PrivacyValueDisallowChatParticipantsTypeID = 0xacae0690

func (p *PrivacyValueDisallowChatParticipants) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Chats == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PrivacyValueDisallowChatParticipants) String() string {
	if p == nil {
		return "PrivacyValueDisallowChatParticipants(nil)"
	}
	type Alias PrivacyValueDisallowChatParticipants
	return fmt.Sprintf("PrivacyValueDisallowChatParticipants%+v", Alias(*p))
}

// FillFrom fills PrivacyValueDisallowChatParticipants from given interface.
func (p *PrivacyValueDisallowChatParticipants) FillFrom(from interface {
	GetChats() (value []int)
}) {
	p.Chats = from.GetChats()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PrivacyValueDisallowChatParticipants) TypeID() uint32 {
	return PrivacyValueDisallowChatParticipantsTypeID
}

// SchemaName returns MTProto type name.
func (p *PrivacyValueDisallowChatParticipants) SchemaName() string {
	return "privacyValueDisallowChatParticipants"
}

// Encode implements bin.Encoder.
func (p *PrivacyValueDisallowChatParticipants) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode privacyValueDisallowChatParticipants#acae0690 as nil")
	}
	b.PutID(PrivacyValueDisallowChatParticipantsTypeID)
	b.PutVectorHeader(len(p.Chats))
	for _, v := range p.Chats {
		b.PutInt(v)
	}
	return nil
}

// GetChats returns value of Chats field.
func (p *PrivacyValueDisallowChatParticipants) GetChats() (value []int) {
	return p.Chats
}

// Decode implements bin.Decoder.
func (p *PrivacyValueDisallowChatParticipants) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode privacyValueDisallowChatParticipants#acae0690 to nil")
	}
	if err := b.ConsumeID(PrivacyValueDisallowChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode privacyValueDisallowChatParticipants#acae0690: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode privacyValueDisallowChatParticipants#acae0690: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode privacyValueDisallowChatParticipants#acae0690: field chats: %w", err)
			}
			p.Chats = append(p.Chats, value)
		}
	}
	return nil
}

// construct implements constructor of PrivacyRuleClass.
func (p PrivacyValueDisallowChatParticipants) construct() PrivacyRuleClass { return &p }

// Ensuring interfaces in compile-time for PrivacyValueDisallowChatParticipants.
var (
	_ bin.Encoder = &PrivacyValueDisallowChatParticipants{}
	_ bin.Decoder = &PrivacyValueDisallowChatParticipants{}

	_ PrivacyRuleClass = &PrivacyValueDisallowChatParticipants{}
)

// PrivacyRuleClass represents PrivacyRule generic type.
//
// See https://core.telegram.org/type/PrivacyRule for reference.
//
// Example:
//  g, err := tg.DecodePrivacyRule(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PrivacyValueAllowContacts: // privacyValueAllowContacts#fffe1bac
//  case *tg.PrivacyValueAllowAll: // privacyValueAllowAll#65427b82
//  case *tg.PrivacyValueAllowUsers: // privacyValueAllowUsers#4d5bbe0c
//  case *tg.PrivacyValueDisallowContacts: // privacyValueDisallowContacts#f888fa1a
//  case *tg.PrivacyValueDisallowAll: // privacyValueDisallowAll#8b73e763
//  case *tg.PrivacyValueDisallowUsers: // privacyValueDisallowUsers#c7f49b7
//  case *tg.PrivacyValueAllowChatParticipants: // privacyValueAllowChatParticipants#18be796b
//  case *tg.PrivacyValueDisallowChatParticipants: // privacyValueDisallowChatParticipants#acae0690
//  default: panic(v)
//  }
type PrivacyRuleClass interface {
	bin.Encoder
	bin.Decoder
	construct() PrivacyRuleClass

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

// AsInput tries to map PrivacyValueAllowChatParticipants to InputPrivacyValueAllowChatParticipants.
func (p *PrivacyValueAllowChatParticipants) AsInput() *InputPrivacyValueAllowChatParticipants {
	value := new(InputPrivacyValueAllowChatParticipants)
	value.Chats = p.GetChats()

	return value
}

// AsInput tries to map PrivacyValueDisallowChatParticipants to InputPrivacyValueDisallowChatParticipants.
func (p *PrivacyValueDisallowChatParticipants) AsInput() *InputPrivacyValueDisallowChatParticipants {
	value := new(InputPrivacyValueDisallowChatParticipants)
	value.Chats = p.GetChats()

	return value
}

// DecodePrivacyRule implements binary de-serialization for PrivacyRuleClass.
func DecodePrivacyRule(buf *bin.Buffer) (PrivacyRuleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PrivacyValueAllowContactsTypeID:
		// Decoding privacyValueAllowContacts#fffe1bac.
		v := PrivacyValueAllowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueAllowAllTypeID:
		// Decoding privacyValueAllowAll#65427b82.
		v := PrivacyValueAllowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueAllowUsersTypeID:
		// Decoding privacyValueAllowUsers#4d5bbe0c.
		v := PrivacyValueAllowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueDisallowContactsTypeID:
		// Decoding privacyValueDisallowContacts#f888fa1a.
		v := PrivacyValueDisallowContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueDisallowAllTypeID:
		// Decoding privacyValueDisallowAll#8b73e763.
		v := PrivacyValueDisallowAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueDisallowUsersTypeID:
		// Decoding privacyValueDisallowUsers#c7f49b7.
		v := PrivacyValueDisallowUsers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueAllowChatParticipantsTypeID:
		// Decoding privacyValueAllowChatParticipants#18be796b.
		v := PrivacyValueAllowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	case PrivacyValueDisallowChatParticipantsTypeID:
		// Decoding privacyValueDisallowChatParticipants#acae0690.
		v := PrivacyValueDisallowChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PrivacyRuleClass: %w", bin.NewUnexpectedID(id))
	}
}

// PrivacyRule boxes the PrivacyRuleClass providing a helper.
type PrivacyRuleBox struct {
	PrivacyRule PrivacyRuleClass
}

// Decode implements bin.Decoder for PrivacyRuleBox.
func (b *PrivacyRuleBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PrivacyRuleBox to nil")
	}
	v, err := DecodePrivacyRule(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PrivacyRule = v
	return nil
}

// Encode implements bin.Encode for PrivacyRuleBox.
func (b *PrivacyRuleBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PrivacyRule == nil {
		return fmt.Errorf("unable to encode PrivacyRuleClass as nil")
	}
	return b.PrivacyRule.Encode(buf)
}

// PrivacyRuleClassSlice is adapter for slice of PrivacyRuleClass.
type PrivacyRuleClassSlice []PrivacyRuleClass

// First returns first element of slice (if exists).
func (s PrivacyRuleClassSlice) First() (v PrivacyRuleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PrivacyRuleClassSlice) Last() (v PrivacyRuleClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PrivacyRuleClassSlice) PopFirst() (v PrivacyRuleClass, ok bool) {
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
func (s *PrivacyRuleClassSlice) Pop() (v PrivacyRuleClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
