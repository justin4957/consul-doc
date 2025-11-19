// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/federation_state_list_mesh_gateways.go
Federation State List Mesh Gateways module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [agent/structs/aclfilter](../agent/structs/aclfilter)

## Tags
agent, networking, persistence, service-mesh, storage

## Exports
CacheFederationStateListMeshGateways, ServerFederationStateListMeshGateways

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/federation_state_list_mesh_gateways.go> a code:Module ;
    code:name "agent/proxycfg-glue/federation_state_list_mesh_gateways.go" ;
    code:description "Federation State List Mesh Gateways module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
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
    code:exports :CacheFederationStateListMeshGateways, :ServerFederationStateListMeshGateways ;
    code:tags "agent", "networking", "persistence", "service-mesh", "storage" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/agent/structs/aclfilter"
)

// CacheFederationStateListMeshGateways satisfies the proxycfg.FederationStateListMeshGateways
// interface by sourcing data from the agent cache.
func CacheFederationStateListMeshGateways(c *cache.Cache) proxycfg.FederationStateListMeshGateways {
	return &cacheProxyDataSource[*structs.DCSpecificRequest]{c, cachetype.FederationStateListMeshGatewaysName}
}

// ServerFederationStateListMeshGateways satisfies the proxycfg.FederationStateListMeshGateways
// interface by sourcing data from a blocking query against the server's state
// store.
func ServerFederationStateListMeshGateways(deps ServerDataSourceDeps) proxycfg.FederationStateListMeshGateways {
	return &serverFederationStateListMeshGateways{deps}
}

type serverFederationStateListMeshGateways struct {
	deps ServerDataSourceDeps
}

func (s *serverFederationStateListMeshGateways) Notify(ctx context.Context, req *structs.DCSpecificRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *structs.DatacenterIndexedCheckServiceNodes, error) {
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &req.EnterpriseMeta, nil)
			if err != nil {
				return 0, nil, err
			}

			index, fedStates, err := store.FederationStateList(ws)
			if err != nil {
				return 0, nil, err
			}

			results := make(map[string]structs.CheckServiceNodes)
			for _, fs := range fedStates {
				if gws := fs.MeshGateways; len(gws) != 0 {
					// Shallow clone to prevent ACL filtering manipulating the slice in memdb.
					results[fs.Datacenter] = gws.ShallowClone()
				}
			}

			rsp := &structs.DatacenterIndexedCheckServiceNodes{
				DatacenterNodes: results,
				QueryMeta: structs.QueryMeta{
					Index:   index,
					Backend: structs.QueryBackendBlocking,
				},
			}
			aclfilter.New(authz, s.deps.Logger).Filter(rsp)

			return index, rsp, nil
		},
		dispatchBlockingQueryUpdate[*structs.DatacenterIndexedCheckServiceNodes](ch),
	)
}
