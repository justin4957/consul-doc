// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: proto/private/pbconfigentry/config_entry_ce.go
Config Entry Ce module for internal layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
configuration, internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbconfigentry/config_entry_ce.go> a code:Module ;
    code:name "proto/private/pbconfigentry/config_entry_ce.go" ;
    code:description "Config Entry Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "configuration", "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbconfigentry

import "github.com/hashicorp/consul/agent/structs"

func gwJWTRequirementToStructs(m *APIGatewayJWTRequirement) *structs.APIGatewayJWTRequirement {
	return &structs.APIGatewayJWTRequirement{}
}

func gwJWTRequirementFromStructs(*structs.APIGatewayJWTRequirement) *APIGatewayJWTRequirement {
	return &APIGatewayJWTRequirement{}
}

func routeJWTFilterToStructs(m *JWTFilter) *structs.JWTFilter {
	return &structs.JWTFilter{}
}

func routeJWTFilterFromStructs(*structs.JWTFilter) *JWTFilter {
	return &JWTFilter{}
}
