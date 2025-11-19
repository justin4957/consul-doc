// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/config/config_ce.go
Config Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, configuration

## Exports
EnterpriseMeta

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/config_ce.go> a code:Module ;
    code:name "agent/config/config_ce.go" ;
    code:description "Config Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :EnterpriseMeta ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

// EnterpriseMeta stub
type EnterpriseMeta struct{}

func (*EnterpriseMeta) ToStructs() acl.EnterpriseMeta {
	return *structs.DefaultEnterpriseMetaInDefaultPartition()
}
