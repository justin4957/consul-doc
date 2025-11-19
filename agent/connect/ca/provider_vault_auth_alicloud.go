// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/ca/provider_vault_auth_alicloud.go
Provider Vault Auth Alicloud module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, authentication, mtls, security, service-mesh

## Exports
AliLoginDataGen, NewAliCloudAuthClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/ca/provider_vault_auth_alicloud.go> a code:Module ;
    code:name "agent/connect/ca/provider_vault_auth_alicloud.go" ;
    code:description "Provider Vault Auth Alicloud module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :AliLoginDataGen, :NewAliCloudAuthClient ;
    code:tags "agent", "authentication", "mtls", "security", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package ca

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials/providers"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/vault-plugin-auth-alicloud/tools"
)

func NewAliCloudAuthClient(authMethod *structs.VaultAuthMethod) (*VaultAuthClient, error) {
	params := authMethod.Params
	authClient := NewVaultAPIAuthClient(authMethod, "")
	// check for login data already in params (for backwards compability)
	legacyKeys := []string{"access_key", "secret_key", "access_token"}
	if legacyCheck(params, legacyKeys...) {
		return authClient, nil
	}

	if r, ok := params["role"].(string); !ok || r == "" {
		return nil, fmt.Errorf("role is required for AliCloud login")
	}
	if r, ok := params["region"].(string); !ok || r == "" {
		return nil, fmt.Errorf("region is required for AliCloud login")
	}
	client := NewVaultAPIAuthClient(authMethod, "")
	client.LoginDataGen = AliLoginDataGen
	return client, nil
}

func AliLoginDataGen(authMethod *structs.VaultAuthMethod) (map[string]any, error) {
	// validity of these params is checked above in New..
	role := authMethod.Params["role"].(string)
	region := authMethod.Params["region"].(string)
	// Credentials can be provided either explicitly via env vars,
	// or we will try to derive them from instance metadata.
	credentialChain := []providers.Provider{
		providers.NewEnvCredentialProvider(),
		providers.NewInstanceMetadataProvider(),
	}
	creds, err := providers.NewChainProvider(credentialChain).Retrieve()
	if err != nil {
		return nil, err
	}

	loginData, err := tools.GenerateLoginData(role, creds, region)
	if err != nil {
		return nil, err
	}

	return loginData, nil
}
