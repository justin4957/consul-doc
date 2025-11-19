// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/service_http_checks.go
Service Http Checks module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/local](../agent/local)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)

## Tags
agent, health-checks, monitoring, networking, service-mesh

## Exports
CacheHTTPChecks, ServerHTTPChecks

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/service_http_checks.go> a code:Module ;
    code:name "agent/proxycfg-glue/service_http_checks.go" ;
    code:description "Service Http Checks module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/cache-types" ;
        code:path "../agent/cache-types"
    ], [
        code:name "agent/local" ;
        code:path "../agent/local"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :CacheHTTPChecks, :ServerHTTPChecks ;
    code:tags "agent", "health-checks", "monitoring", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/local"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
)

// CacheHTTPChecks satisifies the proxycfg.HTTPChecks interface by sourcing
// data from the agent cache.
func CacheHTTPChecks(c *cache.Cache) proxycfg.HTTPChecks {
	return &cacheProxyDataSource[*cachetype.ServiceHTTPChecksRequest]{c, cachetype.ServiceHTTPChecksName}
}

// ServerHTTPChecks satisifies the proxycfg.HTTPChecks interface.
// It sources data from the server agent cache if the service exists in the local state, else
// it is a no-op since the checks can only be performed by local agents.
func ServerHTTPChecks(deps ServerDataSourceDeps, nodeName string, cacheSource proxycfg.HTTPChecks, localState *local.State) proxycfg.HTTPChecks {
	return serverHTTPChecks{deps, nodeName, cacheSource, localState}
}

type serverHTTPChecks struct {
	deps        ServerDataSourceDeps
	nodeName    string
	cacheSource proxycfg.HTTPChecks
	localState  *local.State
}

func (c serverHTTPChecks) Notify(ctx context.Context, req *cachetype.ServiceHTTPChecksRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	// if the service exists in the current server agent local state, delegate to the cache data source
	if req.NodeName == c.nodeName && c.localState.ServiceExists(structs.ServiceID{ID: req.ServiceID, EnterpriseMeta: req.EnterpriseMeta}) {
		return c.cacheSource.Notify(ctx, req, correlationID, ch)
	}
	l := c.deps.Logger.With("service_id", req.ServiceID)
	l.Debug("service-http-checks: no-op")
	return nil
}
