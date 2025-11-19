// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/acl_authmethod.go
Acl Authmethod module for agent layer

## Linked Modules
- [agent/consul/authmethod](../agent/consul/authmethod)
- [agent/consul/authmethod/awsauth](../agent/consul/authmethod/awsauth)
- [agent/consul/authmethod/kubeauth](../agent/consul/authmethod/kubeauth)
- [agent/consul/authmethod/ssoauth](../agent/consul/authmethod/ssoauth)
- [agent/structs](../agent/structs)

## Tags
access-control, agent, authentication, authorization, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/acl_authmethod.go> a code:Module ;
    code:name "agent/consul/acl_authmethod.go" ;
    code:description "Acl Authmethod module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/authmethod" ;
        code:path "../agent/consul/authmethod"
    ], [
        code:name "agent/consul/authmethod/awsauth" ;
        code:path "../agent/consul/authmethod/awsauth"
    ], [
        code:name "agent/consul/authmethod/kubeauth" ;
        code:path "../agent/consul/authmethod/kubeauth"
    ], [
        code:name "agent/consul/authmethod/ssoauth" ;
        code:path "../agent/consul/authmethod/ssoauth"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "access-control", "agent", "authentication", "authorization", "security" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"

	"github.com/hashicorp/consul/agent/consul/authmethod"
	"github.com/hashicorp/consul/agent/structs"

	// register these as a builtin auth method
	_ "github.com/hashicorp/consul/agent/consul/authmethod/awsauth"
	_ "github.com/hashicorp/consul/agent/consul/authmethod/kubeauth"
	_ "github.com/hashicorp/consul/agent/consul/authmethod/ssoauth"
)

type authMethodValidatorEntry struct {
	Validator   authmethod.Validator
	ModifyIndex uint64 // the raft index when this last changed
}

// loadAuthMethodValidator returns an authmethod.Validator for the given auth
// method configuration. If the cache is up to date as-of the provided index
// then the cached version is returned, otherwise a new validator is created
// and cached.
func (s *Server) loadAuthMethodValidator(idx uint64, method *structs.ACLAuthMethod) (authmethod.Validator, error) {
	if prevIdx, v, ok := s.aclAuthMethodValidators.GetValidator(method); ok && idx <= prevIdx {
		return v, nil
	}

	v, err := authmethod.NewValidator(s.logger, method)
	if err != nil {
		return nil, fmt.Errorf("auth method validator for %q could not be initialized: %v", method.Name, err)
	}

	v = s.aclAuthMethodValidators.PutValidatorIfNewer(method, v, idx)

	return v, nil
}
