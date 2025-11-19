// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/acl_endpoint_ce.go
Acl Endpoint Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/consul/authmethod](../agent/consul/authmethod)
- [agent/consul/state](../agent/consul/state)
- [agent/structs](../agent/structs)

## Tags
access-control, agent, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/acl_endpoint_ce.go> a code:Module ;
    code:name "agent/consul/acl_endpoint_ce.go" ;
    code:description "Acl Endpoint Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/consul/authmethod" ;
        code:path "../agent/consul/authmethod"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "access-control", "agent", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	memdb "github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/consul/authmethod"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/structs"
)

func (a *ACL) tokenUpsertValidateEnterprise(token *structs.ACLToken, existing *structs.ACLToken) error {
	state := a.srv.fsm.State()
	return state.ACLTokenUpsertValidateEnterprise(token, existing)
}

func (a *ACL) policyUpsertValidateEnterprise(policy *structs.ACLPolicy, existing *structs.ACLPolicy) error {
	state := a.srv.fsm.State()
	return state.ACLPolicyUpsertValidateEnterprise(policy, existing)
}

func (a *ACL) roleUpsertValidateEnterprise(role *structs.ACLRole, existing *structs.ACLRole) error {
	state := a.srv.fsm.State()
	return state.ACLRoleUpsertValidateEnterprise(role, existing)
}

func enterpriseAuthMethodValidation(method *structs.ACLAuthMethod, validator authmethod.Validator) error {
	return nil
}

func getTokenNamespaceDefaults(ws memdb.WatchSet, state *state.Store, entMeta *acl.EnterpriseMeta) ([]string, []string, error) {
	return nil, nil, nil
}
