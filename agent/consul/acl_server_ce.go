// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/acl_server_ce.go
Acl Server Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
access-control, agent, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/acl_server_ce.go> a code:Module ;
    code:name "agent/consul/acl_server_ce.go" ;
    code:description "Acl Server Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "access-control", "agent", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/agent/structs"
)

// Consul-enterprise only
func (s *Server) validateEnterpriseToken(identity structs.ACLIdentity) error {
	return nil
}

// aclBootstrapAllowed returns whether the server's configuration would allow ACL bootstrapping
//
// This endpoint does not take into account whether bootstrapping has been performed previously
// nor the bootstrap reset file.
func (s *Server) aclBootstrapAllowed() error {
	return nil
}

func (*Server) enterpriseAuthMethodTypeValidation(authMethodType string) error {
	return nil
}
