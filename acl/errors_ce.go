// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/errors_ce.go
Errors Ce module for acl layer

## Tags
access-control, acl, authorization, security

## Exports
NewResourceDescriptor, ResourceDescriptor

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/errors_ce.go> a code:Module ;
    code:name "acl/errors_ce.go" ;
    code:description "Errors Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :NewResourceDescriptor, :ResourceDescriptor ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

// In some sense we really want this to contain an EnterpriseMeta, but
// this turns out to be a convenient place to hang helper functions off of.
type ResourceDescriptor struct {
	Name string
}

func NewResourceDescriptor(name string, _ *AuthorizerContext) ResourceDescriptor {
	return ResourceDescriptor{Name: name}
}

func (od *ResourceDescriptor) ToString() string {
	return "\"" + od.Name + "\""
}
