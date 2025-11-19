// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/services/resource/testing/testing.go
Testing module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/structs](../agent/structs)

## Tags
agent, api, communication, grpc, networking

## Exports
AuthorizerFrom

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/resource/testing/testing.go> a code:Module ;
    code:name "agent/grpc-external/services/resource/testing/testing.go" ;
    code:description "Testing module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :AuthorizerFrom ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testing

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hashicorp/go-uuid"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/structs"
)

func randomACLIdentity(t *testing.T) structs.ACLIdentity {
	id, err := uuid.GenerateUUID()
	require.NoError(t, err)

	return &structs.ACLToken{AccessorID: id}
}

func AuthorizerFrom(t *testing.T, policyStrs ...string) resolver.Result {
	policies := []*acl.Policy{}
	for _, policyStr := range policyStrs {
		policy, err := acl.NewPolicyFromSource(policyStr, nil, nil)
		require.NoError(t, err)
		policies = append(policies, policy)
	}

	authz, err := acl.NewPolicyAuthorizerWithDefaults(acl.DenyAll(), policies, nil)
	require.NoError(t, err)

	return resolver.Result{
		Authorizer:  authz,
		ACLIdentity: randomACLIdentity(t),
	}
}
