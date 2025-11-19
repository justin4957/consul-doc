// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/rpc.go
Rpc module for agent layer

## Tags
agent, communication, data-model, networking, types

## Exports
RPC

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/rpc.go> a code:Module ;
    code:name "agent/cache-types/rpc.go" ;
    code:description "Rpc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :RPC ;
    code:tags "agent", "communication", "data-model", "networking", "types" .
<!-- End LinkedDoc RDF -->
*/
package cachetype

import "context"

// RPC is an interface that an RPC client must implement. This is a helper
// interface that is implemented by the agent delegate so that Type
// implementations can request RPC access.
//
//go:generate mockery --name RPC --inpackage
type RPC interface {
	RPC(ctx context.Context, method string, args interface{}, reply interface{}) error
}
