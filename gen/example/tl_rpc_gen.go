// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}

// Client describes RPC methods of TL schema.
type Client interface {
	Ping(ctx context.Context, id int32) error
	Send(ctx context.Context, msg SMS) (SMS, error)
	SendMultipleSMS(ctx context.Context, messages []SMS) error
}
