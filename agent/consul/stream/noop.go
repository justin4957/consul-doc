// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/stream/noop.go
Noop module for agent layer

## Tags
agent

## Exports
NoOpEventPublisher

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/stream/noop.go> a code:Module ;
    code:name "agent/consul/stream/noop.go" ;
    code:description "Noop module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :NoOpEventPublisher ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package stream

import (
	"context"
	"fmt"
)

type NoOpEventPublisher struct{}

func (NoOpEventPublisher) Publish([]Event) {}

func (NoOpEventPublisher) RegisterHandler(Topic, SnapshotFunc, bool) error {
	return fmt.Errorf("stream event publisher is disabled")
}

func (NoOpEventPublisher) Run(context.Context) {}

func (NoOpEventPublisher) Subscribe(*SubscribeRequest) (*Subscription, error) {
	return nil, fmt.Errorf("stream event publisher is disabled")
}
