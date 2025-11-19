// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/connect.go
Connect module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, data-model, mtls, service-mesh, types

## Exports
ConnectAuthorizeRequest

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/connect.go> a code:Module ;
    code:name "agent/structs/connect.go" ;
    code:description "Connect module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :ConnectAuthorizeRequest ;
    code:tags "agent", "data-model", "mtls", "service-mesh", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import "github.com/hashicorp/consul/acl"

// ConnectAuthorizeRequest is the structure of a request to authorize
// a connection.
type ConnectAuthorizeRequest struct {
	// Target is the name of the service that is being requested.
	Target string

	// EnterpriseMeta is the embedded Consul Enterprise specific metadata
	acl.EnterpriseMeta

	// ClientCertURI is a unique identifier for the requesting client. This
	// is currently the URI SAN from the TLS client certificate.
	//
	// ClientCertSerial is a colon-hex-encoded of the serial number for
	// the requesting client cert. This is used to check against revocation
	// lists.
	ClientCertURI    string
	ClientCertSerial string
}

func (req *ConnectAuthorizeRequest) TargetPartition() string {
	return req.PartitionOrDefault()
}
