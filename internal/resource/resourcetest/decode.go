// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/resourcetest/decode.go
Decode module for internal layer

## Linked Modules
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
MustDecode, MustDecodeList

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/resourcetest/decode.go> a code:Module ;
    code:name "internal/resource/resourcetest/decode.go" ;
    code:description "Decode module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :MustDecode, :MustDecodeList ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcetest

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

func MustDecode[Tp proto.Message](t T, res *pbresource.Resource) *resource.DecodedResource[Tp] {
	dec, err := resource.Decode[Tp](res)
	require.NoError(t, err)
	return dec
}

func MustDecodeList[Tp proto.Message](t T, resources []*pbresource.Resource) []*resource.DecodedResource[Tp] {
	dec, err := resource.DecodeList[Tp](resources)
	require.NoError(t, err)
	return dec
}
