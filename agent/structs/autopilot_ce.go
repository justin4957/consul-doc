// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/autopilot_ce.go
Autopilot Ce module for agent layer

## Tags
agent, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/autopilot_ce.go> a code:Module ;
    code:name "agent/structs/autopilot_ce.go" ;
    code:description "Autopilot Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

func (c *AutopilotConfig) autopilotConfigExt() interface{} {
	return nil
}
