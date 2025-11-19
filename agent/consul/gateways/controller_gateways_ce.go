// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/gateways/controller_gateways_ce.go
Controller Gateways Ce module for agent layer

## Linked Modules
- [agent/consul/controller](../agent/consul/controller)
- [agent/consul/state](../agent/consul/state)
- [agent/structs](../agent/structs)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/gateways/controller_gateways_ce.go> a code:Module ;
    code:name "agent/consul/gateways/controller_gateways_ce.go" ;
    code:description "Controller Gateways Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/controller" ;
        code:path "../agent/consul/controller"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package gateways

import (
	"context"

	"github.com/hashicorp/consul/agent/consul/controller"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/structs"
)

func (r *apiGatewayReconciler) enqueueJWTProviderReferencedGatewaysAndHTTPRoutes(_ *state.Store, _ context.Context, _ controller.Request) error {
	return nil
}

func (m *gatewayMeta) checkJWTProviders() (map[structs.ResourceReference]error, error) {
	return nil, nil
}

func (m *gatewayMeta) validateJWTForRoute(_ *structs.HTTPRouteConfigEntry) (bool, map[structs.ResourceReference]error) {
	return true, nil
}
