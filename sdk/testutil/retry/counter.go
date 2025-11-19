// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: sdk/testutil/retry/counter.go
Counter module for internal layer

## Tags
internal

## Exports
Counter

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/testutil/retry/counter.go> a code:Module ;
    code:name "sdk/testutil/retry/counter.go" ;
    code:description "Counter module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Counter ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package retry

import "time"

// Counter repeats an operation a given number of
// times and waits between subsequent operations.
type Counter struct {
	Count int
	Wait  time.Duration

	count int
}

func (r *Counter) Continue() bool {
	if r.count == r.Count {
		return false
	}
	if r.count > 0 {
		time.Sleep(r.Wait)
	}
	r.count++
	return true
}
