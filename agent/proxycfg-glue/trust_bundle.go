// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/trust_bundle.go
Trust Bundle module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/consul/state](../agent/consul/state)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [proto/private/pbpeering](../proto/private/pbpeering)

## Tags
agent, networking, service-mesh

## Exports
CacheTrustBundle, CacheTrustBundleList, ServerTrustBundle, ServerTrustBundleList

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/trust_bundle.go> a code:Module ;
    code:name "agent/proxycfg-glue/trust_bundle.go" ;
    code:description "Trust Bundle module for agent layer" ;
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
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
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
    code:exports :CacheTrustBundle, :CacheTrustBundleList, :ServerTrustBundle, :ServerTrustBundleList ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"
	"errors"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbpeering"
)

// CacheTrustBundle satisfies the proxycfg.TrustBundle interface by sourcing
// data from the agent cache.
func CacheTrustBundle(c *cache.Cache) proxycfg.TrustBundle {
	return &cacheProxyDataSource[*cachetype.TrustBundleReadRequest]{c, cachetype.TrustBundleReadName}
}

// ServerTrustBundle satisfies the proxycfg.TrustBundle interface by sourcing
// data from a blocking query against the server's state store.
func ServerTrustBundle(deps ServerDataSourceDeps) proxycfg.TrustBundle {
	return &serverTrustBundle{deps}
}

type serverTrustBundle struct {
	deps ServerDataSourceDeps
}

func (s *serverTrustBundle) Notify(ctx context.Context, req *cachetype.TrustBundleReadRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	// Having the ability to write a service in ANY (at least one) namespace should be
	// sufficient for reading the trust bundle, which is why we use a wildcard.
	entMeta := acl.NewEnterpriseMetaWithPartition(req.Request.Partition, acl.WildcardName)

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *pbpeering.TrustBundleReadResponse, error) {
			var authzCtx acl.AuthorizerContext
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &entMeta, &authzCtx)
			if err != nil {
				return 0, nil, err
			}
			if err := authz.ToAllowAuthorizer().ServiceWriteAnyAllowed(&authzCtx); err != nil {
				return 0, nil, err
			}

			index, bundle, err := store.PeeringTrustBundleRead(ws, state.Query{
				Value:          req.Request.Name,
				EnterpriseMeta: *structs.NodeEnterpriseMetaInPartition(req.Request.Partition),
			})
			if err != nil {
				return 0, nil, err
			}
			return index, &pbpeering.TrustBundleReadResponse{
				OBSOLETE_Index: index,
				Bundle:         bundle,
			}, nil
		},
		dispatchBlockingQueryUpdate[*pbpeering.TrustBundleReadResponse](ch),
	)
}

// CacheTrustBundleList satisfies the proxycfg.TrustBundleList interface by sourcing
// data from the agent cache.
func CacheTrustBundleList(c *cache.Cache) proxycfg.TrustBundleList {
	return &cacheProxyDataSource[*cachetype.TrustBundleListRequest]{c, cachetype.TrustBundleListName}
}

// ServerTrustBundleList satisfies the proxycfg.TrustBundle interface by
// sourcing data from a blocking query against the server's state store.
func ServerTrustBundleList(deps ServerDataSourceDeps) proxycfg.TrustBundleList {
	return &serverTrustBundleList{deps}
}

type serverTrustBundleList struct {
	deps ServerDataSourceDeps
}

func (s *serverTrustBundleList) Notify(ctx context.Context, req *cachetype.TrustBundleListRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	entMeta := acl.NewEnterpriseMetaWithPartition(req.Request.Partition, req.Request.Namespace)

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *pbpeering.TrustBundleListByServiceResponse, error) {
			var authzCtx acl.AuthorizerContext
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &entMeta, &authzCtx)
			if err != nil {
				return 0, nil, err
			}
			if err := authz.ToAllowAuthorizer().ServiceWriteAllowed(req.Request.ServiceName, &authzCtx); err != nil {
				return 0, nil, err
			}

			var (
				index   uint64
				bundles []*pbpeering.PeeringTrustBundle
			)
			switch {
			case req.Request.Kind == string(structs.ServiceKindMeshGateway):
				index, bundles, err = store.PeeringTrustBundleList(ws, entMeta)
			case req.Request.ServiceName != "":
				index, bundles, err = store.TrustBundleListByService(ws, req.Request.ServiceName, s.deps.Datacenter, entMeta)
			case req.Request.Kind != "":
				err = errors.New("kind must be mesh-gateway if set")
			default:
				err = errors.New("one of service or kind is required")
			}
			if err != nil {
				return 0, nil, err
			}

			return index, &pbpeering.TrustBundleListByServiceResponse{
				OBSOLETE_Index: index,
				Bundles:        bundles,
			}, nil
		},
		dispatchBlockingQueryUpdate[*pbpeering.TrustBundleListByServiceResponse](ch),
	)
}
