// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/decoded.go
Decoded module for internal layer

## Linked Modules
- [internal/resource](../internal/resource)
- [proto-public/pbmulticluster/v2](../proto-public/pbmulticluster/v2)

## Tags
data-model, internal, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/decoded.go> a code:Module ;
    code:name "internal/multicluster/internal/types/decoded.go" ;
    code:description "Decoded module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbmulticluster/v2" ;
        code:path "../proto-public/pbmulticluster/v2"
    ] ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package types

import (
	"github.com/hashicorp/consul/internal/resource"
	v2 "github.com/hashicorp/consul/proto-public/pbmulticluster/v2"
)

type (
	DecodedExportedServices          = resource.DecodedResource[*v2.ExportedServices]
	DecodedNamespaceExportedServices = resource.DecodedResource[*v2.NamespaceExportedServices]
	DecodedPartitionExportedServices = resource.DecodedResource[*v2.PartitionExportedServices]
)
