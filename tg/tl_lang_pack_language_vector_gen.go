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

// LangPackLanguageVector is a box for Vector<LangPackLanguage>
type LangPackLanguageVector struct {
	// Elements of Vector<LangPackLanguage>
	Elems []LangPackLanguage
}

// Encode implements bin.Encoder.
func (vec *LangPackLanguageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<LangPackLanguage> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<LangPackLanguage>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *LangPackLanguageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<LangPackLanguage> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<LangPackLanguage>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LangPackLanguage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<LangPackLanguage>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for LangPackLanguageVector.
var (
	_ bin.Encoder = &LangPackLanguageVector{}
	_ bin.Decoder = &LangPackLanguageVector{}
)
