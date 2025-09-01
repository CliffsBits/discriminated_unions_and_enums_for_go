// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

// boxValue represents a discriminated union value
type boxValue struct {
	_type *_type
	data  unsafe.Pointer
}

//go:nosplit
//go:norace
func assertB2V(b unsafe.Pointer, typ *byte) unsafe.Pointer {
	// assertB2V performs a box type assertion.
	// This is nosplit because it may be called during bootstrap
	// when stack splitting is unsafe.
	// This is norace because it accesses raw memory structures
	// that the race detector cannot safely instrument.
	
	if b == nil {
		return nil
	}
	
	// Box has same layout as interface: [*_type, data]
	boxType := *(**_type)(b)        // First word: type pointer
	boxData := *(*unsafe.Pointer)(add(b, 8))  // Second word: data pointer
	
	if boxType == nil {
		return nil
	}
	
	// Convert typ parameter to *_type
	// The typ parameter is actually a *_type passed as *byte
	targetType := (*_type)(unsafe.Pointer(typ))
	
	// Use proper type comparison instead of pointer equality
	// This handles types from different compilation scopes
	if typesEqual(boxType, targetType, make(map[_typePair]struct{})) {
		return boxData
	}
	
	// Panic on type mismatch for single-value assertion
	panic(&TypeAssertionError{
		_interface: nil, // Box types don't have an interface type
		concrete:   boxType,
		asserted:   targetType,
	})
}

