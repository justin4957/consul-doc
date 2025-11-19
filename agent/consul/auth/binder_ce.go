// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/auth/binder_ce.go
Binder Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/consul/authmethod](../agent/consul/authmethod)
- [agent/structs](../agent/structs)

## Tags
agent, authentication, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/auth/binder_ce.go> a code:Module ;
    code:name "agent/consul/auth/binder_ce.go" ;
    code:description "Binder Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/consul/authmethod" ;
        code:path "../agent/consul/authmethod"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package auth

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/consul/authmethod"
	"github.com/hashicorp/consul/agent/structs"
)

func bindEnterpriseMeta(authMethod *structs.ACLAuthMethod, verifiedIdentity *authmethod.Identity) (acl.EnterpriseMeta, error) {
	return acl.EnterpriseMeta{}, nil
}
