// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/acl_templated_policy_ce.go
Acl Templated Policy Ce module for agent layer

## Tags
access-control, agent, authorization, data-model, security, types

## Exports
ACLTemplatedPolicyAPIGateway, ACLTemplatedPolicyDNS, ACLTemplatedPolicyNode, ACLTemplatedPolicyNomadClient, ACLTemplatedPolicyNomadServer, ACLTemplatedPolicyService

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/acl_templated_policy_ce.go> a code:Module ;
    code:name "agent/structs/acl_templated_policy_ce.go" ;
    code:description "Acl Templated Policy Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :ACLTemplatedPolicyAPIGateway, :ACLTemplatedPolicyDNS, :ACLTemplatedPolicyNode, :ACLTemplatedPolicyNomadClient, :ACLTemplatedPolicyNomadServer, :ACLTemplatedPolicyService ;
    code:tags "access-control", "agent", "authorization", "data-model", "security", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import _ "embed"

//go:embed acltemplatedpolicy/policies/ce/service.hcl
var ACLTemplatedPolicyService string

//go:embed acltemplatedpolicy/policies/ce/node.hcl
var ACLTemplatedPolicyNode string

//go:embed acltemplatedpolicy/policies/ce/dns.hcl
var ACLTemplatedPolicyDNS string

//go:embed acltemplatedpolicy/policies/ce/nomad-server.hcl
var ACLTemplatedPolicyNomadServer string

//go:embed acltemplatedpolicy/policies/ce/api-gateway.hcl
var ACLTemplatedPolicyAPIGateway string

//go:embed acltemplatedpolicy/policies/ce/nomad-client.hcl
var ACLTemplatedPolicyNomadClient string

func (t *ACLToken) TemplatedPolicyList() []*ACLTemplatedPolicy {
	if len(t.TemplatedPolicies) == 0 {
		return nil
	}

	out := make([]*ACLTemplatedPolicy, 0, len(t.TemplatedPolicies))
	for _, n := range t.TemplatedPolicies {
		out = append(out, n.Clone())
	}
	return out
}

func (t *ACLRole) TemplatedPolicyList() []*ACLTemplatedPolicy {
	if len(t.TemplatedPolicies) == 0 {
		return nil
	}

	out := make([]*ACLTemplatedPolicy, 0, len(t.TemplatedPolicies))
	for _, n := range t.TemplatedPolicies {
		out = append(out, n.Clone())
	}
	return out
}
