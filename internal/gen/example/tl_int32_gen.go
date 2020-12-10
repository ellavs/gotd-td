// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// Int32 represents TL type `int32#5cb934fa`.
//
// See https://localhost:80/doc/constructor/int32 for reference.
type Int32 struct {
}

// Int32TypeID is TL type id of Int32.
const Int32TypeID = 0x5cb934fa

// Encode implements bin.Encoder.
func (i *Int32) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int32#5cb934fa as nil")
	}
	b.PutID(Int32TypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *Int32) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int32#5cb934fa to nil")
	}
	if err := b.ConsumeID(Int32TypeID); err != nil {
		return fmt.Errorf("unable to decode int32#5cb934fa: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Int32.
var (
	_ bin.Encoder = &Int32{}
	_ bin.Decoder = &Int32{}
)
