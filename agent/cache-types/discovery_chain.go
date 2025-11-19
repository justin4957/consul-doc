// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/discovery_chain.go
Discovery Chain module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/structs](../agent/structs)

## Tags
agent, data-model, discovery, dns, types

## Exports
CompiledDiscoveryChain, CompiledDiscoveryChainName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/discovery_chain.go> a code:Module ;
    code:name "agent/cache-types/discovery_chain.go" ;
    code:description "Discovery Chain module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :CompiledDiscoveryChain, :CompiledDiscoveryChainName ;
    code:tags "agent", "data-model", "discovery", "dns", "types" .
<!-- End LinkedDoc RDF -->
*/
package cachetype

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/agent/cache"
	"github.com/hashicorp/consul/agent/structs"
)

// Recommended name for registration.
const CompiledDiscoveryChainName = "compiled-discovery-chain"

// CompiledDiscoveryChain supports fetching the complete discovery chain for a
// service and caching its compilation.
type CompiledDiscoveryChain struct {
	RegisterOptionsBlockingRefresh
	RPC RPC
}

func (c *CompiledDiscoveryChain) Fetch(opts cache.FetchOptions, req cache.Request) (cache.FetchResult, error) {
	var result cache.FetchResult

	// The request should be a DiscoveryChainRequest.
	reqReal, ok := req.(*structs.DiscoveryChainRequest)
	if !ok {
		return result, fmt.Errorf(
			"Internal cache failure: request wrong type: %T", req)
	}

	// Lightweight copy this object so that manipulating QueryOptions doesn't race.
	dup := *reqReal
	reqReal = &dup

	// Set the minimum query index to our current index so we block
	reqReal.MinQueryIndex = opts.MinIndex
	reqReal.MaxQueryTime = opts.Timeout

	// Always allow stale - there's no point in hitting leader if the request is
	// going to be served from cache and endup arbitrarily stale anyway. This
	// allows cached compiled-discovery-chain to automatically read scale across all
	// servers too.
	reqReal.AllowStale = true

	// Fetch
	var reply structs.DiscoveryChainResponse
	if err := c.RPC.RPC(context.Background(), "DiscoveryChain.Get", reqReal, &reply); err != nil {
		return result, err
	}

	result.Value = &reply
	result.Index = reply.Index
	return result, nil
}
