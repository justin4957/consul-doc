// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/uri_agent.go
Uri Agent module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

## Exports
SpiffeIDAgent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_agent.go> a code:Module ;
    code:name "agent/connect/uri_agent.go" ;
    code:description "Uri Agent module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :SpiffeIDAgent ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"net/url"

	"github.com/hashicorp/consul/acl"
)

// SpiffeIDService is the structure to represent the SPIFFE ID for an agent.
type SpiffeIDAgent struct {
	Host       string
	Partition  string
	Datacenter string
	Agent      string
}

func (id SpiffeIDAgent) PartitionOrDefault() string {
	return acl.PartitionOrDefault(id.Partition)
}

// URI returns the *url.URL for this SPIFFE ID.
func (id SpiffeIDAgent) URI() *url.URL {
	var result url.URL
	result.Scheme = "spiffe"
	result.Host = id.Host
	result.Path = id.uriPath()
	return &result
}
