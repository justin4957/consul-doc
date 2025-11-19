// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/uri_server.go
Uri Server module for agent layer

## Tags
agent, mtls, service-mesh

## Exports
PeeringServerSAN, SpiffeIDServer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_server.go> a code:Module ;
    code:name "agent/connect/uri_server.go" ;
    code:description "Uri Server module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :PeeringServerSAN, :SpiffeIDServer ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"fmt"
	"net/url"
)

type SpiffeIDServer struct {
	Host       string
	Datacenter string
}

// URI returns the *url.URL for this SPIFFE ID.
func (id SpiffeIDServer) URI() *url.URL {
	var result url.URL
	result.Scheme = "spiffe"
	result.Host = id.Host
	result.Path = fmt.Sprintf("/agent/server/dc/%s", id.Datacenter)
	return &result
}

// PeeringServerSAN returns the DNS SAN to attach to server certificates
// for control-plane peering traffic.
func PeeringServerSAN(dc, trustDomain string) string {
	return fmt.Sprintf("server.%s.peering.%s", dc, trustDomain)
}
