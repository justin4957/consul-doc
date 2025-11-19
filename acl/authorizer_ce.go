// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/authorizer_ce.go
Authorizer Ce module for acl layer

## Tags
access-control, acl, authentication, authorization, security

## Exports
AuthorizerContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/authorizer_ce.go> a code:Module ;
    code:name "acl/authorizer_ce.go" ;
    code:description "Authorizer Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :AuthorizerContext ;
    code:tags "access-control", "acl", "authentication", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

// AuthorizerContext contains extra information that can be
// used in the determination of an ACL enforcement decision.
type AuthorizerContext struct {
	// Peer is the name of the peer that the resource was imported from.
	Peer string
}

func (c *AuthorizerContext) PeerOrEmpty() string {
	if c == nil {
		return ""
	}
	return c.Peer
}

// enterpriseAuthorizer stub interface
type enterpriseAuthorizer interface{}

func enforceEnterprise(_ Authorizer, _ Resource, _ string, _ string, _ *AuthorizerContext) (bool, EnforcementDecision, error) {
	return false, Deny, nil
}
