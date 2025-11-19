// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/authmethod/testauth/testing_ce.go
Testing Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, authentication, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/authmethod/testauth/testing_ce.go> a code:Module ;
    code:name "agent/consul/authmethod/testauth/testing_ce.go" ;
    code:description "Testing Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package testauth

import "github.com/hashicorp/consul/acl"

type enterpriseConfig struct{}

func (v *Validator) testAuthEntMetaFromFields(fields map[string]string) *acl.EnterpriseMeta {
	return nil
}
