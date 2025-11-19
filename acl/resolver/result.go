// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: acl/resolver/result.go
ACL resolution result combining authorizer with ACL identity information.

## Linked Modules
- [acl/authorizer](../authorizer.go) - Wraps Authorizer interface
- [acl/acl](../acl.go) - Core ACL types
- [agent/structs](../../agent/structs/acl.go) - ACL identity structures

## Tags
security, authorization, resolution, identity

## Exports
Result

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/resolver/result.go> a code:Module ;
    code:name "acl/resolver/result.go" ;
    code:description "ACL resolution result combining authorizer with ACL identity information" ;
    code:layer "acl" ;
    code:linksTo [
        code:name "acl/authorizer" ;
        code:path "../authorizer.go" ;
        code:relationship "Wraps Authorizer interface"
    ], [
        code:name "acl/acl" ;
        code:path "../acl.go" ;
        code:relationship "Core ACL types"
    ], [
        code:name "agent/structs" ;
        code:path "../../agent/structs/acl.go" ;
        code:relationship "ACL identity structures"
    ] ;
    code:exports :Result ;
    code:tags "security", "authorization", "resolution", "identity" .
<!-- End LinkedDoc RDF -->
*/
package resolver

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

type Result struct {
	acl.Authorizer
	// TODO: likely we can reduce this interface
	ACLIdentity structs.ACLIdentity
}

func (a Result) AccessorID() string {
	if a.ACLIdentity == nil {
		return ""
	}
	return a.ACLIdentity.ID()
}

func (a Result) Identity() structs.ACLIdentity {
	return a.ACLIdentity
}

func (a Result) ToAllowAuthorizer() acl.AllowAuthorizer {
	return acl.AllowAuthorizer{Authorizer: a, AccessorID: a.AccessorID()}
}
