// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/authmethod/ssoauth/sso_ce.go
Sso Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [internal/go-sso/oidcauth](../internal/go-sso/oidcauth)

## Tags
agent, authentication, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/authmethod/ssoauth/sso_ce.go> a code:Module ;
    code:name "agent/consul/authmethod/ssoauth/sso_ce.go" ;
    code:description "Sso Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "internal/go-sso/oidcauth" ;
        code:path "../internal/go-sso/oidcauth"
    ] ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package ssoauth

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/internal/go-sso/oidcauth"
)

func validateType(typ string) error {
	if typ != "jwt" {
		return fmt.Errorf("type should be %q", "jwt")
	}
	return nil
}

func (v *Validator) ssoEntMetaFromClaims(_ *oidcauth.Claims) *acl.EnterpriseMeta {
	return nil
}

type enterpriseConfig struct{}

func (c *Config) enterpriseConvertForLibrary(_ *oidcauth.Config) {}
