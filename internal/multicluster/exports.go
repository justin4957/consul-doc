// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/exports.go
Exports module for internal layer

## Linked Modules
- [internal/controller](../internal/controller)
- [internal/multicluster/internal/controllers](../internal/multicluster/internal/controllers)
- [internal/multicluster/internal/controllers/v1compat](../internal/multicluster/internal/controllers/v1compat)
- [internal/multicluster/internal/types](../internal/multicluster/internal/types)
- [internal/resource](../internal/resource)

## Tags
internal

## Exports
CompatControllerDependencies, DefaultCompatControllerDependencies, RegisterCompatControllers, RegisterTypes

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/exports.go> a code:Module ;
    code:name "internal/multicluster/exports.go" ;
    code:description "Exports module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/controller" ;
        code:path "../internal/controller"
    ], [
        code:name "internal/multicluster/internal/controllers" ;
        code:path "../internal/multicluster/internal/controllers"
    ], [
        code:name "internal/multicluster/internal/controllers/v1compat" ;
        code:path "../internal/multicluster/internal/controllers/v1compat"
    ], [
        code:name "internal/multicluster/internal/types" ;
        code:path "../internal/multicluster/internal/types"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ] ;
    code:exports :CompatControllerDependencies, :DefaultCompatControllerDependencies, :RegisterCompatControllers, :RegisterTypes ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package multicluster

import (
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/multicluster/internal/controllers"
	"github.com/hashicorp/consul/internal/multicluster/internal/controllers/v1compat"
	"github.com/hashicorp/consul/internal/multicluster/internal/types"
	"github.com/hashicorp/consul/internal/resource"
)

// RegisterTypes adds all resource types within the "multicluster" API group
// to the given type registry
func RegisterTypes(r resource.Registry) {
	types.Register(r)
}

type CompatControllerDependencies = controllers.CompatDependencies

func DefaultCompatControllerDependencies(ac v1compat.AggregatedConfig) CompatControllerDependencies {
	return CompatControllerDependencies{
		ConfigEntryExports: ac,
	}
}

func RegisterCompatControllers(mgr *controller.Manager, deps CompatControllerDependencies) {
	controllers.RegisterCompat(mgr, deps)
}
