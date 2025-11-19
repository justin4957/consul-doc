// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/uri_service.go
Uri Service module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

## Exports
SpiffeIDService

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_service.go> a code:Module ;
    code:name "agent/connect/uri_service.go" ;
    code:description "Uri Service module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :SpiffeIDService ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/consul/acl"
)

// SpiffeIDService is the structure to represent the SPIFFE ID for a service.
type SpiffeIDService struct {
	Host       string
	Partition  string
	Namespace  string
	Datacenter string
	Service    string
}

func (id SpiffeIDService) NamespaceOrDefault() string {
	return acl.NamespaceOrDefault(id.Namespace)
}

func (id SpiffeIDService) MatchesPartition(partition string) bool {
	return id.PartitionOrDefault() == acl.PartitionOrDefault(partition)
}

// URI returns the *url.URL for this SPIFFE ID.
func (id SpiffeIDService) URI() *url.URL {
	var result url.URL
	result.Scheme = "spiffe"
	result.Host = id.Host
	result.Path = id.uriPath()
	return &result
}

func (id SpiffeIDService) uriPath() string {
	path := fmt.Sprintf("/ns/%s/dc/%s/svc/%s",
		id.NamespaceOrDefault(),
		id.Datacenter,
		id.Service,
	)

	// Although CE has no support for partitions, it still needs to be able to
	// handle exportedPartition from peered Consul Enterprise clusters in order
	// to generate the correct SpiffeID.
	// We intentionally avoid using pbpartition.DefaultName here to be CE friendly.
	if ap := id.PartitionOrDefault(); ap != "" && ap != "default" {
		return "/ap/" + ap + path
	}
	return path
}
