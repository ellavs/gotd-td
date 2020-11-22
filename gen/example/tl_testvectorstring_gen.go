// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}

// A simple object containing a vector of strings; for testing only
type TestVectorString struct {
	// Vector of strings
	Value []string
}

// TestVectorStringTypeID is TL type id of TestVectorString.
const TestVectorStringTypeID = 0x5d6f85bc

// Encode implements bin.Encoder.
func (t *TestVectorString) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorString#5d6f85bc as nil")
	}
	b.PutID(TestVectorStringTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorString) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorString#5d6f85bc to nil")
	}
	if err := b.ConsumeID(TestVectorStringTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorString#5d6f85bc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorString#5d6f85bc: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorString#5d6f85bc: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorString.
var (
	_ bin.Encoder = &TestVectorString{}
	_ bin.Decoder = &TestVectorString{}
)
