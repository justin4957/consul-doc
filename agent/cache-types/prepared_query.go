// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/prepared_query.go
Prepared Query module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/structs](../agent/structs)

## Tags
agent, data-model, types

## Exports
PreparedQuery, PreparedQueryName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/prepared_query.go> a code:Module ;
    code:name "agent/cache-types/prepared_query.go" ;
    code:description "Prepared Query module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :PreparedQuery, :PreparedQueryName ;
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

// Recommended name for registration.
const PreparedQueryName = "prepared-query"

// PreparedQuery supports fetching discovering service instances via prepared
// queries.
type PreparedQuery struct {
	RegisterOptionsNoRefresh
	RPC RPC
}

func (c *PreparedQuery) Fetch(_ cache.FetchOptions, req cache.Request) (cache.FetchResult, error) {
	var result cache.FetchResult

	// The request should be a PreparedQueryExecuteRequest.
	reqReal, ok := req.(*structs.PreparedQueryExecuteRequest)
	if !ok {
		return result, fmt.Errorf(
			"Internal cache failure: request wrong type: %T", req)
	}

	// Lightweight copy this object so that manipulating QueryOptions doesn't race.
	dup := *reqReal
	reqReal = &dup

	// Always allow stale - there's no point in hitting leader if the request is
	// going to be served from cache and endup arbitrarily stale anyway. This
	// allows cached service-discover to automatically read scale across all
	// servers too.
	reqReal.AllowStale = true

	// Fetch
	var reply structs.PreparedQueryExecuteResponse
	if err := c.RPC.RPC(context.Background(), "PreparedQuery.Execute", reqReal, &reply); err != nil {
		return result, err
	}

	result.Value = &reply
	result.Index = reply.Index

	return result, nil
}
