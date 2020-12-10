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

// True represents TL type `true#3fedd339`.
// See predefined identifiers.
//
// See https://core.telegram.org/constructor/true for reference.
type True struct {
}

// TrueTypeID is TL type id of True.
const TrueTypeID = 0x3fedd339

// Encode implements bin.Encoder.
func (t *True) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode true#3fedd339 as nil")
	}
	b.PutID(TrueTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (t *True) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode true#3fedd339 to nil")
	}
	if err := b.ConsumeID(TrueTypeID); err != nil {
		return fmt.Errorf("unable to decode true#3fedd339: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for True.
var (
	_ bin.Encoder = &True{}
	_ bin.Decoder = &True{}
)
