// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/controllers/register.go
Register module for internal layer

## Linked Modules
- [internal/controller](../internal/controller)
- [internal/multicluster/internal/controllers/v1compat](../internal/multicluster/internal/controllers/v1compat)

## Tags
internal

## Exports
CompatDependencies, RegisterCompat

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/controllers/register.go> a code:Module ;
    code:name "internal/multicluster/internal/controllers/register.go" ;
    code:description "Register module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/controller" ;
        code:path "../internal/controller"
    ], [
        code:name "internal/multicluster/internal/controllers/v1compat" ;
        code:path "../internal/multicluster/internal/controllers/v1compat"
    ] ;
    code:exports :CompatDependencies, :RegisterCompat ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package controllers

import (
	"github.com/hashicorp/consul/internal/controller"
	"github.com/hashicorp/consul/internal/multicluster/internal/controllers/v1compat"
)

type CompatDependencies struct {
	ConfigEntryExports v1compat.AggregatedConfig
}

func RegisterCompat(mgr *controller.Manager, deps CompatDependencies) {
	mgr.Register(v1compat.Controller(deps.ConfigEntryExports))
}
