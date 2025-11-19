// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/stop_context.go
Stop Context module for internal layer

## Tags
internal

## Exports
StopChannelContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/stop_context.go> a code:Module ;
    code:name "lib/stop_context.go" ;
    code:description "Stop Context module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :StopChannelContext ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package lib

import (
	"context"
	"time"
)

// StopChannelContext implements the context.Context interface
// You provide the channel to select on to determine whether
// the context should be canceled and other code such
// as the rate.Limiter will automatically use the channel
// appropriately
type StopChannelContext struct {
	StopCh <-chan struct{}
}

func (c *StopChannelContext) Deadline() (deadline time.Time, ok bool) {
	ok = false
	return
}

func (c *StopChannelContext) Done() <-chan struct{} {
	return c.StopCh
}

func (c *StopChannelContext) Err() error {
	select {
	case <-c.StopCh:
		return context.Canceled
	default:
		return nil
	}
}

func (c *StopChannelContext) Value(key interface{}) interface{} {
	return nil
}
