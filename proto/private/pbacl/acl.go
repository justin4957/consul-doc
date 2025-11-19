// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbacl/acl.go
Acl module for internal layer

## Linked Modules
- [api](../api)

## Tags
access-control, authorization, internal, security

## Exports
ACLLinkFromAPI

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbacl/acl.go> a code:Module ;
    code:name "proto/private/pbacl/acl.go" ;
    code:description "Acl module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :ACLLinkFromAPI ;
    code:tags "access-control", "authorization", "internal", "security" .
<!-- End LinkedDoc RDF -->
*/
package pbacl

import (
	"github.com/hashicorp/consul/api"
)

func (a *ACLLink) ToAPI() api.ACLLink {
	return api.ACLLink{
		ID:   a.ID,
		Name: a.Name,
	}
}

func ACLLinkFromAPI(a api.ACLLink) *ACLLink {
	return &ACLLink{
		ID:   a.ID,
		Name: a.Name,
	}
}
