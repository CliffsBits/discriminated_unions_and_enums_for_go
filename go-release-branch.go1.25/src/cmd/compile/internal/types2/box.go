// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

import (
	"cmd/compile/internal/syntax"
	. "internal/types/errors"
)

// ----------------------------------------------------------------------------
// API

// A Box represents a box type (discriminated union).
type Box struct {
	types []Type // variant types
}

// NewBox returns a new box with the given variant types.
func NewBox(types []Type) *Box {
	b := &Box{types: types}
	b.markComplete()
	return b
}

// NumTypes returns the number of variant types in the box.
func (b *Box) NumTypes() int { return len(b.types) }

// Type returns the i'th variant type for 0 <= i < NumTypes().
func (b *Box) Type(i int) Type { return b.types[i] }

// ----------------------------------------------------------------------------
// Implementation

func (b *Box) Underlying() Type { return b }
func (b *Box) String() string   { return TypeString(b, nil) }

// markComplete marks the box type as complete.
func (b *Box) markComplete() {}

// ----------------------------------------------------------------------------
// type checking

func (check *Checker) boxType(typ *Box, e *syntax.BoxType) {
	if e.TypeList == nil {
		return
	}

	var types []Type
	for _, f := range e.TypeList {
		if f.Type != nil {
			t := check.varType(f.Type)
			if isValid(t) {
				// Validate that only value types and named types are allowed in boxes
				if !isValidBoxType(t) {
					check.errorf(f, InvalidTypeCycle, "invalid type %s in box (only basic types and named types are allowed; anonymous structs are not permitted)", t)
					continue
				}
				types = append(types, t)
			}
		}
	}
	typ.types = types
}

// isValidBoxType checks if a type is valid for use in a box.
// Only basic types, named types, and struct types are allowed.
// Slices, maps, channels, functions, interfaces, and pointers are not allowed.
func isValidBoxType(t Type) bool {
	// Allow named types (which includes user-defined types)
	if named, ok := t.(*Named); ok {
		// Check the underlying type of the named type
		underlying := named.Underlying()
		// Named types are OK as long as their underlying type is valid
		// For structs, we allow them only through named types
		if _, isStruct := underlying.(*Struct); isStruct {
			return true // Named struct types are allowed
		}
		// For other types, recursively check validity
		return isValidBoxType(underlying)
	}
	
	// Allow basic types (int, string, float64, bool, etc.)
	if _, ok := t.(*Basic); ok {
		return true
	}
	
	// Reject anonymous struct types and everything else
	// Anonymous structs, slices, maps, channels, functions, interfaces, pointers are not allowed
	return false
}