// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test/integration/consul-container/libs/utils/defer.go
Defer module for internal layer

## Tags
internal

## Exports
ResettableDefer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/libs/utils/defer.go> a code:Module ;
    code:name "test/integration/consul-container/libs/utils/defer.go" ;
    code:description "Defer module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :ResettableDefer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package utils

// ResettableDefer is a way to capture a series of cleanup functions and
// bulk-cancel them. Ideal to use in a long constructor function before the
// overall Close/Stop/Terminate method is ready to use to tear down all of the
// portions properly.
type ResettableDefer struct {
	cleanupFns []func()
}

// Add registers another function to call at Execute time.
func (d *ResettableDefer) Add(f func()) {
	d.cleanupFns = append(d.cleanupFns, f)
}

// Reset clears the pending defer work.
func (d *ResettableDefer) Reset() {
	d.cleanupFns = nil
}

// Execute actually executes the functions registered by Add in the reverse
// order of their call order (like normal defer blocks).
func (d *ResettableDefer) Execute() {
	// Run these in reverse order, like defer blocks.
	for i := len(d.cleanupFns) - 1; i >= 0; i-- {
		d.cleanupFns[i]()
	}
}
