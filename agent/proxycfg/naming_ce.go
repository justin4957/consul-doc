// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/proxycfg/naming_ce.go
Naming Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, networking, service-mesh

## Exports
UpstreamIDString

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/naming_ce.go> a code:Module ;
    code:name "agent/proxycfg/naming_ce.go" ;
    code:description "Naming Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :UpstreamIDString ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func UpstreamIDString(typ, dc, name string, _ *acl.EnterpriseMeta, peerName string) string {
	ret := name

	if peerName != "" {
		ret += "?peer=" + peerName
	} else if dc != "" {
		ret += "?dc=" + dc
	}

	if typ == "" || typ == structs.UpstreamDestTypeService {
		return ret
	}

	return typ + ":" + ret
}

func parseInnerUpstreamIDString(input string) (string, *acl.EnterpriseMeta) {
	return input, structs.DefaultEnterpriseMetaInDefaultPartition()
}

func (u UpstreamID) enterpriseIdentifierPrefix() string {
	return ""
}
