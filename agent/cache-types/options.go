// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/cache-types/options.go
Options module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)

## Tags
agent, data-model, types

## Exports
RegisterOptionsBlockingNoRefresh, RegisterOptionsBlockingRefresh, RegisterOptionsNoRefresh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/cache-types/options.go> a code:Module ;
    code:name "agent/cache-types/options.go" ;
    code:description "Options module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ] ;
    code:exports :RegisterOptionsBlockingNoRefresh, :RegisterOptionsBlockingRefresh, :RegisterOptionsNoRefresh ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package cachetype

import (
	"time"

	"github.com/hashicorp/consul/agent/cache"
)

// RegisterOptionsBlockingRefresh can be embedded into a struct to implement
// part of the agent/cache.Type interface.
// When embedded into a struct it identifies the cache type as one which
// supports blocking, and uses refresh to keep the cache fresh.
type RegisterOptionsBlockingRefresh struct{}

func (r RegisterOptionsBlockingRefresh) RegisterOptions() cache.RegisterOptions {
	return cache.RegisterOptions{
		// Maintain a blocking query, retry dropped connections quickly
		Refresh:          true,
		SupportsBlocking: true,
		RefreshTimer:     0 * time.Second,
		QueryTimeout:     10 * time.Minute,
	}
}

type RegisterOptionsBlockingNoRefresh struct{}

func (r RegisterOptionsBlockingNoRefresh) RegisterOptions() cache.RegisterOptions {
	return cache.RegisterOptions{
		Refresh:          false,
		SupportsBlocking: true,
		QueryTimeout:     10 * time.Minute,
	}
}

// RegisterOptionsNoRefresh can be embedded into a struct to implement
// part of the agent/cache.Type interface.
// When embedded into a struct it identifies the cache type as one which
// does not support blocking, and should not be refreshed.
type RegisterOptionsNoRefresh struct{}

func (r RegisterOptionsNoRefresh) RegisterOptions() cache.RegisterOptions {
	return cache.RegisterOptions{
		Refresh:          false,
		SupportsBlocking: false,
	}
}
