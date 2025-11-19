// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/mesh_gateway_ce.go
Mesh Gateway Ce module for agent layer

## Tags
agent, networking, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/mesh_gateway_ce.go> a code:Module ;
    code:name "agent/proxycfg/mesh_gateway_ce.go" ;
    code:description "Mesh Gateway Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

import (
	"context"

	"github.com/hashicorp/go-hclog"
)

func (s *handlerMeshGateway) initializeEntWatches(_ context.Context) error {
	return nil
}

func (s *handlerMeshGateway) handleEntUpdate(_ hclog.Logger, _ context.Context, _ UpdateEvent, _ *ConfigSnapshot) error {
	return nil
}
