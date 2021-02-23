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

// UpdatesGetDifferenceRequest represents TL type `updates.getDifference#25939651`.
// Get new updates¹.
//
// Links:
//  1) https://core.telegram.org/api/updates
//
// See https://core.telegram.org/method/updates.getDifference for reference.
type UpdatesGetDifferenceRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields `schemaname:"flags"`
	// PTS, see updates¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int `schemaname:"pts"`
	// For fast updating: if provided and pts + pts_total_limit < remote pts, updates.differenceTooLong¹ will be returned.Simply tells the server to not return the difference if it is bigger than pts_total_limitIf the remote pts is too big (> ~4000000), this field will default to 1000000
	//
	// Links:
	//  1) https://core.telegram.org/constructor/updates.differenceTooLong
	//
	// Use SetPtsTotalLimit and GetPtsTotalLimit helpers.
	PtsTotalLimit int `schemaname:"pts_total_limit"`
	// date, see updates¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Date int `schemaname:"date"`
	// QTS, see updates¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Qts int `schemaname:"qts"`
}

// UpdatesGetDifferenceRequestTypeID is TL type id of UpdatesGetDifferenceRequest.
const UpdatesGetDifferenceRequestTypeID = 0x25939651

func (g *UpdatesGetDifferenceRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Pts == 0) {
		return false
	}
	if !(g.PtsTotalLimit == 0) {
		return false
	}
	if !(g.Date == 0) {
		return false
	}
	if !(g.Qts == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UpdatesGetDifferenceRequest) String() string {
	if g == nil {
		return "UpdatesGetDifferenceRequest(nil)"
	}
	type Alias UpdatesGetDifferenceRequest
	return fmt.Sprintf("UpdatesGetDifferenceRequest%+v", Alias(*g))
}

// FillFrom fills UpdatesGetDifferenceRequest from given interface.
func (g *UpdatesGetDifferenceRequest) FillFrom(from interface {
	GetPts() (value int)
	GetPtsTotalLimit() (value int, ok bool)
	GetDate() (value int)
	GetQts() (value int)
}) {
	g.Pts = from.GetPts()
	if val, ok := from.GetPtsTotalLimit(); ok {
		g.PtsTotalLimit = val
	}

	g.Date = from.GetDate()
	g.Qts = from.GetQts()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *UpdatesGetDifferenceRequest) TypeID() uint32 {
	return UpdatesGetDifferenceRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *UpdatesGetDifferenceRequest) SchemaName() string {
	return "updates.getDifference"
}

// Encode implements bin.Encoder.
func (g *UpdatesGetDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getDifference#25939651 as nil")
	}
	b.PutID(UpdatesGetDifferenceRequestTypeID)
	if !(g.PtsTotalLimit == 0) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getDifference#25939651: field flags: %w", err)
	}
	b.PutInt(g.Pts)
	if g.Flags.Has(0) {
		b.PutInt(g.PtsTotalLimit)
	}
	b.PutInt(g.Date)
	b.PutInt(g.Qts)
	return nil
}

// GetPts returns value of Pts field.
func (g *UpdatesGetDifferenceRequest) GetPts() (value int) {
	return g.Pts
}

// SetPtsTotalLimit sets value of PtsTotalLimit conditional field.
func (g *UpdatesGetDifferenceRequest) SetPtsTotalLimit(value int) {
	g.Flags.Set(0)
	g.PtsTotalLimit = value
}

// GetPtsTotalLimit returns value of PtsTotalLimit conditional field and
// boolean which is true if field was set.
func (g *UpdatesGetDifferenceRequest) GetPtsTotalLimit() (value int, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.PtsTotalLimit, true
}

// GetDate returns value of Date field.
func (g *UpdatesGetDifferenceRequest) GetDate() (value int) {
	return g.Date
}

// GetQts returns value of Qts field.
func (g *UpdatesGetDifferenceRequest) GetQts() (value int) {
	return g.Qts
}

// Decode implements bin.Decoder.
func (g *UpdatesGetDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getDifference#25939651 to nil")
	}
	if err := b.ConsumeID(UpdatesGetDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getDifference#25939651: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field pts: %w", err)
		}
		g.Pts = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field pts_total_limit: %w", err)
		}
		g.PtsTotalLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field date: %w", err)
		}
		g.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getDifference#25939651: field qts: %w", err)
		}
		g.Qts = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesGetDifferenceRequest.
var (
	_ bin.Encoder = &UpdatesGetDifferenceRequest{}
	_ bin.Decoder = &UpdatesGetDifferenceRequest{}
)

// UpdatesGetDifference invokes method updates.getDifference#25939651 returning error if any.
// Get new updates¹.
//
// Links:
//  1) https://core.telegram.org/api/updates
//
// Possible errors:
//  401 AUTH_KEY_PERM_EMPTY: The temporary auth key must be binded to the permanent auth key to use these methods.
//  400 CDN_METHOD_INVALID: You can't call this method in a CDN DC
//  400 DATE_EMPTY: Date empty
//  400 PERSISTENT_TIMESTAMP_EMPTY: Persistent timestamp empty
//  400 PERSISTENT_TIMESTAMP_INVALID: Persistent timestamp invalid
//
// See https://core.telegram.org/method/updates.getDifference for reference.
// Can be used by bots.
func (c *Client) UpdatesGetDifference(ctx context.Context, request *UpdatesGetDifferenceRequest) (UpdatesDifferenceClass, error) {
	var result UpdatesDifferenceBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Difference, nil
}
