// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: command/acl/authmethod/create/authmethod_create_ce.go
Authmethod Create Ce module for cli layer

## Linked Modules
- [api](../api)

## Tags
access-control, authentication, authorization, cli, security, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/acl/authmethod/create/authmethod_create_ce.go> a code:Module ;
    code:name "command/acl/authmethod/create/authmethod_create_ce.go" ;
    code:description "Authmethod Create Ce module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "access-control", "authentication", "authorization", "cli", "security", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package authmethodcreate

import "github.com/hashicorp/consul/api"

type enterpriseCmd struct {
}

func (c *cmd) initEnterpriseFlags() {}

func (c *cmd) enterprisePopulateAuthMethod(method *api.ACLAuthMethod) error {
	return nil
}
