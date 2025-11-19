// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/intention_upstreams_destination.go
Intention Upstreams Destination module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/structs](../agent/structs)

## Tags
agent, data-model, types

## Exports
IntentionUpstreamsDestination, IntentionUpstreamsDestinationName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/intention_upstreams_destination.go> a code:Module ;
    code:name "agent/cache-types/intention_upstreams_destination.go" ;
    code:description "Intention Upstreams Destination module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :IntentionUpstreamsDestination, :IntentionUpstreamsDestinationName ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package cachetype

import (
	"context"
	"fmt"

	"github.com/hashicorp/consul/agent/cache"
	"github.com/hashicorp/consul/agent/structs"
)

// IntentionUpstreamsDestinationName Recommended name for registration.
const IntentionUpstreamsDestinationName = "intention-upstreams-destination"

// IntentionUpstreamsDestination supports fetching upstreams for a given gateway name.
type IntentionUpstreamsDestination struct {
	RegisterOptionsBlockingRefresh
	RPC RPC
}

func (i *IntentionUpstreamsDestination) Fetch(opts cache.FetchOptions, req cache.Request) (cache.FetchResult, error) {
	var result cache.FetchResult

	// The request should be a ServiceSpecificRequest.
	reqReal, ok := req.(*structs.ServiceSpecificRequest)
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
	// going to be served from cache and end up arbitrarily stale anyway. This
	// allows cached service-discover to automatically read scale across all
	// servers too.
	reqReal.AllowStale = true

	// Fetch
	var reply structs.IndexedServiceList
	if err := i.RPC.RPC(context.Background(), "Internal.IntentionUpstreamsDestination", reqReal, &reply); err != nil {
		return result, err
	}

	result.Value = &reply
	result.Index = reply.Index
	return result, nil
}
