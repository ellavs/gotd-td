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

// JsonNull represents TL type `jsonNull#3f6d7b68`.
// null JSON value
//
// See https://core.telegram.org/constructor/jsonNull for reference.
type JsonNull struct {
}

// JsonNullTypeID is TL type id of JsonNull.
const JsonNullTypeID = 0x3f6d7b68

// Encode implements bin.Encoder.
func (j *JsonNull) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonNull#3f6d7b68 as nil")
	}
	b.PutID(JsonNullTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonNull) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonNull#3f6d7b68 to nil")
	}
	if err := b.ConsumeID(JsonNullTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonNull#3f6d7b68: %w", err)
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonNull) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonNull.
var (
	_ bin.Encoder = &JsonNull{}
	_ bin.Decoder = &JsonNull{}

	_ JSONValueClass = &JsonNull{}
)

// JsonBool represents TL type `jsonBool#c7345e6a`.
// JSON boolean value
//
// See https://core.telegram.org/constructor/jsonBool for reference.
type JsonBool struct {
	// Value
	Value bool
}

// JsonBoolTypeID is TL type id of JsonBool.
const JsonBoolTypeID = 0xc7345e6a

// Encode implements bin.Encoder.
func (j *JsonBool) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonBool#c7345e6a as nil")
	}
	b.PutID(JsonBoolTypeID)
	b.PutBool(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonBool) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonBool#c7345e6a to nil")
	}
	if err := b.ConsumeID(JsonBoolTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonBool#c7345e6a: %w", err)
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode jsonBool#c7345e6a: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonBool) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonBool.
var (
	_ bin.Encoder = &JsonBool{}
	_ bin.Decoder = &JsonBool{}

	_ JSONValueClass = &JsonBool{}
)

// JsonNumber represents TL type `jsonNumber#2be0dfa4`.
// JSON numeric value
//
// See https://core.telegram.org/constructor/jsonNumber for reference.
type JsonNumber struct {
	// Value
	Value float64
}

// JsonNumberTypeID is TL type id of JsonNumber.
const JsonNumberTypeID = 0x2be0dfa4

// Encode implements bin.Encoder.
func (j *JsonNumber) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonNumber#2be0dfa4 as nil")
	}
	b.PutID(JsonNumberTypeID)
	b.PutDouble(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonNumber) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonNumber#2be0dfa4 to nil")
	}
	if err := b.ConsumeID(JsonNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonNumber#2be0dfa4: %w", err)
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode jsonNumber#2be0dfa4: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonNumber) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonNumber.
var (
	_ bin.Encoder = &JsonNumber{}
	_ bin.Decoder = &JsonNumber{}

	_ JSONValueClass = &JsonNumber{}
)

// JsonString represents TL type `jsonString#b71e767a`.
// JSON string
//
// See https://core.telegram.org/constructor/jsonString for reference.
type JsonString struct {
	// Value
	Value string
}

// JsonStringTypeID is TL type id of JsonString.
const JsonStringTypeID = 0xb71e767a

// Encode implements bin.Encoder.
func (j *JsonString) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonString#b71e767a as nil")
	}
	b.PutID(JsonStringTypeID)
	b.PutString(j.Value)
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonString) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonString#b71e767a to nil")
	}
	if err := b.ConsumeID(JsonStringTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonString#b71e767a: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode jsonString#b71e767a: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonString) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonString.
var (
	_ bin.Encoder = &JsonString{}
	_ bin.Decoder = &JsonString{}

	_ JSONValueClass = &JsonString{}
)

// JsonArray represents TL type `jsonArray#f7444763`.
// JSON array
//
// See https://core.telegram.org/constructor/jsonArray for reference.
type JsonArray struct {
	// JSON values
	Value []JSONValueClass
}

// JsonArrayTypeID is TL type id of JsonArray.
const JsonArrayTypeID = 0xf7444763

// Encode implements bin.Encoder.
func (j *JsonArray) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonArray#f7444763 as nil")
	}
	b.PutID(JsonArrayTypeID)
	b.PutVectorHeader(len(j.Value))
	for idx, v := range j.Value {
		if v == nil {
			return fmt.Errorf("unable to encode jsonArray#f7444763: field value element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode jsonArray#f7444763: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonArray) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonArray#f7444763 to nil")
	}
	if err := b.ConsumeID(JsonArrayTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonArray#f7444763: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode jsonArray#f7444763: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeJSONValue(b)
			if err != nil {
				return fmt.Errorf("unable to decode jsonArray#f7444763: field value: %w", err)
			}
			j.Value = append(j.Value, value)
		}
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonArray) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonArray.
var (
	_ bin.Encoder = &JsonArray{}
	_ bin.Decoder = &JsonArray{}

	_ JSONValueClass = &JsonArray{}
)

// JsonObject represents TL type `jsonObject#99c1d49d`.
// JSON object value
//
// See https://core.telegram.org/constructor/jsonObject for reference.
type JsonObject struct {
	// Values
	Value []JsonObjectValue
}

// JsonObjectTypeID is TL type id of JsonObject.
const JsonObjectTypeID = 0x99c1d49d

// Encode implements bin.Encoder.
func (j *JsonObject) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonObject#99c1d49d as nil")
	}
	b.PutID(JsonObjectTypeID)
	b.PutVectorHeader(len(j.Value))
	for idx, v := range j.Value {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode jsonObject#99c1d49d: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JsonObject) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonObject#99c1d49d to nil")
	}
	if err := b.ConsumeID(JsonObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonObject#99c1d49d: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode jsonObject#99c1d49d: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value JsonObjectValue
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode jsonObject#99c1d49d: field value: %w", err)
			}
			j.Value = append(j.Value, value)
		}
	}
	return nil
}

// construct implements constructor of JSONValueClass.
func (j JsonObject) construct() JSONValueClass { return &j }

// Ensuring interfaces in compile-time for JsonObject.
var (
	_ bin.Encoder = &JsonObject{}
	_ bin.Decoder = &JsonObject{}

	_ JSONValueClass = &JsonObject{}
)

// JSONValueClass represents JSONValue generic type.
//
// See https://core.telegram.org/type/JSONValue for reference.
//
// Example:
//  g, err := DecodeJSONValue(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *JsonNull: // jsonNull#3f6d7b68
//  case *JsonBool: // jsonBool#c7345e6a
//  case *JsonNumber: // jsonNumber#2be0dfa4
//  case *JsonString: // jsonString#b71e767a
//  case *JsonArray: // jsonArray#f7444763
//  case *JsonObject: // jsonObject#99c1d49d
//  default: panic(v)
//  }
type JSONValueClass interface {
	bin.Encoder
	bin.Decoder
	construct() JSONValueClass
}

// DecodeJSONValue implements binary de-serialization for JSONValueClass.
func DecodeJSONValue(buf *bin.Buffer) (JSONValueClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case JsonNullTypeID:
		// Decoding jsonNull#3f6d7b68.
		v := JsonNull{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JsonBoolTypeID:
		// Decoding jsonBool#c7345e6a.
		v := JsonBool{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JsonNumberTypeID:
		// Decoding jsonNumber#2be0dfa4.
		v := JsonNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JsonStringTypeID:
		// Decoding jsonString#b71e767a.
		v := JsonString{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JsonArrayTypeID:
		// Decoding jsonArray#f7444763.
		v := JsonArray{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	case JsonObjectTypeID:
		// Decoding jsonObject#99c1d49d.
		v := JsonObject{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode JSONValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode JSONValueClass: %w", bin.NewUnexpectedID(id))
	}
}

// JSONValue boxes the JSONValueClass providing a helper.
type JSONValueBox struct {
	JSONValue JSONValueClass
}

// Decode implements bin.Decoder for JSONValueBox.
func (b *JSONValueBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode JSONValueBox to nil")
	}
	v, err := DecodeJSONValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.JSONValue = v
	return nil
}

// Encode implements bin.Encode for JSONValueBox.
func (b *JSONValueBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.JSONValue == nil {
		return fmt.Errorf("unable to encode JSONValueClass as nil")
	}
	return b.JSONValue.Encode(buf)
}
