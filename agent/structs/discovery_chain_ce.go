// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/discovery_chain_ce.go
Discovery Chain Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, data-model, discovery, dns, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/discovery_chain_ce.go> a code:Module ;
    code:name "agent/structs/discovery_chain_ce.go" ;
    code:description "Discovery Chain Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "data-model", "discovery", "dns", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/hashicorp/consul/acl"
)

func (t *DiscoveryTarget) GetEnterpriseMetadata() *acl.EnterpriseMeta {
	return DefaultEnterpriseMetaInDefaultPartition()
}
