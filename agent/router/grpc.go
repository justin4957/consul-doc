// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/router/grpc.go
Grpc module for agent layer

## Linked Modules
- [agent/metadata](../agent/metadata)
- [types](../types)

## Tags
agent, api, communication, grpc, networking

## Exports
NoOpServerTracker, Rebalancer, ServerTracker

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/router/grpc.go> a code:Module ;
    code:name "agent/router/grpc.go" ;
    code:description "Grpc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/metadata" ;
        code:path "../agent/metadata"
    ], [
        code:name "types" ;
        code:path "../types"
    ] ;
    code:exports :NoOpServerTracker, :Rebalancer, :ServerTracker ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package router

import (
	"github.com/hashicorp/consul/agent/metadata"
	"github.com/hashicorp/consul/types"
)

// ServerTracker is called when Router is notified of a server being added or
// removed.
type ServerTracker interface {
	NewRebalancer(dc string) func()
	AddServer(types.AreaID, *metadata.Server)
	RemoveServer(types.AreaID, *metadata.Server)
}

// Rebalancer is called periodically to re-order the servers so that the load on the
// servers is evenly balanced.
type Rebalancer func()

// NoOpServerTracker is a ServerTracker that does nothing. Used when gRPC is not
// enabled.
type NoOpServerTracker struct{}

// Rebalance does nothing
func (NoOpServerTracker) NewRebalancer(string) func() {
	return func() {}
}

// AddServer does nothing
func (NoOpServerTracker) AddServer(types.AreaID, *metadata.Server) {}

// RemoveServer does nothing
func (NoOpServerTracker) RemoveServer(types.AreaID, *metadata.Server) {}
