// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/acl_ce.go
Acl Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
access-control, agent, authorization, security

## Exports
EnterpriseACLResolverDelegate

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/acl_ce.go> a code:Module ;
    code:name "agent/consul/acl_ce.go" ;
    code:description "Acl Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :EnterpriseACLResolverDelegate ;
    code:tags "access-control", "agent", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

// EnterpriseACLResolverDelegate stub
type EnterpriseACLResolverDelegate interface{}

func (s *Server) replicationEnterpriseMeta() *acl.EnterpriseMeta {
	return structs.ReplicationEnterpriseMeta()
}

func serverPartitionInfo(s *Server) acl.ExportFetcher {
	return &partitionInfoNoop{}
}

func newACLConfig(_ acl.ExportFetcher, _ hclog.Logger) *acl.Config {
	return &acl.Config{
		WildcardName: structs.WildcardSpecifier,
	}
}

func (r *ACLResolver) resolveEnterpriseDefaultsForIdentity(identity structs.ACLIdentity) (acl.Authorizer, error) {
	return nil, nil
}

// resolveEnterpriseIdentityAndRoles will resolve an enterprise identity to an additional set of roles
func (*ACLResolver) resolveEnterpriseIdentityAndRoles(_ structs.ACLIdentity) (structs.ACLIdentity, structs.ACLRoles, error) {
	// this function does nothing in CE
	return nil, nil, nil
}

// resolveEnterpriseIdentityAndPolicies will resolve an enterprise identity to an additional set of policies
func (*ACLResolver) resolveEnterpriseIdentityAndPolicies(_ structs.ACLIdentity) (structs.ACLIdentity, structs.ACLPolicies, error) {
	// this function does nothing in CE
	return nil, nil, nil
}

// resolveLocallyManagedEnterpriseToken will resolve a managed service provider token to an identity and authorizer
func (*ACLResolver) resolveLocallyManagedEnterpriseToken(_ string) (structs.ACLIdentity, acl.Authorizer, bool) {
	return nil, nil, false
}

func setEnterpriseConf(entMeta *acl.EnterpriseMeta, conf *acl.Config) {}
