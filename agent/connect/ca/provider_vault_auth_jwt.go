// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/ca/provider_vault_auth_jwt.go
Provider Vault Auth Jwt module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, authentication, mtls, security, service-mesh

## Exports
JwtLoginDataGen, NewJwtAuthClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/ca/provider_vault_auth_jwt.go> a code:Module ;
    code:name "agent/connect/ca/provider_vault_auth_jwt.go" ;
    code:description "Provider Vault Auth Jwt module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :JwtLoginDataGen, :NewJwtAuthClient ;
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

func NewJwtAuthClient(authMethod *structs.VaultAuthMethod) (*VaultAuthClient, error) {
	params := authMethod.Params

	role, ok := params["role"].(string)
	if !ok || strings.TrimSpace(role) == "" {
		return nil, fmt.Errorf("missing 'role' value")
	}

	authClient := NewVaultAPIAuthClient(authMethod, "")
	if legacyCheck(params, "jwt") {
		return authClient, nil
	}

	// The path is required for the auto-auth config, but this auth provider
	// seems to be used for jwt based auth by directly passing the jwt token.
	// So we only require the token file path if the token string isn't
	// present.
	tokenPath, ok := params["path"].(string)
	if !ok || strings.TrimSpace(tokenPath) == "" {
		return nil, fmt.Errorf("missing 'path' value")
	}
	authClient.LoginDataGen = JwtLoginDataGen
	return authClient, nil
}

func JwtLoginDataGen(authMethod *structs.VaultAuthMethod) (map[string]any, error) {
	params := authMethod.Params
	role := params["role"].(string)

	tokenPath := params["path"].(string)
	rawToken, err := os.ReadFile(tokenPath)
	if err != nil {
		return nil, err
	}

	return map[string]any{
		"role": role,
		"jwt":  strings.TrimSpace(string(rawToken)),
	}, nil
}
