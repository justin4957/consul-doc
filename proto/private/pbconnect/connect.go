// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbconnect/connect.go
Connect module for internal layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)
- [proto/private/pbcommon](../proto/private/pbcommon)

## Tags
internal, mtls, service-mesh

## Exports
CARootToStructs, CARootsToStructs, EnterpriseMetaFrom, EnterpriseMetaTo, IssuedCertToStructs, NewCARootFromStructs, NewCARootsFromStructs, NewIssuedCertFromStructs, QueryMetaFrom, QueryMetaTo, RaftIndexFrom, RaftIndexTo

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbconnect/connect.go> a code:Module ;
    code:name "proto/private/pbconnect/connect.go" ;
    code:description "Connect module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pbcommon" ;
        code:path "../proto/private/pbcommon"
    ] ;
    code:exports :CARootToStructs, :CARootsToStructs, :EnterpriseMetaFrom, :EnterpriseMetaTo, :IssuedCertToStructs, :NewCARootFromStructs, :NewCARootsFromStructs, :NewIssuedCertFromStructs, :QueryMetaFrom, :QueryMetaTo, :RaftIndexFrom, :RaftIndexTo ;
    code:tags "internal", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package pbconnect

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbcommon"
)

func QueryMetaFrom(f structs.QueryMeta) *pbcommon.QueryMeta {
	t := new(pbcommon.QueryMeta)
	pbcommon.QueryMetaFromStructs(&f, t)
	return t
}

func QueryMetaTo(f *pbcommon.QueryMeta) structs.QueryMeta {
	t := new(structs.QueryMeta)
	pbcommon.QueryMetaToStructs(f, t)
	return *t
}

func RaftIndexFrom(f structs.RaftIndex) *pbcommon.RaftIndex {
	t := new(pbcommon.RaftIndex)
	pbcommon.RaftIndexFromStructs(&f, t)
	return t
}

func RaftIndexTo(f *pbcommon.RaftIndex) structs.RaftIndex {
	t := new(structs.RaftIndex)
	pbcommon.RaftIndexToStructs(f, t)
	return *t
}

func EnterpriseMetaFrom(f acl.EnterpriseMeta) *pbcommon.EnterpriseMeta {
	t := new(pbcommon.EnterpriseMeta)
	pbcommon.EnterpriseMetaFromStructs(&f, t)
	return t
}

func EnterpriseMetaTo(f *pbcommon.EnterpriseMeta) acl.EnterpriseMeta {
	t := new(acl.EnterpriseMeta)
	pbcommon.EnterpriseMetaToStructs(f, t)
	return *t
}

func NewIssuedCertFromStructs(in *structs.IssuedCert) (*IssuedCert, error) {
	t := new(IssuedCert)
	IssuedCertFromStructsIssuedCert(in, t)
	return t, nil
}

func NewCARootsFromStructs(in *structs.IndexedCARoots) (*CARoots, error) {
	t := new(CARoots)
	CARootsFromStructsIndexedCARoots(in, t)
	return t, nil
}

func CARootsToStructs(in *CARoots) (*structs.IndexedCARoots, error) {
	t := new(structs.IndexedCARoots)
	CARootsToStructsIndexedCARoots(in, t)
	return t, nil
}

func NewCARootFromStructs(in *structs.CARoot) (*CARoot, error) {
	t := new(CARoot)
	CARootFromStructsCARoot(in, t)
	return t, nil
}

func CARootToStructs(in *CARoot) (*structs.CARoot, error) {
	t := new(structs.CARoot)
	CARootToStructsCARoot(in, t)
	return t, nil
}

func IssuedCertToStructs(in *IssuedCert) (*structs.IssuedCert, error) {
	t := new(structs.IssuedCert)
	IssuedCertToStructsIssuedCert(in, t)
	return t, nil
}
