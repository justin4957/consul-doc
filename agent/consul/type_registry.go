// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/type_registry.go
Type Registry module for agent layer

## Linked Modules
- [internal/multicluster](../internal/multicluster)
- [internal/resource](../internal/resource)
- [internal/resource/demo](../internal/resource/demo)

## Tags
agent

## Exports
NewTypeRegistry

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/type_registry.go> a code:Module ;
    code:name "agent/consul/type_registry.go" ;
    code:description "Type Registry module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "internal/multicluster" ;
        code:path "../internal/multicluster"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "internal/resource/demo" ;
        code:path "../internal/resource/demo"
    ] ;
    code:exports :NewTypeRegistry ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/internal/multicluster"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/resource/demo"
)

// NewTypeRegistry returns a registry populated with all supported resource
// types.
//
// Note: the registry includes resource types that may not be suitable for
// production use (e.g. experimental or development resource types) because
// it is used in the CLI, where feature flags and other runtime configuration
// may not be available.
func NewTypeRegistry() resource.Registry {
	registry := resource.NewRegistry()

	demo.RegisterTypes(registry)
	multicluster.RegisterTypes(registry)

	return registry
}
