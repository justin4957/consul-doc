// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/structs.deepcopy_ce.go
Structs.Deepcopy Ce module for agent layer

## Tags
agent, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/structs.deepcopy_ce.go> a code:Module ;
    code:name "agent/structs/structs.deepcopy_ce.go" ;
    code:description "Structs.Deepcopy Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

// DeepCopy generates a deep copy of *APIGatewayJWTRequirement
func (o *APIGatewayJWTRequirement) DeepCopy() *APIGatewayJWTRequirement {
	return new(APIGatewayJWTRequirement)
}

// DeepCopy generates a deep copy of *JWTFilter
func (o *JWTFilter) DeepCopy() *JWTFilter {
	return new(JWTFilter)
}
