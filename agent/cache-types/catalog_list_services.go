// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/catalog_list_services.go
Catalog List Services module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/structs](../agent/structs)

## Tags
agent, data-model, discovery, registry, types

## Exports
CatalogListServices, CatalogListServicesName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/catalog_list_services.go> a code:Module ;
    code:name "agent/cache-types/catalog_list_services.go" ;
    code:description "Catalog List Services module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :CatalogListServices, :CatalogListServicesName ;
    code:tags "agent", "data-model", "discovery", "registry", "types" .
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
const CatalogListServicesName = "catalog-list-services"

// CatalogListServices supports fetching discovering service names via the catalog.
type CatalogListServices struct {
	RegisterOptionsBlockingRefresh
	RPC RPC
}

func (c *CatalogListServices) Fetch(opts cache.FetchOptions, req cache.Request) (cache.FetchResult, error) {
	var result cache.FetchResult

	// The request should be a DCSpecificRequest.
	reqReal, ok := req.(*structs.DCSpecificRequest)
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

	if opts.LastResult != nil {
		reqReal.AllowNotModifiedResponse = true
	}

	var reply structs.IndexedServices
	if err := c.RPC.RPC(context.Background(), "Catalog.ListServices", reqReal, &reply); err != nil {
		return result, err
	}

	result.Value = &reply
	result.Index = reply.Index
	result.NotModified = reply.NotModified
	return result, nil
}
