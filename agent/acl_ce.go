// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/acl_ce.go
Acl Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [api](../api)

## Tags
access-control, agent, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/acl_ce.go> a code:Module ;
    code:name "agent/acl_ce.go" ;
    code:description "Acl Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "access-control", "agent", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"github.com/hashicorp/serf/serf"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/api"
)

func serfMemberFillAuthzContext(m *serf.Member, ctx *acl.AuthorizerContext) {
	// no-op
}

func agentServiceFillAuthzContext(s *api.AgentService, ctx *acl.AuthorizerContext) {
	// no-op
}
