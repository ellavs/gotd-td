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

// GlobalPrivacySettings represents TL type `globalPrivacySettings#734c4ccb`.
// Global privacy settings
//
// See https://core.telegram.org/constructor/globalPrivacySettings for reference.
type GlobalPrivacySettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to archive and mute new chats from non-contacts
	ArchiveAndMuteNewNoncontactPeers bool
	// Whether unmuted chats will be kept in the Archive chat list when they get a new
	// message.
	KeepArchivedUnmuted bool
	// Whether unmuted chats that are always included or pinned in a folder¹, will be kept
	// in the Archive chat list when they get a new message. Ignored if keep_archived_unmuted
	// is set.
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	KeepArchivedFolders bool
	// HideReadMarks field of GlobalPrivacySettings.
	HideReadMarks bool
	// NewNoncontactPeersRequirePremium field of GlobalPrivacySettings.
	NewNoncontactPeersRequirePremium bool
}

// GlobalPrivacySettingsTypeID is TL type id of GlobalPrivacySettings.
const GlobalPrivacySettingsTypeID = 0x734c4ccb

// Ensuring interfaces in compile-time for GlobalPrivacySettings.
var (
	_ bin.Encoder     = &GlobalPrivacySettings{}
	_ bin.Decoder     = &GlobalPrivacySettings{}
	_ bin.BareEncoder = &GlobalPrivacySettings{}
	_ bin.BareDecoder = &GlobalPrivacySettings{}
)

func (g *GlobalPrivacySettings) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		return false
	}
	if !(g.KeepArchivedUnmuted == false) {
		return false
	}
	if !(g.KeepArchivedFolders == false) {
		return false
	}
	if !(g.HideReadMarks == false) {
		return false
	}
	if !(g.NewNoncontactPeersRequirePremium == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GlobalPrivacySettings) String() string {
	if g == nil {
		return "GlobalPrivacySettings(nil)"
	}
	type Alias GlobalPrivacySettings
	return fmt.Sprintf("GlobalPrivacySettings%+v", Alias(*g))
}

// FillFrom fills GlobalPrivacySettings from given interface.
func (g *GlobalPrivacySettings) FillFrom(from interface {
	GetArchiveAndMuteNewNoncontactPeers() (value bool)
	GetKeepArchivedUnmuted() (value bool)
	GetKeepArchivedFolders() (value bool)
	GetHideReadMarks() (value bool)
	GetNewNoncontactPeersRequirePremium() (value bool)
}) {
	g.ArchiveAndMuteNewNoncontactPeers = from.GetArchiveAndMuteNewNoncontactPeers()
	g.KeepArchivedUnmuted = from.GetKeepArchivedUnmuted()
	g.KeepArchivedFolders = from.GetKeepArchivedFolders()
	g.HideReadMarks = from.GetHideReadMarks()
	g.NewNoncontactPeersRequirePremium = from.GetNewNoncontactPeersRequirePremium()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GlobalPrivacySettings) TypeID() uint32 {
	return GlobalPrivacySettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*GlobalPrivacySettings) TypeName() string {
	return "globalPrivacySettings"
}

// TypeInfo returns info about TL type.
func (g *GlobalPrivacySettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "globalPrivacySettings",
		ID:   GlobalPrivacySettingsTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ArchiveAndMuteNewNoncontactPeers",
			SchemaName: "archive_and_mute_new_noncontact_peers",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "KeepArchivedUnmuted",
			SchemaName: "keep_archived_unmuted",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "KeepArchivedFolders",
			SchemaName: "keep_archived_folders",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "HideReadMarks",
			SchemaName: "hide_read_marks",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "NewNoncontactPeersRequirePremium",
			SchemaName: "new_noncontact_peers_require_premium",
			Null:       !g.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (g *GlobalPrivacySettings) SetFlags() {
	if !(g.ArchiveAndMuteNewNoncontactPeers == false) {
		g.Flags.Set(0)
	}
	if !(g.KeepArchivedUnmuted == false) {
		g.Flags.Set(1)
	}
	if !(g.KeepArchivedFolders == false) {
		g.Flags.Set(2)
	}
	if !(g.HideReadMarks == false) {
		g.Flags.Set(3)
	}
	if !(g.NewNoncontactPeersRequirePremium == false) {
		g.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (g *GlobalPrivacySettings) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#734c4ccb as nil")
	}
	b.PutID(GlobalPrivacySettingsTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GlobalPrivacySettings) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode globalPrivacySettings#734c4ccb as nil")
	}
	g.SetFlags()
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode globalPrivacySettings#734c4ccb: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GlobalPrivacySettings) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#734c4ccb to nil")
	}
	if err := b.ConsumeID(GlobalPrivacySettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode globalPrivacySettings#734c4ccb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GlobalPrivacySettings) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode globalPrivacySettings#734c4ccb to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode globalPrivacySettings#734c4ccb: field flags: %w", err)
		}
	}
	g.ArchiveAndMuteNewNoncontactPeers = g.Flags.Has(0)
	g.KeepArchivedUnmuted = g.Flags.Has(1)
	g.KeepArchivedFolders = g.Flags.Has(2)
	g.HideReadMarks = g.Flags.Has(3)
	g.NewNoncontactPeersRequirePremium = g.Flags.Has(4)
	return nil
}

// SetArchiveAndMuteNewNoncontactPeers sets value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) SetArchiveAndMuteNewNoncontactPeers(value bool) {
	if value {
		g.Flags.Set(0)
		g.ArchiveAndMuteNewNoncontactPeers = true
	} else {
		g.Flags.Unset(0)
		g.ArchiveAndMuteNewNoncontactPeers = false
	}
}

// GetArchiveAndMuteNewNoncontactPeers returns value of ArchiveAndMuteNewNoncontactPeers conditional field.
func (g *GlobalPrivacySettings) GetArchiveAndMuteNewNoncontactPeers() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(0)
}

// SetKeepArchivedUnmuted sets value of KeepArchivedUnmuted conditional field.
func (g *GlobalPrivacySettings) SetKeepArchivedUnmuted(value bool) {
	if value {
		g.Flags.Set(1)
		g.KeepArchivedUnmuted = true
	} else {
		g.Flags.Unset(1)
		g.KeepArchivedUnmuted = false
	}
}

// GetKeepArchivedUnmuted returns value of KeepArchivedUnmuted conditional field.
func (g *GlobalPrivacySettings) GetKeepArchivedUnmuted() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(1)
}

// SetKeepArchivedFolders sets value of KeepArchivedFolders conditional field.
func (g *GlobalPrivacySettings) SetKeepArchivedFolders(value bool) {
	if value {
		g.Flags.Set(2)
		g.KeepArchivedFolders = true
	} else {
		g.Flags.Unset(2)
		g.KeepArchivedFolders = false
	}
}

// GetKeepArchivedFolders returns value of KeepArchivedFolders conditional field.
func (g *GlobalPrivacySettings) GetKeepArchivedFolders() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(2)
}

// SetHideReadMarks sets value of HideReadMarks conditional field.
func (g *GlobalPrivacySettings) SetHideReadMarks(value bool) {
	if value {
		g.Flags.Set(3)
		g.HideReadMarks = true
	} else {
		g.Flags.Unset(3)
		g.HideReadMarks = false
	}
}

// GetHideReadMarks returns value of HideReadMarks conditional field.
func (g *GlobalPrivacySettings) GetHideReadMarks() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(3)
}

// SetNewNoncontactPeersRequirePremium sets value of NewNoncontactPeersRequirePremium conditional field.
func (g *GlobalPrivacySettings) SetNewNoncontactPeersRequirePremium(value bool) {
	if value {
		g.Flags.Set(4)
		g.NewNoncontactPeersRequirePremium = true
	} else {
		g.Flags.Unset(4)
		g.NewNoncontactPeersRequirePremium = false
	}
}

// GetNewNoncontactPeersRequirePremium returns value of NewNoncontactPeersRequirePremium conditional field.
func (g *GlobalPrivacySettings) GetNewNoncontactPeersRequirePremium() (value bool) {
	if g == nil {
		return
	}
	return g.Flags.Has(4)
}
