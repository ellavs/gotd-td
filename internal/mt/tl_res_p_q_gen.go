// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ResPQ represents TL type `resPQ#5162463`.
type ResPQ struct {
	// Nonce field of ResPQ.
	Nonce bin.Int128 `schemaname:"nonce"`
	// ServerNonce field of ResPQ.
	ServerNonce bin.Int128 `schemaname:"server_nonce"`
	// Pq field of ResPQ.
	Pq []byte `schemaname:"pq"`
	// ServerPublicKeyFingerprints field of ResPQ.
	ServerPublicKeyFingerprints []int64 `schemaname:"server_public_key_fingerprints"`
}

// ResPQTypeID is TL type id of ResPQ.
const ResPQTypeID = 0x5162463

func (r *ResPQ) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}
	if !(r.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(r.Pq == nil) {
		return false
	}
	if !(r.ServerPublicKeyFingerprints == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResPQ) String() string {
	if r == nil {
		return "ResPQ(nil)"
	}
	type Alias ResPQ
	return fmt.Sprintf("ResPQ%+v", Alias(*r))
}

// FillFrom fills ResPQ from given interface.
func (r *ResPQ) FillFrom(from interface {
	GetNonce() (value bin.Int128)
	GetServerNonce() (value bin.Int128)
	GetPq() (value []byte)
	GetServerPublicKeyFingerprints() (value []int64)
}) {
	r.Nonce = from.GetNonce()
	r.ServerNonce = from.GetServerNonce()
	r.Pq = from.GetPq()
	r.ServerPublicKeyFingerprints = from.GetServerPublicKeyFingerprints()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *ResPQ) TypeID() uint32 {
	return ResPQTypeID
}

// SchemaName returns MTProto type name.
func (r *ResPQ) SchemaName() string {
	return "resPQ"
}

// Encode implements bin.Encoder.
func (r *ResPQ) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resPQ#5162463 as nil")
	}
	b.PutID(ResPQTypeID)
	b.PutInt128(r.Nonce)
	b.PutInt128(r.ServerNonce)
	b.PutBytes(r.Pq)
	b.PutVectorHeader(len(r.ServerPublicKeyFingerprints))
	for _, v := range r.ServerPublicKeyFingerprints {
		b.PutLong(v)
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ResPQ) GetNonce() (value bin.Int128) {
	return r.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (r *ResPQ) GetServerNonce() (value bin.Int128) {
	return r.ServerNonce
}

// GetPq returns value of Pq field.
func (r *ResPQ) GetPq() (value []byte) {
	return r.Pq
}

// GetServerPublicKeyFingerprints returns value of ServerPublicKeyFingerprints field.
func (r *ResPQ) GetServerPublicKeyFingerprints() (value []int64) {
	return r.ServerPublicKeyFingerprints
}

// Decode implements bin.Decoder.
func (r *ResPQ) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resPQ#5162463 to nil")
	}
	if err := b.ConsumeID(ResPQTypeID); err != nil {
		return fmt.Errorf("unable to decode resPQ#5162463: %w", err)
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode resPQ#5162463: field nonce: %w", err)
		}
		r.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode resPQ#5162463: field server_nonce: %w", err)
		}
		r.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode resPQ#5162463: field pq: %w", err)
		}
		r.Pq = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode resPQ#5162463: field server_public_key_fingerprints: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode resPQ#5162463: field server_public_key_fingerprints: %w", err)
			}
			r.ServerPublicKeyFingerprints = append(r.ServerPublicKeyFingerprints, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ResPQ.
var (
	_ bin.Encoder = &ResPQ{}
	_ bin.Decoder = &ResPQ{}
)
