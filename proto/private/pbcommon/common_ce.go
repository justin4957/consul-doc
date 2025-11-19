// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: proto/private/pbcommon/common_ce.go
Common Ce module for internal layer

## Linked Modules
- [acl](../acl)

## Tags
internal

## Exports
DefaultEnterpriseMeta, EnterpriseMetaFromStructs, EnterpriseMetaToStructs, NewEnterpriseMetaFromStructs

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbcommon/common_ce.go> a code:Module ;
    code:name "proto/private/pbcommon/common_ce.go" ;
    code:description "Common Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :DefaultEnterpriseMeta, :EnterpriseMetaFromStructs, :EnterpriseMetaToStructs, :NewEnterpriseMetaFromStructs ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbcommon

import "github.com/hashicorp/consul/acl"

var DefaultEnterpriseMeta = &EnterpriseMeta{}

func NewEnterpriseMetaFromStructs(_ acl.EnterpriseMeta) *EnterpriseMeta {
	return &EnterpriseMeta{}
}
func EnterpriseMetaToStructs(s *EnterpriseMeta, t *acl.EnterpriseMeta) {
	if s == nil {
		return
	}
}
func EnterpriseMetaFromStructs(t *acl.EnterpriseMeta, s *EnterpriseMeta) {
	if s == nil {
		return
	}
}
