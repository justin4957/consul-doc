// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/token/store_ce.go
Store Ce module for agent layer

## Tags
agent

## Exports
EnterpriseConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/token/store_ce.go> a code:Module ;
    code:name "agent/token/store_ce.go" ;
    code:description "Store Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EnterpriseConfig ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package token

type EnterpriseConfig struct {
}

// Stub for enterpriseTokens
type enterpriseTokens struct {
}

// enterpriseAgentToken CE stub
func (t *Store) enterpriseAgentToken() string {
	return ""
}

// loadEnterpriseTokens is a noop stub for the func defined agent_ent.go
func loadEnterpriseTokens(_ *Store, _ Config) {
}
