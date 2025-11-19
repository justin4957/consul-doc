// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/testing_service_definition.go
Testing Service Definition module for agent layer

## Tags
agent, data-model, types

## Exports
TestServiceDefinition, TestServiceDefinitionProxy

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/testing_service_definition.go> a code:Module ;
    code:name "agent/structs/testing_service_definition.go" ;
    code:description "Testing Service Definition module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :TestServiceDefinition, :TestServiceDefinitionProxy ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/mitchellh/go-testing-interface"
)

// TestServiceDefinition returns a ServiceDefinition for a typical service.
func TestServiceDefinition(t testing.T) *ServiceDefinition {
	return &ServiceDefinition{
		Name: "db",
		Port: 1234,
	}
}

// TestServiceDefinitionProxy returns a ServiceDefinition for a proxy.
func TestServiceDefinitionProxy(t testing.T) *ServiceDefinition {
	return &ServiceDefinition{
		Kind: ServiceKindConnectProxy,
		Name: "foo-proxy",
		Port: 1234,
		Proxy: &ConnectProxyConfig{
			DestinationServiceName: "db",
		},
	}
}
