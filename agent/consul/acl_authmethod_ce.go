// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/acl_authmethod_ce.go
Acl Authmethod Ce module for agent layer

## Tags
access-control, agent, authentication, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/acl_authmethod_ce.go> a code:Module ;
    code:name "agent/consul/acl_authmethod_ce.go" ;
    code:description "Acl Authmethod Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "access-control", "agent", "authentication", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package consul

func (s *Server) enterpriseEvaluateRoleBindings() error {
	return nil
}
