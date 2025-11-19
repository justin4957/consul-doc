// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/xds/locality_policy_ce.go
Locality Policy Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/locality_policy_ce.go> a code:Module ;
    code:name "agent/xds/locality_policy_ce.go" ;
    code:description "Locality Policy Ce module for agent layer" ;
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
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/go-hclog"
)

func prioritizeByLocalityFailover(_ hclog.Logger, _ *structs.Locality, _ structs.CheckServiceNodes) []structs.CheckServiceNodes {
	return nil
}
