// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/grpc-external/services/resource/testing/testing_ce.go
Testing Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, api, communication, grpc, networking

## Exports
FillAuthorizerContext, FillEntMeta

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/resource/testing/testing_ce.go> a code:Module ;
    code:name "agent/grpc-external/services/resource/testing/testing_ce.go" ;
    code:description "Testing Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :FillAuthorizerContext, :FillEntMeta ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testing

import (
	"github.com/hashicorp/consul/acl"
)

func FillEntMeta(entMeta *acl.EnterpriseMeta) {
	// nothing to to in CE.
}

func FillAuthorizerContext(authzContext *acl.AuthorizerContext) {
	// nothing to to in CE.
}
