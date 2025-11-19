// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-internal/tracker.go
Tracker module for agent layer

## Linked Modules
- [agent/grpc-internal/balancer](../agent/grpc-internal/balancer)
- [agent/grpc-internal/resolver](../agent/grpc-internal/resolver)
- [agent/metadata](../agent/metadata)
- [types](../types)

## Tags
agent, api, communication, grpc, networking

## Exports
NewTracker, Tracker

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-internal/tracker.go> a code:Module ;
    code:name "agent/grpc-internal/tracker.go" ;
    code:description "Tracker module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/grpc-internal/balancer" ;
        code:path "../agent/grpc-internal/balancer"
    ], [
        code:name "agent/grpc-internal/resolver" ;
        code:path "../agent/grpc-internal/resolver"
    ], [
        code:name "agent/metadata" ;
        code:path "../agent/metadata"
    ], [
        code:name "types" ;
        code:path "../types"
    ] ;
    code:exports :NewTracker, :Tracker ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package internal

import (
	"fmt"
	"net/url"

	gresolver "google.golang.org/grpc/resolver"

	"github.com/hashicorp/consul/agent/grpc-internal/balancer"
	"github.com/hashicorp/consul/agent/grpc-internal/resolver"
	"github.com/hashicorp/consul/agent/metadata"
	"github.com/hashicorp/consul/types"
)

// NewTracker returns an implementation of the router.ServerTracker interface
// backed by the given ServerResolverBuilder and Balancer.
func NewTracker(rb *resolver.ServerResolverBuilder, bb *balancer.Builder) *Tracker {
	return &Tracker{rb, bb}
}

// Tracker satisfies the ServerTracker interface the router manager uses to
// register/deregister servers and trigger rebalances.
type Tracker struct {
	rb *resolver.ServerResolverBuilder
	bb *balancer.Builder
}

// AddServer adds the given server to the resolver.
func (t *Tracker) AddServer(a types.AreaID, s *metadata.Server) { t.rb.AddServer(a, s) }

// RemoveServer removes the given server from the resolver.
func (t *Tracker) RemoveServer(a types.AreaID, s *metadata.Server) { t.rb.RemoveServer(a, s) }

// NewRebalancer returns a function that can be called to randomize the
// priority order of connections for the given datacenter.
func (t *Tracker) NewRebalancer(dc string) func() {
	return func() {
		t.bb.Rebalance(gresolver.Target{
			URL: url.URL{
				Scheme: "consul",
				Host:   t.rb.Authority(),
				Path:   fmt.Sprintf("server.%s", dc),
			},
		})
	}
}
