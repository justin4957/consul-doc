// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/configentry/config_entry.go
Config Entry module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, configuration

## Exports
KindName, NewKindName, NewKindNameForEntry

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/configentry/config_entry.go> a code:Module ;
    code:name "agent/configentry/config_entry.go" ;
    code:description "Config Entry module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :KindName, :NewKindName, :NewKindNameForEntry ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package configentry

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

// KindName is a value type useful for maps. You can use:
//
//	map[KindName]Payload
//
// instead of:
//
//	map[string]map[string]Payload
type KindName struct {
	Kind string
	Name string
	acl.EnterpriseMeta
}

// NewKindName returns a new KindName. The EnterpriseMeta values will be
// normalized based on the kind.
//
// Any caller which modifies the EnterpriseMeta field must call Normalize
// before persisting or using the value as a map key.
func NewKindName(kind, name string, entMeta *acl.EnterpriseMeta) KindName {
	ret := KindName{
		Kind: kind,
		Name: name,
	}
	if entMeta == nil {
		entMeta = structs.DefaultEnterpriseMetaInDefaultPartition()
	}

	ret.EnterpriseMeta = *entMeta
	ret.Normalize()
	return ret
}

func NewKindNameForEntry(entry structs.ConfigEntry) KindName {
	return NewKindName(entry.GetKind(), entry.GetName(), entry.GetEnterpriseMeta())
}
