// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/resolved_service_config.go
Resolved Service Config module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/configentry](../agent/configentry)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, networking, service-mesh

## Exports
CacheResolvedServiceConfig, ServerResolvedServiceConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/resolved_service_config.go> a code:Module ;
    code:name "agent/proxycfg-glue/resolved_service_config.go" ;
    code:description "Resolved Service Config module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/cache-types" ;
        code:path "../agent/cache-types"
    ], [
        code:name "agent/configentry" ;
        code:path "../agent/configentry"
    ], [
        code:name "agent/consul/watch" ;
        code:path "../agent/consul/watch"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :CacheResolvedServiceConfig, :ServerResolvedServiceConfig ;
    code:tags "agent", "configuration", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
)

// CacheResolvedServiceConfig satisfies the proxycfg.ResolvedServiceConfig
// interface by sourcing data from the agent cache.
func CacheResolvedServiceConfig(c *cache.Cache) proxycfg.ResolvedServiceConfig {
	return &cacheProxyDataSource[*structs.ServiceConfigRequest]{c, cachetype.ResolvedServiceConfigName}
}

// ServerResolvedServiceConfig satisfies the proxycfg.ResolvedServiceConfig
// interface by sourcing data from a blocking query against the server's state
// store.
func ServerResolvedServiceConfig(deps ServerDataSourceDeps, remoteSource proxycfg.ResolvedServiceConfig) proxycfg.ResolvedServiceConfig {
	return &serverResolvedServiceConfig{deps, remoteSource}
}

type serverResolvedServiceConfig struct {
	deps         ServerDataSourceDeps
	remoteSource proxycfg.ResolvedServiceConfig
}

func (s *serverResolvedServiceConfig) Notify(ctx context.Context, req *structs.ServiceConfigRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	if req.Datacenter != s.deps.Datacenter {
		return s.remoteSource.Notify(ctx, req, correlationID, ch)
	}

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *structs.ServiceConfigResponse, error) {
			var authzContext acl.AuthorizerContext
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &req.EnterpriseMeta, &authzContext)
			if err != nil {
				return 0, nil, err
			}

			if err := authz.ToAllowAuthorizer().ServiceReadAllowed(req.Name, &authzContext); err != nil {
				return 0, nil, err
			}

			idx, entries, err := store.ReadResolvedServiceConfigEntries(ws, req.Name, &req.EnterpriseMeta, req.GetLocalUpstreamIDs(), req.Mode)
			if err != nil {
				return 0, nil, err
			}

			reply, err := configentry.ComputeResolvedServiceConfig(req, entries, s.deps.Logger)
			if err != nil {
				return 0, nil, err
			}
			reply.Index = idx

			return idx, reply, nil
		},
		dispatchBlockingQueryUpdate[*structs.ServiceConfigResponse](ch),
	)
}
