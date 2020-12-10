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

// EmojiURL represents TL type `emojiURL#a575739d`.
// An HTTP URL which can be used to automatically log in into translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
//
// See https://core.telegram.org/constructor/emojiURL for reference.
type EmojiURL struct {
	// An HTTP URL which can be used to automatically log in into translation platform and suggest new emoji replacements. The URL will be valid for 30 seconds after generation
	URL string
}

// EmojiURLTypeID is TL type id of EmojiURL.
const EmojiURLTypeID = 0xa575739d

// Encode implements bin.Encoder.
func (e *EmojiURL) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiURL#a575739d as nil")
	}
	b.PutID(EmojiURLTypeID)
	b.PutString(e.URL)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiURL) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiURL#a575739d to nil")
	}
	if err := b.ConsumeID(EmojiURLTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiURL#a575739d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emojiURL#a575739d: field url: %w", err)
		}
		e.URL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for EmojiURL.
var (
	_ bin.Encoder = &EmojiURL{}
	_ bin.Decoder = &EmojiURL{}
)
