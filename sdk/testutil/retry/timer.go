// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: sdk/testutil/retry/timer.go
Timer module for internal layer

## Tags
internal

## Exports
Timer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/testutil/retry/timer.go> a code:Module ;
    code:name "sdk/testutil/retry/timer.go" ;
    code:description "Timer module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Timer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package retry

import "time"

// Timer repeats an operation for a given amount
// of time and waits between subsequent operations.
type Timer struct {
	Timeout time.Duration
	Wait    time.Duration

	// stop is the timeout deadline.
	// TODO: Next()?
	// Set on the first invocation of Next().
	stop time.Time
}

func (r *Timer) Continue() bool {
	if r.stop.IsZero() {
		r.stop = time.Now().Add(r.Timeout)
		return true
	}
	if time.Now().After(r.stop) {
		return false
	}
	time.Sleep(r.Wait)
	return true
}
