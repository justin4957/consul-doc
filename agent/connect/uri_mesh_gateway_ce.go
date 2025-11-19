// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/connect/uri_mesh_gateway_ce.go
Uri Mesh Gateway Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_mesh_gateway_ce.go> a code:Module ;
    code:name "agent/connect/uri_mesh_gateway_ce.go" ;
    code:description "Uri Mesh Gateway Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

// GetEnterpriseMeta will synthesize an EnterpriseMeta struct from the SpiffeIDAgent.
// in CE this just returns an empty (but never nil) struct pointer
func (id SpiffeIDMeshGateway) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return &acl.EnterpriseMeta{}
}

func (id SpiffeIDMeshGateway) uriPath() string {
	return fmt.Sprintf("/gateway/mesh/dc/%s", id.Datacenter)
}
