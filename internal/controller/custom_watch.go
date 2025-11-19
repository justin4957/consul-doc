// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/controller/custom_watch.go
Custom Watch module for internal layer

## Linked Modules
- [agent/consul/controller/queue](../agent/consul/controller/queue)

## Tags
internal

## Exports
CustomDependencyMapper, Event, Source

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/controller/custom_watch.go> a code:Module ;
    code:name "internal/controller/custom_watch.go" ;
    code:description "Custom Watch module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "agent/consul/controller/queue" ;
        code:path "../agent/consul/controller/queue"
    ] ;
    code:exports :CustomDependencyMapper, :Event, :Source ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package controller

import (
	"context"

	"github.com/hashicorp/consul/agent/consul/controller/queue"
)

// CustomDependencyMapper is called when an Event occurs to determine which of the
// controller's managed resources need to be reconciled.
type CustomDependencyMapper func(
	ctx context.Context,
	rt Runtime,
	event Event,
) ([]Request, error)

// Watch is responsible for watching for custom events from source and adding them to
// the event queue.
func (s *Source) Watch(ctx context.Context, add func(e Event)) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case evt, ok := <-s.Source:
			if !ok {
				return nil
			}
			add(evt)
		}
	}
}

// Source is used as a generic source of events. This can be used when events aren't coming from resources
// stored by the resource API.
type Source struct {
	Source <-chan Event
}

// Event captures an event in the system which the API can choose to respond to.
type Event struct {
	Obj queue.ItemType
}

// Key returns a string that will be used to de-duplicate items in the queue.
func (e Event) Key() string {
	return e.Obj.Key()
}

// customWatch represent a Watch on a custom Event source and a Mapper to map said
// Events into Requests that the controller can respond to.
type customWatch struct {
	source *Source
	mapper CustomDependencyMapper
}
