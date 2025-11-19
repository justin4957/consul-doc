// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/xds/failover_policy_ce.go
Failover Policy Ce module for agent layer

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/failover_policy_ce.go> a code:Module ;
    code:name "agent/xds/failover_policy_ce.go" ;
    code:description "Failover Policy Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	"fmt"
)

func (ft discoChainTargets) orderByLocality() ([]discoChainTargetGroup, error) {
	return nil, fmt.Errorf("order-by-locality is a Consul Enterprise feature")
}
