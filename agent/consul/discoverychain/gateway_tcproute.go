// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/discoverychain/gateway_tcproute.go
Gateway Tcproute module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, discovery, dns

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/discoverychain/gateway_tcproute.go> a code:Module ;
    code:name "agent/consul/discoverychain/gateway_tcproute.go" ;
    code:description "Gateway Tcproute module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "discovery", "dns" .
<!-- End LinkedDoc RDF -->
*/
package discoverychain

import "github.com/hashicorp/consul/agent/structs"

func synthesizeTCPRouteDiscoveryChain(route structs.TCPRouteConfigEntry) []structs.IngressService {
	services := make([]structs.IngressService, 0, len(route.Services))
	for _, service := range route.Services {
		ingress := structs.IngressService{
			Name:           service.Name,
			EnterpriseMeta: service.EnterpriseMeta,
		}

		services = append(services, ingress)
	}

	return services
}
