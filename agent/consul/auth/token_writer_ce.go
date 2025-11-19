// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/auth/token_writer_ce.go
Token Writer Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, authentication, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/auth/token_writer_ce.go> a code:Module ;
    code:name "agent/consul/auth/token_writer_ce.go" ;
    code:description "Token Writer Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package auth

import "github.com/hashicorp/consul/agent/structs"

func (w *TokenWriter) enterpriseValidation(token, existing *structs.ACLToken) error {
	return nil
}
