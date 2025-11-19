// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/xds/locality_policy.go
Locality Policy module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/locality_policy.go> a code:Module ;
    code:name "agent/xds/locality_policy.go" ;
    code:description "Locality Policy module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	"fmt"
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/agent/structs"
)

func groupedEndpoints(logger hclog.Logger, locality *structs.Locality, policy *structs.DiscoveryPrioritizeByLocality, csns structs.CheckServiceNodes) ([]structs.CheckServiceNodes, error) {
	switch {
	case policy == nil || policy.Mode == "" || policy.Mode == "none":
		return []structs.CheckServiceNodes{csns}, nil
	case policy.Mode == "failover":
		log := logger.Named("locality")
		return prioritizeByLocalityFailover(log, locality, csns), nil
	default:
		return nil, fmt.Errorf("unexpected priortize-by-locality mode %q", policy.Mode)
	}
}
