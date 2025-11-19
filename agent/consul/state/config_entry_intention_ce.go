// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/config_entry_intention_ce.go
Config Entry Intention Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/config_entry_intention_ce.go> a code:Module ;
    code:name "agent/consul/state/config_entry_intention_ce.go" ;
    code:description "Config Entry Intention Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "configuration", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func getIntentionPrecedenceMatchServiceNames(serviceName string, entMeta *acl.EnterpriseMeta) []structs.ServiceName {
	if serviceName == structs.WildcardSpecifier {
		return []structs.ServiceName{
			structs.NewServiceName(structs.WildcardSpecifier, entMeta),
		}
	}

	return []structs.ServiceName{
		structs.NewServiceName(serviceName, entMeta),
		structs.NewServiceName(structs.WildcardSpecifier, entMeta),
	}
}
