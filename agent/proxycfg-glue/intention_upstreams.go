// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/intention_upstreams.go
Intention Upstreams module for agent layer

## Linked Modules
- [agent/cache](../agent/cache)
- [agent/cache-types](../agent/cache-types)
- [agent/consul](../agent/consul)
- [agent/consul/watch](../agent/consul/watch)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [agent/structs/aclfilter](../agent/structs/aclfilter)

## Tags
agent, networking, service-mesh

## Exports
CacheIntentionUpstreams, CacheIntentionUpstreamsDestination, ServerIntentionUpstreams, ServerIntentionUpstreamsDestination

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/intention_upstreams.go> a code:Module ;
    code:name "agent/proxycfg-glue/intention_upstreams.go" ;
    code:description "Intention Upstreams module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/cache" ;
        code:path "../agent/cache"
    ], [
        code:name "agent/cache-types" ;
        code:path "../agent/cache-types"
    ], [
        code:name "agent/consul" ;
        code:path "../agent/consul"
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
    code:exports :CacheIntentionUpstreams, :CacheIntentionUpstreamsDestination, :ServerIntentionUpstreams, :ServerIntentionUpstreamsDestination ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/cache"
	cachetype "github.com/hashicorp/consul/agent/cache-types"
	"github.com/hashicorp/consul/agent/consul"
	"github.com/hashicorp/consul/agent/consul/watch"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/agent/structs/aclfilter"
)

// CacheIntentionUpstreams satisfies the proxycfg.IntentionUpstreams interface
// by sourcing upstreams for the given service, inferred from intentions, from
// the agent cache.
func CacheIntentionUpstreams(c *cache.Cache) proxycfg.IntentionUpstreams {
	return &cacheProxyDataSource[*structs.ServiceSpecificRequest]{c, cachetype.IntentionUpstreamsName}
}

// CacheIntentionUpstreamsDestination satisfies the proxycfg.IntentionUpstreams
// interface by sourcing upstreams for the given destination, inferred from
// intentions, from the agent cache.
func CacheIntentionUpstreamsDestination(c *cache.Cache) proxycfg.IntentionUpstreams {
	return &cacheProxyDataSource[*structs.ServiceSpecificRequest]{c, cachetype.IntentionUpstreamsDestinationName}
}

// ServerIntentionUpstreams satisfies the proxycfg.IntentionUpstreams interface
// by sourcing upstreams for the given service, inferred from intentions, from
// the server's state store.
func ServerIntentionUpstreams(deps ServerDataSourceDeps, defaultIntentionPolicy string) proxycfg.IntentionUpstreams {
	return serverIntentionUpstreams{deps, structs.IntentionTargetService, defaultIntentionPolicy}
}

// ServerIntentionUpstreamsDestination satisfies the proxycfg.IntentionUpstreams
// interface by sourcing upstreams for the given destination, inferred from
// intentions, from the server's state store.
func ServerIntentionUpstreamsDestination(deps ServerDataSourceDeps, defaultIntentionPolicy string) proxycfg.IntentionUpstreams {
	return serverIntentionUpstreams{deps, structs.IntentionTargetDestination, defaultIntentionPolicy}
}

type serverIntentionUpstreams struct {
	deps                   ServerDataSourceDeps
	target                 structs.IntentionTargetType
	defaultIntentionPolicy string
}

func (s serverIntentionUpstreams) Notify(ctx context.Context, req *structs.ServiceSpecificRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	target := structs.NewServiceName(req.ServiceName, &req.EnterpriseMeta)

	return watch.ServerLocalNotify(ctx, correlationID, s.deps.GetStore,
		func(ws memdb.WatchSet, store Store) (uint64, *structs.IndexedServiceList, error) {
			authz, err := s.deps.ACLResolver.ResolveTokenAndDefaultMeta(req.Token, &req.EnterpriseMeta, nil)
			if err != nil {
				return 0, nil, err
			}

			defaultAllow := consul.DefaultIntentionAllow(authz, s.defaultIntentionPolicy)

			index, services, err := store.IntentionTopology(ws, target, false, defaultAllow, s.target)
			if err != nil {
				return 0, nil, err
			}

			result := &structs.IndexedServiceList{
				Services: services,
				QueryMeta: structs.QueryMeta{
					Index:   index,
					Backend: structs.QueryBackendBlocking,
				},
			}
			aclfilter.New(authz, s.deps.Logger).Filter(result)

			return index, result, nil
		},
		dispatchBlockingQueryUpdate[*structs.IndexedServiceList](ch),
	)
}
