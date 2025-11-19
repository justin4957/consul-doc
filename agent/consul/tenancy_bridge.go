// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/tenancy_bridge.go
Tenancy Bridge module for agent layer

## Linked Modules
- [agent/grpc-external/services/resource](../agent/grpc-external/services/resource)

## Tags
agent

## Exports
NewV1TenancyBridge, V1TenancyBridge

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/tenancy_bridge.go> a code:Module ;
    code:name "agent/consul/tenancy_bridge.go" ;
    code:description "Tenancy Bridge module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/grpc-external/services/resource" ;
        code:path "../agent/grpc-external/services/resource"
    ] ;
    code:exports :NewV1TenancyBridge, :V1TenancyBridge ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import "github.com/hashicorp/consul/agent/grpc-external/services/resource"

// V1TenancyBridge is used by the resource service to access V1 implementations of
// partitions and namespaces. This bridge will be removed when V2 implemenations
// of partitions and namespaces are available.
type V1TenancyBridge struct {
	server *Server
}

func NewV1TenancyBridge(server *Server) resource.TenancyBridge {
	return &V1TenancyBridge{server: server}
}
