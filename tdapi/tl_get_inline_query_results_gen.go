// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// GetInlineQueryResultsRequest represents TL type `getInlineQueryResults#79dcf86c`.
type GetInlineQueryResultsRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// Identifier of the chat where the query was sent
	ChatID int64
	// Location of the user; pass null if unknown or the bot doesn't need user's location
	UserLocation Location
	// Text of the query
	Query string
	// Offset of the first entry to return; use empty string to get the first chunk of
	// results
	Offset string
}

// GetInlineQueryResultsRequestTypeID is TL type id of GetInlineQueryResultsRequest.
const GetInlineQueryResultsRequestTypeID = 0x79dcf86c

// Ensuring interfaces in compile-time for GetInlineQueryResultsRequest.
var (
	_ bin.Encoder     = &GetInlineQueryResultsRequest{}
	_ bin.Decoder     = &GetInlineQueryResultsRequest{}
	_ bin.BareEncoder = &GetInlineQueryResultsRequest{}
	_ bin.BareDecoder = &GetInlineQueryResultsRequest{}
)

func (g *GetInlineQueryResultsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BotUserID == 0) {
		return false
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.UserLocation.Zero()) {
		return false
	}
	if !(g.Query == "") {
		return false
	}
	if !(g.Offset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetInlineQueryResultsRequest) String() string {
	if g == nil {
		return "GetInlineQueryResultsRequest(nil)"
	}
	type Alias GetInlineQueryResultsRequest
	return fmt.Sprintf("GetInlineQueryResultsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetInlineQueryResultsRequest) TypeID() uint32 {
	return GetInlineQueryResultsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetInlineQueryResultsRequest) TypeName() string {
	return "getInlineQueryResults"
}

// TypeInfo returns info about TL type.
func (g *GetInlineQueryResultsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getInlineQueryResults",
		ID:   GetInlineQueryResultsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserLocation",
			SchemaName: "user_location",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetInlineQueryResultsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineQueryResults#79dcf86c as nil")
	}
	b.PutID(GetInlineQueryResultsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetInlineQueryResultsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineQueryResults#79dcf86c as nil")
	}
	b.PutInt53(g.BotUserID)
	b.PutInt53(g.ChatID)
	if err := g.UserLocation.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getInlineQueryResults#79dcf86c: field user_location: %w", err)
	}
	b.PutString(g.Query)
	b.PutString(g.Offset)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetInlineQueryResultsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineQueryResults#79dcf86c to nil")
	}
	if err := b.ConsumeID(GetInlineQueryResultsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetInlineQueryResultsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineQueryResults#79dcf86c to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field bot_user_id: %w", err)
		}
		g.BotUserID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		if err := g.UserLocation.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field user_location: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field query: %w", err)
		}
		g.Query = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field offset: %w", err)
		}
		g.Offset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetInlineQueryResultsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineQueryResults#79dcf86c as nil")
	}
	b.ObjStart()
	b.PutID("getInlineQueryResults")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(g.BotUserID)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("user_location")
	if err := g.UserLocation.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getInlineQueryResults#79dcf86c: field user_location: %w", err)
	}
	b.Comma()
	b.FieldStart("query")
	b.PutString(g.Query)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetInlineQueryResultsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineQueryResults#79dcf86c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getInlineQueryResults"); err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field bot_user_id: %w", err)
			}
			g.BotUserID = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field chat_id: %w", err)
			}
			g.ChatID = value
		case "user_location":
			if err := g.UserLocation.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field user_location: %w", err)
			}
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field query: %w", err)
			}
			g.Query = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineQueryResults#79dcf86c: field offset: %w", err)
			}
			g.Offset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (g *GetInlineQueryResultsRequest) GetBotUserID() (value int64) {
	if g == nil {
		return
	}
	return g.BotUserID
}

// GetChatID returns value of ChatID field.
func (g *GetInlineQueryResultsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetUserLocation returns value of UserLocation field.
func (g *GetInlineQueryResultsRequest) GetUserLocation() (value Location) {
	if g == nil {
		return
	}
	return g.UserLocation
}

// GetQuery returns value of Query field.
func (g *GetInlineQueryResultsRequest) GetQuery() (value string) {
	if g == nil {
		return
	}
	return g.Query
}

// GetOffset returns value of Offset field.
func (g *GetInlineQueryResultsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetInlineQueryResults invokes method getInlineQueryResults#79dcf86c returning error if any.
func (c *Client) GetInlineQueryResults(ctx context.Context, request *GetInlineQueryResultsRequest) (*InlineQueryResults, error) {
	var result InlineQueryResults

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
