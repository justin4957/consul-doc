// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/agent_endpoint_ce.go
Agent Endpoint Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/agent_endpoint_ce.go> a code:Module ;
    code:name "agent/agent_endpoint_ce.go" ;
    code:description "Agent Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"net/http"

	"github.com/hashicorp/consul/acl"
)

func (s *HTTPHandlers) validateRequestPartition(_ http.ResponseWriter, _ *acl.EnterpriseMeta) bool {
	return true
}

func (s *HTTPHandlers) defaultMetaPartitionToAgent(entMeta *acl.EnterpriseMeta) {
}
