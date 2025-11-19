// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/acls.go
Acls module for internal layer

## Linked Modules
- [acl](../acl)

## Tags
access-control, authorization, internal, security

## Exports
NoOpACLListHook

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/acls.go> a code:Module ;
    code:name "internal/resource/acls.go" ;
    code:description "Acls module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :NoOpACLListHook ;
    code:tags "access-control", "authorization", "internal", "security" .
<!-- End LinkedDoc RDF -->
*/
package resource

import "github.com/hashicorp/consul/acl"

// NoOpACLListHook is a common function that can be used if no special list permission is required for a resource.
func NoOpACLListHook(_ acl.Authorizer, _ *acl.AuthorizerContext) error {
	// No-op List permission as we want to default to filtering resources
	// from the list using the Read enforcement.
	return nil
}
