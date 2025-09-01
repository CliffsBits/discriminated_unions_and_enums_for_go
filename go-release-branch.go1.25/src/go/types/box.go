// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// ----------------------------------------------------------------------------
// API

// A Box represents a box type (discriminated union).
type Box struct {
	variants []Type // list of variant types
}

// NewBox returns a new [Box] type with the given variant types.
// It is an error to create an empty box; they are syntactically not possible.
func NewBox(variants []Type) *Box {
	if len(variants) == 0 {
		panic("empty box")
	}
	return &Box{variants}
}

func (b *Box) Len() int          { return len(b.variants) }
func (b *Box) Variant(i int) Type { return b.variants[i] }

func (b *Box) Underlying() Type { return b }
func (b *Box) String() string   { return TypeString(b, nil) }