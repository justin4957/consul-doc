// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/leader_connect_ca_ce.go
Leader Connect Ca Ce module for agent layer

## Linked Modules
- [agent/connect](../agent/connect)

## Tags
agent, mtls, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/leader_connect_ca_ce.go> a code:Module ;
    code:name "agent/consul/leader_connect_ca_ce.go" ;
    code:description "Leader Connect Ca Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/connect" ;
        code:path "../agent/connect"
    ] ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/agent/connect"
)

func (c *CAManager) validateSupportedIdentityScopesInCertificate(spiffeID connect.CertURI) error {
	switch v := spiffeID.(type) {
	case *connect.SpiffeIDService:
		if v.Namespace != "default" || v.Partition != "default" {
			return connect.InvalidCSRError("Non default partition or namespace is supported in Enterprise only."+
				"Provided namespace is %s and partition is %s", v.Namespace, v.Partition)
		}
	case *connect.SpiffeIDMeshGateway:
		if v.Partition != "default" {
			return connect.InvalidCSRError("Non default partition is supported in Enterprise only."+
				"Provided partition is %s", v.Partition)
		}
	case *connect.SpiffeIDAgent, *connect.SpiffeIDServer:
		return nil
	default:
		c.logger.Trace("spiffe ID type is not expected", "spiffeID", spiffeID, "spiffeIDType", v)
		return connect.InvalidCSRError("SPIFFE ID in CSR must be a service, mesh-gateway, or agent ID")
	}
	return nil
}
