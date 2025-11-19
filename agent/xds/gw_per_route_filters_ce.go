// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/xds/gw_per_route_filters_ce.go
Gw Per Route Filters Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/gw_per_route_filters_ce.go> a code:Module ;
    code:name "agent/xds/gw_per_route_filters_ce.go" ;
    code:description "Gw Per Route Filters Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	envoy_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/hashicorp/consul/agent/structs"
)

type perRouteFilterBuilder struct {
	providerMap map[string]*structs.JWTProviderConfigEntry
	listener    *structs.APIGatewayListener
	route       *structs.HTTPRouteConfigEntry
}

func (p perRouteFilterBuilder) buildTypedPerFilterConfig(match *envoy_route_v3.RouteMatch, routeAction *envoy_route_v3.Route_Route) (map[string]*anypb.Any, error) {
	return nil, nil
}
