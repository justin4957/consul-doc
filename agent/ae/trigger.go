// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/ae/trigger.go
Trigger module for agent layer

## Tags
agent

## Exports
NewTrigger, Trigger

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/ae/trigger.go> a code:Module ;
    code:name "agent/ae/trigger.go" ;
    code:description "Trigger module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :NewTrigger, :Trigger ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package ae

// Trigger implements a non-blocking event notifier. Events can be
// triggered without blocking and notifications happen only when the
// previous event was consumed.
type Trigger struct {
	ch chan struct{}
}

func NewTrigger() *Trigger {
	return &Trigger{make(chan struct{}, 1)}
}

func (t Trigger) Trigger() {
	select {
	case t.ch <- struct{}{}:
	default:
	}
}

func (t Trigger) Notif() <-chan struct{} {
	return t.ch
}
