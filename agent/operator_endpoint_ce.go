// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/operator_endpoint_ce.go
Operator Endpoint Ce module for agent layer

## Linked Modules
- [api](../api)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/operator_endpoint_ce.go> a code:Module ;
    code:name "agent/operator_endpoint_ce.go" ;
    code:description "Operator Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"github.com/hashicorp/consul/api"
	autopilot "github.com/hashicorp/raft-autopilot"
)

func autopilotToAPIServerEnterprise(_ *autopilot.ServerState, _ *api.AutopilotServer) {
	// noop in ce
}

func autopilotToAPIStateEnterprise(state *autopilot.State, apiState *api.AutopilotState) {
	// without the enterprise features there is no different between these two and we don't want to
	// alarm anyone by leaving this as the zero value.
	apiState.OptimisticFailureTolerance = state.FailureTolerance
}
