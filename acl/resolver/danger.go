// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: acl/resolver/danger.go
Danger module for acl layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
access-control, acl, authorization, security

## Exports
DANGER

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/resolver/danger.go> a code:Module ;
    code:name "acl/resolver/danger.go" ;
    code:description "Danger module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :DANGER ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package resolver

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

// DANGER_NO_AUTH implements an ACL resolver short-circuit authorization in
// cases where it is handled somewhere else or expressly not required.
type DANGER_NO_AUTH struct{}

// ResolveTokenAndDefaultMeta returns an authorizer with unfettered permissions.
func (DANGER_NO_AUTH) ResolveTokenAndDefaultMeta(_ string, entMeta *acl.EnterpriseMeta, _ *acl.AuthorizerContext) (Result, error) {
	entMeta.Merge(structs.DefaultEnterpriseMetaInDefaultPartition())
	return Result{Authorizer: acl.ManageAll()}, nil
}
