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

// ContactsFound represents TL type `contacts.found#b3134d9d`.
// Users found by name substring and auxiliary data.
//
// See https://core.telegram.org/constructor/contacts.found for reference.
type ContactsFound struct {
	// Personalized results
	MyResults []PeerClass `schemaname:"my_results"`
	// List of found user identifiers
	Results []PeerClass `schemaname:"results"`
	// Found chats
	Chats []ChatClass `schemaname:"chats"`
	// List of users
	Users []UserClass `schemaname:"users"`
}

// ContactsFoundTypeID is TL type id of ContactsFound.
const ContactsFoundTypeID = 0xb3134d9d

func (f *ContactsFound) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.MyResults == nil) {
		return false
	}
	if !(f.Results == nil) {
		return false
	}
	if !(f.Chats == nil) {
		return false
	}
	if !(f.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *ContactsFound) String() string {
	if f == nil {
		return "ContactsFound(nil)"
	}
	type Alias ContactsFound
	return fmt.Sprintf("ContactsFound%+v", Alias(*f))
}

// FillFrom fills ContactsFound from given interface.
func (f *ContactsFound) FillFrom(from interface {
	GetMyResults() (value []PeerClass)
	GetResults() (value []PeerClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	f.MyResults = from.GetMyResults()
	f.Results = from.GetResults()
	f.Chats = from.GetChats()
	f.Users = from.GetUsers()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *ContactsFound) TypeID() uint32 {
	return ContactsFoundTypeID
}

// SchemaName returns MTProto type name.
func (f *ContactsFound) SchemaName() string {
	return "contacts.found"
}

// Encode implements bin.Encoder.
func (f *ContactsFound) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode contacts.found#b3134d9d as nil")
	}
	b.PutID(ContactsFoundTypeID)
	b.PutVectorHeader(len(f.MyResults))
	for idx, v := range f.MyResults {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field my_results element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field my_results element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Results))
	for idx, v := range f.Results {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field results element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field results element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Chats))
	for idx, v := range f.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(f.Users))
	for idx, v := range f.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.found#b3134d9d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetMyResults returns value of MyResults field.
func (f *ContactsFound) GetMyResults() (value []PeerClass) {
	return f.MyResults
}

// MapMyResults returns field MyResults wrapped in PeerClassSlice helper.
func (f *ContactsFound) MapMyResults() (value PeerClassSlice) {
	return PeerClassSlice(f.MyResults)
}

// GetResults returns value of Results field.
func (f *ContactsFound) GetResults() (value []PeerClass) {
	return f.Results
}

// MapResults returns field Results wrapped in PeerClassSlice helper.
func (f *ContactsFound) MapResults() (value PeerClassSlice) {
	return PeerClassSlice(f.Results)
}

// GetChats returns value of Chats field.
func (f *ContactsFound) GetChats() (value []ChatClass) {
	return f.Chats
}

// MapChats returns field Chats wrapped in ChatClassSlice helper.
func (f *ContactsFound) MapChats() (value ChatClassSlice) {
	return ChatClassSlice(f.Chats)
}

// GetUsers returns value of Users field.
func (f *ContactsFound) GetUsers() (value []UserClass) {
	return f.Users
}

// MapUsers returns field Users wrapped in UserClassSlice helper.
func (f *ContactsFound) MapUsers() (value UserClassSlice) {
	return UserClassSlice(f.Users)
}

// Decode implements bin.Decoder.
func (f *ContactsFound) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode contacts.found#b3134d9d to nil")
	}
	if err := b.ConsumeID(ContactsFoundTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.found#b3134d9d: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field my_results: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field my_results: %w", err)
			}
			f.MyResults = append(f.MyResults, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field results: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field results: %w", err)
			}
			f.Results = append(f.Results, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field chats: %w", err)
			}
			f.Chats = append(f.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.found#b3134d9d: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.found#b3134d9d: field users: %w", err)
			}
			f.Users = append(f.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsFound.
var (
	_ bin.Encoder = &ContactsFound{}
	_ bin.Decoder = &ContactsFound{}
)
