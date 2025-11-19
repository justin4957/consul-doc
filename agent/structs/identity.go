// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/identity.go
Identity module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, data-model, types

## Exports
Identity

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/identity.go> a code:Module ;
    code:name "agent/structs/identity.go" ;
    code:description "Identity module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :Identity ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import "github.com/hashicorp/consul/acl"

// Identity of some entity (ex: service, node, check).
//
// TODO: this type should replace ServiceID, ServiceName, and CheckID which all
// have roughly identical implementations.
type Identity struct {
	ID string
	acl.EnterpriseMeta
}
