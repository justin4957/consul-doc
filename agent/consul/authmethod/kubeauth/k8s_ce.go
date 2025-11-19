// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/authmethod/kubeauth/k8s_ce.go
K8S Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, authentication, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/authmethod/kubeauth/k8s_ce.go> a code:Module ;
    code:name "agent/consul/authmethod/kubeauth/k8s_ce.go" ;
    code:description "K8S Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "authentication", "security" .
<!-- End LinkedDoc RDF -->
*/
package kubeauth

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

type enterpriseConfig struct{}

func enterpriseValidation(method *structs.ACLAuthMethod, config *Config) error {
	return nil
}

func (v *Validator) k8sEntMetaFromFields(fields map[string]string) *acl.EnterpriseMeta {
	return nil
}
