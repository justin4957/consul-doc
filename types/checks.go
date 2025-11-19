// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: types/checks.go
Check identifier type definition for unique health check identification on agents.

## Linked Modules
- [agent/structs](../agent/structs/check_definition.go) - Check definition structures
- [agent/local](../agent/local/state.go) - Local agent check state
- [api/agent](../api/agent.go) - Agent API check types

## Tags
types, data-model, health-checks, identifiers

## Exports
CheckID

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<types/checks.go> a code:Module ;
    code:language "go" ;
    code:name "types/checks.go" ;
    code:description "Check identifier type definition for unique health check identification on agents" ;
    code:layer "api" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs/check_definition.go" ;
        code:relationship "Check definition structures"
    ], [
        code:name "agent/local" ;
        code:path "../agent/local/state.go" ;
        code:relationship "Local agent check state"
    ], [
        code:name "api/agent" ;
        code:path "../api/agent.go" ;
        code:relationship "Agent API check types"
    ] ;
    code:exports :CheckID ;
    code:tags "types", "data-model", "health-checks", "identifiers" .
<!-- End LinkedDoc RDF -->
*/
package types

// CheckID is a strongly typed string used to uniquely represent a Consul
// Check on an Agent (a CheckID is not globally unique).
type CheckID string
