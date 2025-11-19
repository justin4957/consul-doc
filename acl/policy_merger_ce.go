// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/policy_merger_ce.go
Policy Merger Ce module for acl layer

## Tags
access-control, acl, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/policy_merger_ce.go> a code:Module ;
    code:name "acl/policy_merger_ce.go" ;
    code:description "Policy Merger Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

type enterprisePolicyRulesMergeContext struct{}

func (ctx *enterprisePolicyRulesMergeContext) init() {
	// do nothing
}

func (ctx *enterprisePolicyRulesMergeContext) merge(*EnterprisePolicyRules) {
	// do nothing
}

func (ctx *enterprisePolicyRulesMergeContext) fill(*EnterprisePolicyRules) {
	// do nothing
}
