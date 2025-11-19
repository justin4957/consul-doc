// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/types.go
Types module for internal layer

## Linked Modules
- [internal/resource](../internal/resource)

## Tags
data-model, internal, types

## Exports
Register

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/types.go> a code:Module ;
    code:name "internal/multicluster/internal/types/types.go" ;
    code:description "Types module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ] ;
    code:exports :Register ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package types

import (
	"github.com/hashicorp/consul/internal/resource"
)

func Register(r resource.Registry) {
	RegisterExportedServices(r)
	RegisterNamespaceExportedServices(r)
	RegisterPartitionExportedServices(r)
	RegisterComputedExportedServices(r)
}
