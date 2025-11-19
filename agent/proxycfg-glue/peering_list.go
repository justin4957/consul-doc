// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/peering_list.go
Peering List module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [proto/private/pbpeering](../proto/private/pbpeering)

## Tags
agent, networking, service-mesh

## Exports
CachePeeringList, ServerPeeringList

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/peering_list.go> a code:Module ;
    code:name "agent/proxycfg-glue/peering_list.go" ;
    code:description "Peering List module for agent layer" ;
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
        code:name "proto/private/pbpeering" ;
        code:path "../proto/private/pbpeering"
    ] ;
    code:exports :CachePeeringList, :ServerPeeringList ;
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
	"github.com/hashicorp/consul/proto/private/pbpeering"
)

// CachePeeringList satisfies the proxycfg.PeeringList interface by sourcing
// data from the agent cache.
func CachePeeringList(c *cache.Cache) proxycfg.PeeringList {
	return &cacheProxyDataSource[*cachetype.PeeringListRequest]{c, cachetype.PeeringListName}
}

// ServerPeeringList satisfies the proxycfg.PeeringList interface by sourcing
// data from a blocking query against the server's state store.
func ServerPeeringList(deps ServerDataSourceDeps) proxycfg.PeeringList {
	return &serverPeeringList{deps}
}

type serverPeeringList struct {
	deps ServerDataSourceDeps
}

func (s *serverPeeringList) Notify(ctx context.Context, req *cachetype.PeeringListRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	entMeta := structs.DefaultEnterpriseMetaInPartition(req.Request.Partition)

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *pbpeering.PeeringListResponse, error) {
			var authzCtx acl.AuthorizerContext
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, entMeta, &authzCtx)
			if err != nil {
				return 0, nil, err
			}
			if err := authz.ToAllowAuthorizer().PeeringReadAllowed(&authzCtx); err != nil {
				return 0, nil, err
			}

			index, peerings, err := store.PeeringList(ws, *entMeta)
			if err != nil {
				return 0, nil, err
			}
			return index, &pbpeering.PeeringListResponse{
				OBSOLETE_Index: index,
				Peerings:       peerings,
			}, nil
		},
		dispatchBlockingQueryUpdate[*pbpeering.PeeringListResponse](ch),
	)
}
