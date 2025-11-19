// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: acl/policy_ce.go
Policy Ce module for acl layer

## Tags
access-control, acl, authorization, security

## Exports
EnterprisePolicyMeta, EnterprisePolicyRules, EnterpriseRule

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<acl/policy_ce.go> a code:Module ;
    code:name "acl/policy_ce.go" ;
    code:description "Policy Ce module for acl layer" ;
    code:language "go" ;
    code:layer "acl" ;
    code:exports :EnterprisePolicyMeta, :EnterprisePolicyRules, :EnterpriseRule ;
    code:tags "access-control", "acl", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package acl

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/hcl"
	"strings"
)

// EnterprisePolicyMeta stub
type EnterprisePolicyMeta struct{}

// EnterpriseRule stub
type EnterpriseRule struct{}

func (r *EnterpriseRule) Validate(string, *Config) error {
	// nothing to validate
	return nil
}

// EnterprisePolicyRules stub
type EnterprisePolicyRules struct{}

func (r *EnterprisePolicyRules) Validate(*Config) error {
	// nothing to validate
	return nil
}

func decodeRules(rules string, warnOnDuplicateKey bool, _ *Config, _ *EnterprisePolicyMeta) (*Policy, error) {
	p := &Policy{}

	err := hcl.DecodeErrorOnDuplicates(p, rules)

	if errIsDuplicateKey(err) && warnOnDuplicateKey {
		//because the snapshot saves the unparsed rules we have to assume some snapshots exist that shouldn't fail, but
		// have duplicates
		if err := hcl.Decode(p, rules); err != nil {
			hclog.Default().Warn("Warning- Duplicate key in ACL Policy ignored", "errorMessage", err.Error())
			return nil, fmt.Errorf("Failed to parse ACL rules: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("Failed to parse ACL rules: %v", err)
	}

	return p, nil
}

func errIsDuplicateKey(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "was already set. Each argument can only be defined once")
}
