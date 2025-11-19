// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/gateway_services.go
Gateway Services module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [agent/structs/aclfilter](../agent/structs/aclfilter)

## Tags
agent, networking, service-mesh

## Exports
CacheGatewayServices, ServerGatewayServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/gateway_services.go> a code:Module ;
    code:name "agent/proxycfg-glue/gateway_services.go" ;
    code:description "Gateway Services module for agent layer" ;
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
        code:name "agent/consul/watch" ;
        code:path "../agent/consul/watch"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "agent/structs/aclfilter" ;
        code:path "../agent/structs/aclfilter"
    ] ;
    code:exports :CacheGatewayServices, :ServerGatewayServices ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/agent/structs/aclfilter"
)

// CacheGatewayServices satisfies the proxycfg.GatewayServices interface by
// sourcing data from the agent cache.
func CacheGatewayServices(c *cache.Cache) proxycfg.GatewayServices {
	return &cacheProxyDataSource[*structs.ServiceSpecificRequest]{c, cachetype.GatewayServicesName}
}

// ServerGatewayServices satisfies the proxycfg.GatewayServices interface by
// sourcing data from a blocking query against the server's state store.
func ServerGatewayServices(deps ServerDataSourceDeps) proxycfg.GatewayServices {
	return &serverGatewayServices{deps}
}

type serverGatewayServices struct {
	deps ServerDataSourceDeps
}

func (s *serverGatewayServices) Notify(ctx context.Context, req *structs.ServiceSpecificRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *structs.IndexedGatewayServices, error) {
			var authzContext acl.AuthorizerContext
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &req.EnterpriseMeta, &authzContext)
			if err != nil {
				return 0, nil, err
			}
			if err := authz.ToAllowAuthorizer().ServiceReadAllowed(req.ServiceName, &authzContext); err != nil {
				return 0, nil, err
			}

			index, services, err := store.GatewayServices(ws, req.ServiceName, &req.EnterpriseMeta)
			if err != nil {
				return 0, nil, err
			}

			response := &structs.IndexedGatewayServices{
				Services: services,
				QueryMeta: structs.QueryMeta{
					Backend: structs.QueryBackendBlocking,
					Index:   index,
				},
			}
			aclfilter.New(authz, s.deps.Logger).Filter(response)

			return index, response, nil
		},
		dispatchBlockingQueryUpdate[*structs.IndexedGatewayServices](ch),
	)
}
