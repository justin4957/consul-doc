// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/connect/uri_agent_ce.go
Uri Agent Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, mtls, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/uri_agent_ce.go> a code:Module ;
    code:name "agent/connect/uri_agent_ce.go" ;
    code:description "Uri Agent Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

// GetEnterpriseMeta will synthesize an EnterpriseMeta struct from the SpiffeIDAgent.
// in CE this just returns an empty (but never nil) struct pointer
func (id SpiffeIDAgent) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return &acl.EnterpriseMeta{}
}

func (id SpiffeIDAgent) uriPath() string {
	return fmt.Sprintf("/agent/client/dc/%s/id/%s", id.Datacenter, id.Agent)
}
