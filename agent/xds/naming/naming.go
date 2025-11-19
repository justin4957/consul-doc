// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/xds/naming/naming.go
Naming module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, envoy, service-mesh

## Exports
CustomizeClusterName

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/naming/naming.go> a code:Module ;
    code:name "agent/xds/naming/naming.go" ;
    code:description "Naming module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :CustomizeClusterName ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package naming

import (
	"fmt"

	"github.com/hashicorp/consul/agent/structs"
)

const (
	// OriginalDestinationClusterName is the name we give to the passthrough
	// cluster which redirects transparently-proxied requests to their original
	// destination outside the mesh. This cluster prevents Consul from blocking
	// connections to destinations outside of the catalog when in transparent
	// proxy mode.
	OriginalDestinationClusterName = "original-destination"
	VirtualIPTag                   = "virtual"
)

func CustomizeClusterName(clusterName string, chain *structs.CompiledDiscoveryChain) string {
	if chain == nil || chain.CustomizationHash == "" {
		return clusterName
	}
	// Use a colon to delimit this prefix instead of a dot to avoid a
	// theoretical collision problem with subsets.
	return fmt.Sprintf("%s~%s", chain.CustomizationHash, clusterName)
}
