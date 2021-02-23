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

// AccountGetWallPaperRequest represents TL type `account.getWallPaper#fc8ddbea`.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
type AccountGetWallPaperRequest struct {
	// The wallpaper to get info about
	Wallpaper InputWallPaperClass `schemaname:"wallpaper"`
}

// AccountGetWallPaperRequestTypeID is TL type id of AccountGetWallPaperRequest.
const AccountGetWallPaperRequestTypeID = 0xfc8ddbea

func (g *AccountGetWallPaperRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Wallpaper == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetWallPaperRequest) String() string {
	if g == nil {
		return "AccountGetWallPaperRequest(nil)"
	}
	type Alias AccountGetWallPaperRequest
	return fmt.Sprintf("AccountGetWallPaperRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetWallPaperRequest from given interface.
func (g *AccountGetWallPaperRequest) FillFrom(from interface {
	GetWallpaper() (value InputWallPaperClass)
}) {
	g.Wallpaper = from.GetWallpaper()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *AccountGetWallPaperRequest) TypeID() uint32 {
	return AccountGetWallPaperRequestTypeID
}

// SchemaName returns MTProto type name.
func (g *AccountGetWallPaperRequest) SchemaName() string {
	return "account.getWallPaper"
}

// Encode implements bin.Encoder.
func (g *AccountGetWallPaperRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getWallPaper#fc8ddbea as nil")
	}
	b.PutID(AccountGetWallPaperRequestTypeID)
	if g.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper is nil")
	}
	if err := g.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
	}
	return nil
}

// GetWallpaper returns value of Wallpaper field.
func (g *AccountGetWallPaperRequest) GetWallpaper() (value InputWallPaperClass) {
	return g.Wallpaper
}

// Decode implements bin.Decoder.
func (g *AccountGetWallPaperRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getWallPaper#fc8ddbea to nil")
	}
	if err := b.ConsumeID(AccountGetWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: %w", err)
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getWallPaper#fc8ddbea: field wallpaper: %w", err)
		}
		g.Wallpaper = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetWallPaperRequest.
var (
	_ bin.Encoder = &AccountGetWallPaperRequest{}
	_ bin.Decoder = &AccountGetWallPaperRequest{}
)

// AccountGetWallPaper invokes method account.getWallPaper#fc8ddbea returning error if any.
// Get info about a certain wallpaper
//
// See https://core.telegram.org/method/account.getWallPaper for reference.
func (c *Client) AccountGetWallPaper(ctx context.Context, wallpaper InputWallPaperClass) (WallPaperClass, error) {
	var result WallPaperBox

	request := &AccountGetWallPaperRequest{
		Wallpaper: wallpaper,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}
