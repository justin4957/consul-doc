// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/channels/deliver_latest.go
Deliver Latest module for internal layer

## Tags
internal

## Exports
DeliverLatest

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/channels/deliver_latest.go> a code:Module ;
    code:name "lib/channels/deliver_latest.go" ;
    code:description "Deliver Latest module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :DeliverLatest ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package channels

import "fmt"

// DeliverLatest will drain the channel discarding any messages if there are any and sends the current message.
func DeliverLatest[T any](val T, ch chan T) error {
	// Send if chan is empty
	select {
	case ch <- val:
		return nil
	default:
	}

	// If it falls through to here, the channel is not empty.
	// Drain the channel.
	done := false
	for !done {
		select {
		case <-ch:
			continue
		default:
			done = true
		}
	}

	// Attempt to send again.  If it is not empty, throw an error
	select {
	case ch <- val:
		return nil
	default:
		return fmt.Errorf("failed to deliver latest event: chan full again after draining")
	}
}
