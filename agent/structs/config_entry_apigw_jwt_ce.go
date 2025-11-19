// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_apigw_jwt_ce.go
Config Entry Apigw Jwt Ce module for agent layer

## Tags
agent, api, client, configuration, data-model, types

## Exports
APIGatewayJWTRequirement, JWTFilter

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_apigw_jwt_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_apigw_jwt_ce.go" ;
    code:description "Config Entry Apigw Jwt Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :APIGatewayJWTRequirement, :JWTFilter ;
    code:tags "agent", "api", "client", "configuration", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

// APIGatewayJWTRequirement holds the list of JWT providers to be verified against
type APIGatewayJWTRequirement struct{}

// JWTFilter holds the JWT Filter configuration for an HTTPRoute
type JWTFilter struct{}
