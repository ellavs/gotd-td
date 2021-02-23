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

// KeyboardButtonRow represents TL type `keyboardButtonRow#77608b83`.
// Inline keyboard row
//
// See https://core.telegram.org/constructor/keyboardButtonRow for reference.
type KeyboardButtonRow struct {
	// Bot or inline keyboard buttons
	Buttons []KeyboardButtonClass `schemaname:"buttons"`
}

// KeyboardButtonRowTypeID is TL type id of KeyboardButtonRow.
const KeyboardButtonRowTypeID = 0x77608b83

func (k *KeyboardButtonRow) Zero() bool {
	if k == nil {
		return true
	}
	if !(k.Buttons == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (k *KeyboardButtonRow) String() string {
	if k == nil {
		return "KeyboardButtonRow(nil)"
	}
	type Alias KeyboardButtonRow
	return fmt.Sprintf("KeyboardButtonRow%+v", Alias(*k))
}

// FillFrom fills KeyboardButtonRow from given interface.
func (k *KeyboardButtonRow) FillFrom(from interface {
	GetButtons() (value []KeyboardButtonClass)
}) {
	k.Buttons = from.GetButtons()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (k *KeyboardButtonRow) TypeID() uint32 {
	return KeyboardButtonRowTypeID
}

// SchemaName returns MTProto type name.
func (k *KeyboardButtonRow) SchemaName() string {
	return "keyboardButtonRow"
}

// Encode implements bin.Encoder.
func (k *KeyboardButtonRow) Encode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't encode keyboardButtonRow#77608b83 as nil")
	}
	b.PutID(KeyboardButtonRowTypeID)
	b.PutVectorHeader(len(k.Buttons))
	for idx, v := range k.Buttons {
		if v == nil {
			return fmt.Errorf("unable to encode keyboardButtonRow#77608b83: field buttons element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode keyboardButtonRow#77608b83: field buttons element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetButtons returns value of Buttons field.
func (k *KeyboardButtonRow) GetButtons() (value []KeyboardButtonClass) {
	return k.Buttons
}

// MapButtons returns field Buttons wrapped in KeyboardButtonClassSlice helper.
func (k *KeyboardButtonRow) MapButtons() (value KeyboardButtonClassSlice) {
	return KeyboardButtonClassSlice(k.Buttons)
}

// Decode implements bin.Decoder.
func (k *KeyboardButtonRow) Decode(b *bin.Buffer) error {
	if k == nil {
		return fmt.Errorf("can't decode keyboardButtonRow#77608b83 to nil")
	}
	if err := b.ConsumeID(KeyboardButtonRowTypeID); err != nil {
		return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: field buttons: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeKeyboardButton(b)
			if err != nil {
				return fmt.Errorf("unable to decode keyboardButtonRow#77608b83: field buttons: %w", err)
			}
			k.Buttons = append(k.Buttons, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for KeyboardButtonRow.
var (
	_ bin.Encoder = &KeyboardButtonRow{}
	_ bin.Decoder = &KeyboardButtonRow{}
)
