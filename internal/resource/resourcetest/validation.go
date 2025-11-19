// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/resourcetest/validation.go
Validation module for internal layer

## Linked Modules
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
ValidateAndNormalize

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/resourcetest/validation.go> a code:Module ;
    code:name "internal/resource/resourcetest/validation.go" ;
    code:description "Validation module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :ValidateAndNormalize ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcetest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func ValidateAndNormalize(t *testing.T, registry resource.Registry, res *pbresource.Resource) {
	t.Helper()
	typ := res.Id.Type

	typeInfo, ok := registry.Resolve(typ)
	if !ok {
		t.Fatalf("unhandled resource type: %q", resource.ToGVK(typ))
	}

	if typeInfo.Mutate != nil {
		require.NoError(t, typeInfo.Mutate(res), "failed to apply type mutation to resource")
	}

	if typeInfo.Validate != nil {
		require.NoError(t, typeInfo.Validate(res), "failed to validate resource")
	}
}
