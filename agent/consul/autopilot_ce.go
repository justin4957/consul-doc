// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/autopilot_ce.go
Autopilot Ce module for agent layer

## Linked Modules
- [agent/metadata](../agent/metadata)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/autopilot_ce.go> a code:Module ;
    code:name "agent/consul/autopilot_ce.go" ;
    code:description "Autopilot Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/metadata" ;
        code:path "../agent/metadata"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/agent/metadata"
	autopilot "github.com/hashicorp/raft-autopilot"
)

func (s *Server) autopilotPromoter() autopilot.Promoter {
	return autopilot.DefaultPromoter()
}

func (*Server) autopilotServerExt(_ *metadata.Server) interface{} {
	return nil
}
