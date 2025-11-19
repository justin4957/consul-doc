// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/enterprisemeta_ce.go
Enterprisemeta Ce module for acl layer

## Tags
access-control, acl, authorization, security

## Exports
DefaultEnterpriseMeta, EnterpriseMeta, EqualNamespaces, EqualPartitions, IsDefaultPartition, NamespaceOrDefault, NewEnterpriseMetaWithPartition, NormalizeNamespace, PartitionOrDefault, WildcardEnterpriseMeta

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/enterprisemeta_ce.go> a code:Module ;
    code:name "acl/enterprisemeta_ce.go" ;
    code:description "Enterprisemeta Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :DefaultEnterpriseMeta, :EnterpriseMeta, :EqualNamespaces, :EqualPartitions, :IsDefaultPartition, :NamespaceOrDefault, :NewEnterpriseMetaWithPartition, :NormalizeNamespace, :PartitionOrDefault, :WildcardEnterpriseMeta ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

import "hash"

var emptyEnterpriseMeta = EnterpriseMeta{}

// EnterpriseMeta stub
type EnterpriseMeta struct{}

func (m *EnterpriseMeta) ToEnterprisePolicyMeta() *EnterprisePolicyMeta {
	return nil
}

func DefaultEnterpriseMeta() *EnterpriseMeta {
	return &EnterpriseMeta{}
}

func WildcardEnterpriseMeta() *EnterpriseMeta {
	return &EnterpriseMeta{}
}

func (m *EnterpriseMeta) EstimateSize() int {
	return 0
}

func (m *EnterpriseMeta) AddToHash(_ hash.Hash, _ bool) {
	// do nothing
}

func (m *EnterpriseMeta) PartitionOrDefault() string {
	return "default"
}

func EqualPartitions(_, _ string) bool {
	return true
}

func IsDefaultPartition(partition string) bool {
	return true
}

func PartitionOrDefault(_ string) string {
	return "default"
}

func (m *EnterpriseMeta) PartitionOrEmpty() string {
	return ""
}

func (m *EnterpriseMeta) InDefaultPartition() bool {
	return true
}

func (m *EnterpriseMeta) NamespaceOrDefault() string {
	return DefaultNamespaceName
}

func EqualNamespaces(_, _ string) bool {
	return true
}

func NamespaceOrDefault(_ string) string {
	return DefaultNamespaceName
}

func (m *EnterpriseMeta) NamespaceOrEmpty() string {
	return ""
}

func (m *EnterpriseMeta) InDefaultNamespace() bool {
	return true
}

func (m *EnterpriseMeta) Merge(_ *EnterpriseMeta) {
	// do nothing
}

func (m *EnterpriseMeta) MergeNoWildcard(_ *EnterpriseMeta) {
	// do nothing
}

func (*EnterpriseMeta) Normalize()          {}
func (*EnterpriseMeta) NormalizePartition() {}
func (*EnterpriseMeta) NormalizeNamespace() {}

func (m *EnterpriseMeta) Matches(_ *EnterpriseMeta) bool {
	return true
}

func (m *EnterpriseMeta) IsSame(_ *EnterpriseMeta) bool {
	return true
}

func (m *EnterpriseMeta) LessThan(_ *EnterpriseMeta) bool {
	return false
}

func (m *EnterpriseMeta) WithWildcardNamespace() *EnterpriseMeta {
	return &emptyEnterpriseMeta
}

func (m *EnterpriseMeta) UnsetPartition() {
	// do nothing
}

func (m *EnterpriseMeta) OverridePartition(_ string) {
	// do nothing
}

func NewEnterpriseMetaWithPartition(_, _ string) EnterpriseMeta {
	return emptyEnterpriseMeta
}

// FillAuthzContext stub
func (*EnterpriseMeta) FillAuthzContext(_ *AuthorizerContext) {}

func NormalizeNamespace(_ string) string {
	return ""
}
