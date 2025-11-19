// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/policy_authorizer_ce.go
Policy Authorizer Ce module for acl layer

## Tags
access-control, acl, authentication, authorization, security

## Exports
NewPolicyAuthorizer, NewPolicyAuthorizerWithDefaults

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/policy_authorizer_ce.go> a code:Module ;
    code:name "acl/policy_authorizer_ce.go" ;
    code:description "Policy Authorizer Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :NewPolicyAuthorizer, :NewPolicyAuthorizerWithDefaults ;
    code:tags "access-control", "acl", "authentication", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

// enterprisePolicyAuthorizer stub
type enterprisePolicyAuthorizer struct{}

func (authz *enterprisePolicyAuthorizer) init(*Config) {
	// nothing to do
}

func (authz *enterprisePolicyAuthorizer) enforce(_ *EnterpriseRule, _ *AuthorizerContext) EnforcementDecision {
	return Default
}

// NewPolicyAuthorizer merges the policies and returns an Authorizer that will enforce them
func NewPolicyAuthorizer(policies []*Policy, entConfig *Config) (Authorizer, error) {
	return newPolicyAuthorizer(policies, entConfig)
}

// NewPolicyAuthorizerWithDefaults will actually created a ChainedAuthorizer with
// the policies compiled into one Authorizer and the backup policy of the defaultAuthz
func NewPolicyAuthorizerWithDefaults(defaultAuthz Authorizer, policies []*Policy, entConfig *Config) (Authorizer, error) {
	authz, err := newPolicyAuthorizer(policies, entConfig)
	if err != nil {
		return nil, err
	}

	return NewChainedAuthorizer([]Authorizer{authz, defaultAuthz}), nil
}
