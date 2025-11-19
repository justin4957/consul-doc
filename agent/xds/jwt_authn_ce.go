// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/xds/jwt_authn_ce.go
Jwt Authn Ce module for agent layer

## Linked Modules
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)

## Tags
agent, authentication, envoy, security, service-mesh

## Exports
GatewayAuthFilterBuilder

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/jwt_authn_ce.go> a code:Module ;
    code:name "agent/xds/jwt_authn_ce.go" ;
    code:description "Jwt Authn Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :GatewayAuthFilterBuilder ;
    code:tags "agent", "authentication", "envoy", "security", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	envoy_http_jwt_authn_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"
	envoy_http_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/go-hclog"
)

type GatewayAuthFilterBuilder struct {
	listener       structs.APIGatewayListener
	routes         []*structs.HTTPRouteConfigEntry
	providers      map[string]*structs.JWTProviderConfigEntry
	envoyProviders map[string]*envoy_http_jwt_authn_v3.JwtProvider
}

func (g *GatewayAuthFilterBuilder) makeGatewayAuthFilters() ([]*envoy_http_v3.HttpFilter, error) {
	return nil, nil
}

func makeAPIGatewayJWKClusters(_ hclog.Logger, _ *proxycfg.ConfigSnapshot) []proto.Message {
	return nil
}
