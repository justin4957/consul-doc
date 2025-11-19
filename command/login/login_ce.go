// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: command/login/login_ce.go
Login Ce module for cli layer

## Tags
cli, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/login/login_ce.go> a code:Module ;
    code:name "command/login/login_ce.go" ;
    code:description "Login Ce module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package login

type enterpriseCmd struct {
}

func (c *cmd) initEnterpriseFlags() {
}

func (c *cmd) login() int {
	return c.bearerTokenLogin()
}
