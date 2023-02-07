//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesEmojiGroupsClassArray is adapter for slice of MessagesEmojiGroupsClass.
type MessagesEmojiGroupsClassArray []MessagesEmojiGroupsClass

// Sort sorts slice of MessagesEmojiGroupsClass.
func (s MessagesEmojiGroupsClassArray) Sort(less func(a, b MessagesEmojiGroupsClass) bool) MessagesEmojiGroupsClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesEmojiGroupsClass.
func (s MessagesEmojiGroupsClassArray) SortStable(less func(a, b MessagesEmojiGroupsClass) bool) MessagesEmojiGroupsClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesEmojiGroupsClass.
func (s MessagesEmojiGroupsClassArray) Retain(keep func(x MessagesEmojiGroupsClass) bool) MessagesEmojiGroupsClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesEmojiGroupsClassArray) First() (v MessagesEmojiGroupsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesEmojiGroupsClassArray) Last() (v MessagesEmojiGroupsClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesEmojiGroupsClassArray) PopFirst() (v MessagesEmojiGroupsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesEmojiGroupsClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesEmojiGroupsClassArray) Pop() (v MessagesEmojiGroupsClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesEmojiGroups returns copy with only MessagesEmojiGroups constructors.
func (s MessagesEmojiGroupsClassArray) AsMessagesEmojiGroups() (to MessagesEmojiGroupsArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesEmojiGroups)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesEmojiGroupsClassArray) AppendOnlyModified(to []*MessagesEmojiGroups) []*MessagesEmojiGroups {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesEmojiGroupsClassArray) AsModified() (to []*MessagesEmojiGroups) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesEmojiGroupsClassArray) FirstAsModified() (v *MessagesEmojiGroups, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesEmojiGroupsClassArray) LastAsModified() (v *MessagesEmojiGroups, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesEmojiGroupsClassArray) PopFirstAsModified() (v *MessagesEmojiGroups, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesEmojiGroupsClassArray) PopAsModified() (v *MessagesEmojiGroups, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesEmojiGroupsArray is adapter for slice of MessagesEmojiGroups.
type MessagesEmojiGroupsArray []MessagesEmojiGroups

// Sort sorts slice of MessagesEmojiGroups.
func (s MessagesEmojiGroupsArray) Sort(less func(a, b MessagesEmojiGroups) bool) MessagesEmojiGroupsArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesEmojiGroups.
func (s MessagesEmojiGroupsArray) SortStable(less func(a, b MessagesEmojiGroups) bool) MessagesEmojiGroupsArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesEmojiGroups.
func (s MessagesEmojiGroupsArray) Retain(keep func(x MessagesEmojiGroups) bool) MessagesEmojiGroupsArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesEmojiGroupsArray) First() (v MessagesEmojiGroups, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesEmojiGroupsArray) Last() (v MessagesEmojiGroups, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesEmojiGroupsArray) PopFirst() (v MessagesEmojiGroups, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesEmojiGroups
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesEmojiGroupsArray) Pop() (v MessagesEmojiGroups, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
