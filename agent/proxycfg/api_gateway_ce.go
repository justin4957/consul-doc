// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/api_gateway_ce.go
Api Gateway Ce module for agent layer

## Tags
agent, api, client, networking, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/api_gateway_ce.go> a code:Module ;
    code:name "agent/proxycfg/api_gateway_ce.go" ;
    code:description "Api Gateway Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "api", "client", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

import "context"

func watchJWTProviders(cxt context.Context, h *handlerAPIGateway) error {
	return nil
}

func setJWTProvider(u UpdateEvent, snap *ConfigSnapshot) error {
	return nil
}
