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

// StatsGroupTopInviter represents TL type `statsGroupTopInviter#31962a4c`.
// Information about an active supergroup inviter
//
// See https://core.telegram.org/constructor/statsGroupTopInviter for reference.
type StatsGroupTopInviter struct {
	// User ID
	UserID int
	// Number of invitations for statistics period in consideration
	Invitations int
}

// StatsGroupTopInviterTypeID is TL type id of StatsGroupTopInviter.
const StatsGroupTopInviterTypeID = 0x31962a4c

// Encode implements bin.Encoder.
func (s *StatsGroupTopInviter) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopInviter#31962a4c as nil")
	}
	b.PutID(StatsGroupTopInviterTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Invitations)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopInviter) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopInviter#31962a4c to nil")
	}
	if err := b.ConsumeID(StatsGroupTopInviterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: field invitations: %w", err)
		}
		s.Invitations = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopInviter.
var (
	_ bin.Encoder = &StatsGroupTopInviter{}
	_ bin.Decoder = &StatsGroupTopInviter{}
)
