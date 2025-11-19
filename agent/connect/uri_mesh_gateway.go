// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/uri_mesh_gateway.go
Uri Mesh Gateway module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

## Exports
SpiffeIDMeshGateway

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_mesh_gateway.go> a code:Module ;
    code:name "agent/connect/uri_mesh_gateway.go" ;
    code:description "Uri Mesh Gateway module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :SpiffeIDMeshGateway ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"net/url"

	"github.com/hashicorp/consul/acl"
)

type SpiffeIDMeshGateway struct {
	Host       string
	Partition  string
	Datacenter string
}

func (id SpiffeIDMeshGateway) MatchesPartition(partition string) bool {
	return id.PartitionOrDefault() == acl.PartitionOrDefault(partition)
}

func (id SpiffeIDMeshGateway) PartitionOrDefault() string {
	return acl.PartitionOrDefault(id.Partition)
}

// URI returns the *url.URL for this SPIFFE ID.
func (id SpiffeIDMeshGateway) URI() *url.URL {
	var result url.URL
	result.Scheme = "spiffe"
	result.Host = id.Host
	result.Path = id.uriPath()
	return &result
}
