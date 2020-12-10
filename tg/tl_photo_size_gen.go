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

// PhotoSizeEmpty represents TL type `photoSizeEmpty#e17e23c`.
// Empty constructor. Image with this thumbnail is unavailable.
//
// See https://core.telegram.org/constructor/photoSizeEmpty for reference.
type PhotoSizeEmpty struct {
	// Thumbnail type (see. photoSize)
	Type string
}

// PhotoSizeEmptyTypeID is TL type id of PhotoSizeEmpty.
const PhotoSizeEmptyTypeID = 0xe17e23c

// Encode implements bin.Encoder.
func (p *PhotoSizeEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeEmpty#e17e23c as nil")
	}
	b.PutID(PhotoSizeEmptyTypeID)
	b.PutString(p.Type)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSizeEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeEmpty#e17e23c to nil")
	}
	if err := b.ConsumeID(PhotoSizeEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeEmpty#e17e23c: field type: %w", err)
		}
		p.Type = value
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoSizeEmpty) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSizeEmpty.
var (
	_ bin.Encoder = &PhotoSizeEmpty{}
	_ bin.Decoder = &PhotoSizeEmpty{}

	_ PhotoSizeClass = &PhotoSizeEmpty{}
)

// PhotoSize represents TL type `photoSize#77bfb61b`.
// Image description.
//
// See https://core.telegram.org/constructor/photoSize for reference.
type PhotoSize struct {
	// Thumbnail type
	Type string
	// File location
	Location FileLocationToBeDeprecated
	// Image width
	W int
	// Image height
	H int
	// File size
	Size int
}

// PhotoSizeTypeID is TL type id of PhotoSize.
const PhotoSizeTypeID = 0x77bfb61b

// Encode implements bin.Encoder.
func (p *PhotoSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSize#77bfb61b as nil")
	}
	b.PutID(PhotoSizeTypeID)
	b.PutString(p.Type)
	if err := p.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photoSize#77bfb61b: field location: %w", err)
	}
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutInt(p.Size)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSize#77bfb61b to nil")
	}
	if err := b.ConsumeID(PhotoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSize#77bfb61b: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field type: %w", err)
		}
		p.Type = value
	}
	{
		if err := p.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field location: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSize#77bfb61b: field size: %w", err)
		}
		p.Size = value
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSize.
var (
	_ bin.Encoder = &PhotoSize{}
	_ bin.Decoder = &PhotoSize{}

	_ PhotoSizeClass = &PhotoSize{}
)

// PhotoCachedSize represents TL type `photoCachedSize#e9a734fa`.
// Description of an image and its content.
//
// See https://core.telegram.org/constructor/photoCachedSize for reference.
type PhotoCachedSize struct {
	// Thumbnail type
	Type string
	// File location
	Location FileLocationToBeDeprecated
	// Image width
	W int
	// Image height
	H int
	// Binary data, file content
	Bytes []byte
}

// PhotoCachedSizeTypeID is TL type id of PhotoCachedSize.
const PhotoCachedSizeTypeID = 0xe9a734fa

// Encode implements bin.Encoder.
func (p *PhotoCachedSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoCachedSize#e9a734fa as nil")
	}
	b.PutID(PhotoCachedSizeTypeID)
	b.PutString(p.Type)
	if err := p.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photoCachedSize#e9a734fa: field location: %w", err)
	}
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoCachedSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoCachedSize#e9a734fa to nil")
	}
	if err := b.ConsumeID(PhotoCachedSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field type: %w", err)
		}
		p.Type = value
	}
	{
		if err := p.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field location: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field h: %w", err)
		}
		p.H = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoCachedSize#e9a734fa: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoCachedSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoCachedSize.
var (
	_ bin.Encoder = &PhotoCachedSize{}
	_ bin.Decoder = &PhotoCachedSize{}

	_ PhotoSizeClass = &PhotoCachedSize{}
)

// PhotoStrippedSize represents TL type `photoStrippedSize#e0b0bc2e`.
// A low-resolution compressed JPG payload
//
// See https://core.telegram.org/constructor/photoStrippedSize for reference.
type PhotoStrippedSize struct {
	// Thumbnail type
	Type string
	// Thumbnail data, see here for more info on decompression »
	Bytes []byte
}

// PhotoStrippedSizeTypeID is TL type id of PhotoStrippedSize.
const PhotoStrippedSizeTypeID = 0xe0b0bc2e

// Encode implements bin.Encoder.
func (p *PhotoStrippedSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoStrippedSize#e0b0bc2e as nil")
	}
	b.PutID(PhotoStrippedSizeTypeID)
	b.PutString(p.Type)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoStrippedSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoStrippedSize#e0b0bc2e to nil")
	}
	if err := b.ConsumeID(PhotoStrippedSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoStrippedSize#e0b0bc2e: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoStrippedSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoStrippedSize.
var (
	_ bin.Encoder = &PhotoStrippedSize{}
	_ bin.Decoder = &PhotoStrippedSize{}

	_ PhotoSizeClass = &PhotoStrippedSize{}
)

// PhotoSizeProgressive represents TL type `photoSizeProgressive#5aa86a51`.
// Progressively encoded photosize
//
// See https://core.telegram.org/constructor/photoSizeProgressive for reference.
type PhotoSizeProgressive struct {
	// Photosize type
	Type string
	// File location
	Location FileLocationToBeDeprecated
	// Photo width
	W int
	// Photo height
	H int
	// Sizes of progressive JPEG file prefixes, which can be used to preliminarily show the image.
	Sizes []int
}

// PhotoSizeProgressiveTypeID is TL type id of PhotoSizeProgressive.
const PhotoSizeProgressiveTypeID = 0x5aa86a51

// Encode implements bin.Encoder.
func (p *PhotoSizeProgressive) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoSizeProgressive#5aa86a51 as nil")
	}
	b.PutID(PhotoSizeProgressiveTypeID)
	b.PutString(p.Type)
	if err := p.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photoSizeProgressive#5aa86a51: field location: %w", err)
	}
	b.PutInt(p.W)
	b.PutInt(p.H)
	b.PutVectorHeader(len(p.Sizes))
	for _, v := range p.Sizes {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoSizeProgressive) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoSizeProgressive#5aa86a51 to nil")
	}
	if err := b.ConsumeID(PhotoSizeProgressiveTypeID); err != nil {
		return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field type: %w", err)
		}
		p.Type = value
	}
	{
		if err := p.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field location: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field w: %w", err)
		}
		p.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field h: %w", err)
		}
		p.H = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field sizes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode photoSizeProgressive#5aa86a51: field sizes: %w", err)
			}
			p.Sizes = append(p.Sizes, value)
		}
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoSizeProgressive) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoSizeProgressive.
var (
	_ bin.Encoder = &PhotoSizeProgressive{}
	_ bin.Decoder = &PhotoSizeProgressive{}

	_ PhotoSizeClass = &PhotoSizeProgressive{}
)

// PhotoPathSize represents TL type `photoPathSize#d8214d41`.
// Messages with animated stickers can have a compressed svg (< 300 bytes) to show the outline of the sticker before fetching the actual lottie animation.
//
// See https://core.telegram.org/constructor/photoPathSize for reference.
type PhotoPathSize struct {
	// Always j
	Type string
	// Compressed SVG path payload, see here for decompression instructions
	Bytes []byte
}

// PhotoPathSizeTypeID is TL type id of PhotoPathSize.
const PhotoPathSizeTypeID = 0xd8214d41

// Encode implements bin.Encoder.
func (p *PhotoPathSize) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoPathSize#d8214d41 as nil")
	}
	b.PutID(PhotoPathSizeTypeID)
	b.PutString(p.Type)
	b.PutBytes(p.Bytes)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoPathSize) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoPathSize#d8214d41 to nil")
	}
	if err := b.ConsumeID(PhotoPathSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode photoPathSize#d8214d41: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode photoPathSize#d8214d41: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photoPathSize#d8214d41: field bytes: %w", err)
		}
		p.Bytes = value
	}
	return nil
}

// construct implements constructor of PhotoSizeClass.
func (p PhotoPathSize) construct() PhotoSizeClass { return &p }

// Ensuring interfaces in compile-time for PhotoPathSize.
var (
	_ bin.Encoder = &PhotoPathSize{}
	_ bin.Decoder = &PhotoPathSize{}

	_ PhotoSizeClass = &PhotoPathSize{}
)

// PhotoSizeClass represents PhotoSize generic type.
//
// See https://core.telegram.org/type/PhotoSize for reference.
//
// Example:
//  g, err := DecodePhotoSize(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PhotoSizeEmpty: // photoSizeEmpty#e17e23c
//  case *PhotoSize: // photoSize#77bfb61b
//  case *PhotoCachedSize: // photoCachedSize#e9a734fa
//  case *PhotoStrippedSize: // photoStrippedSize#e0b0bc2e
//  case *PhotoSizeProgressive: // photoSizeProgressive#5aa86a51
//  case *PhotoPathSize: // photoPathSize#d8214d41
//  default: panic(v)
//  }
type PhotoSizeClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhotoSizeClass
}

// DecodePhotoSize implements binary de-serialization for PhotoSizeClass.
func DecodePhotoSize(buf *bin.Buffer) (PhotoSizeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotoSizeEmptyTypeID:
		// Decoding photoSizeEmpty#e17e23c.
		v := PhotoSizeEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoSizeTypeID:
		// Decoding photoSize#77bfb61b.
		v := PhotoSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoCachedSizeTypeID:
		// Decoding photoCachedSize#e9a734fa.
		v := PhotoCachedSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoStrippedSizeTypeID:
		// Decoding photoStrippedSize#e0b0bc2e.
		v := PhotoStrippedSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoSizeProgressiveTypeID:
		// Decoding photoSizeProgressive#5aa86a51.
		v := PhotoSizeProgressive{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	case PhotoPathSizeTypeID:
		// Decoding photoPathSize#d8214d41.
		v := PhotoPathSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoSizeClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhotoSize boxes the PhotoSizeClass providing a helper.
type PhotoSizeBox struct {
	PhotoSize PhotoSizeClass
}

// Decode implements bin.Decoder for PhotoSizeBox.
func (b *PhotoSizeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotoSizeBox to nil")
	}
	v, err := DecodePhotoSize(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhotoSize = v
	return nil
}

// Encode implements bin.Encode for PhotoSizeBox.
func (b *PhotoSizeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhotoSize == nil {
		return fmt.Errorf("unable to encode PhotoSizeClass as nil")
	}
	return b.PhotoSize.Encode(buf)
}
