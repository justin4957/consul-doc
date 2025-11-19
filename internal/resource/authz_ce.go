// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: internal/resource/authz_ce.go
Authz Ce module for internal layer

## Linked Modules
- [acl](../acl)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
authentication, internal, security

## Exports
AuthorizerContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/authz_ce.go> a code:Module ;
    code:name "internal/resource/authz_ce.go" ;
    code:description "Authz Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :AuthorizerContext ;
    code:tags "authentication", "internal", "security" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// AuthorizerContext builds an ACL AuthorizerContext for the given tenancy.
func AuthorizerContext(t *pbresource.Tenancy) *acl.AuthorizerContext {
	// TODO(peering/v2) handle non-local peers here
	return &acl.AuthorizerContext{}
}
