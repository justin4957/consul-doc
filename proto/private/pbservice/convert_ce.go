// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: proto/private/pbservice/convert_ce.go
Convert Ce module for internal layer

## Linked Modules
- [acl](../acl)
- [proto/private/pbcommon](../proto/private/pbcommon)

## Tags
internal

## Exports
EnterpriseMetaToStructs, NewEnterpriseMetaFromStructs

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbservice/convert_ce.go> a code:Module ;
    code:name "proto/private/pbservice/convert_ce.go" ;
    code:description "Convert Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "proto/private/pbcommon" ;
        code:path "../proto/private/pbcommon"
    ] ;
    code:exports :EnterpriseMetaToStructs, :NewEnterpriseMetaFromStructs ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbservice

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/proto/private/pbcommon"
)

func EnterpriseMetaToStructs(_ *pbcommon.EnterpriseMeta) acl.EnterpriseMeta {
	return acl.EnterpriseMeta{}
}

func NewEnterpriseMetaFromStructs(_ acl.EnterpriseMeta) *pbcommon.EnterpriseMeta {
	return &pbcommon.EnterpriseMeta{}
}
