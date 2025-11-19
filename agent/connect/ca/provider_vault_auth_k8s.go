// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/ca/provider_vault_auth_k8s.go
Provider Vault Auth K8S module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, authentication, mtls, security, service-mesh

## Exports
K8sLoginDataGen, NewK8sAuthClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/ca/provider_vault_auth_k8s.go> a code:Module ;
    code:name "agent/connect/ca/provider_vault_auth_k8s.go" ;
    code:description "Provider Vault Auth K8S module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :K8sLoginDataGen, :NewK8sAuthClient ;
    code:tags "agent", "authentication", "mtls", "security", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package ca

import (
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/consul/agent/structs"
)

func NewK8sAuthClient(authMethod *structs.VaultAuthMethod) (*VaultAuthClient, error) {
	params := authMethod.Params
	role, ok := params["role"].(string)
	if !ok || strings.TrimSpace(role) == "" {
		return nil, fmt.Errorf("missing 'role' value")
	}
	// don't check for `token_path` as it is optional

	authClient := NewVaultAPIAuthClient(authMethod, "")
	// Note the `jwt` can be passed directly in the authMethod as a Param value
	// is a freeform map in the config where they could hardcode it.
	if legacyCheck(params, "jwt") {
		return authClient, nil
	}

	authClient.LoginDataGen = K8sLoginDataGen
	return authClient, nil
}

func K8sLoginDataGen(authMethod *structs.VaultAuthMethod) (map[string]any, error) {
	params := authMethod.Params
	role := params["role"].(string)

	// read token from file on path
	tokenPath, ok := params["token_path"].(string)
	if !ok || strings.TrimSpace(tokenPath) == "" {
		tokenPath = defaultK8SServiceAccountTokenPath
	}
	rawToken, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}
	return map[string]any{
		"role": role,
		"jwt":  strings.TrimSpace(string(rawToken)),
	}, nil
}
