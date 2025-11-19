// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/discoverychain/compile_ce.go
Compile Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, discovery, dns

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/discoverychain/compile_ce.go> a code:Module ;
    code:name "agent/consul/discoverychain/compile_ce.go" ;
    code:description "Compile Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "discovery", "dns" .
<!-- End LinkedDoc RDF -->
*/
package discoverychain

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func (c *compiler) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return structs.DefaultEnterpriseMetaInDefaultPartition()
}
