// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/ca/provider_vault_auth_gcp.go
Provider Vault Auth Gcp module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, authentication, mtls, security, service-mesh

## Exports
NewGCPAuthClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/ca/provider_vault_auth_gcp.go> a code:Module ;
    code:name "agent/connect/ca/provider_vault_auth_gcp.go" ;
    code:description "Provider Vault Auth Gcp module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :NewGCPAuthClient ;
    code:tags "agent", "authentication", "mtls", "security", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package ca

import (
	"fmt"

	"github.com/hashicorp/vault/api/auth/gcp"

	"github.com/hashicorp/consul/agent/structs"
)

var _ VaultAuthenticator = (*gcp.GCPAuth)(nil)

// NewGCPAuthClient returns a VaultAuthenticator that can log into Vault using the GCP auth method.
func NewGCPAuthClient(authMethod *structs.VaultAuthMethod) (VaultAuthenticator, error) {
	// Check if the configuration already contains a JWT auth token. If so we want to
	// perform a direct request to the login API with the config that is provided.
	// This supports the  Vault CA config in a backwards compatible way so that we don't
	// break existing configurations.
	if legacyCheck(authMethod.Params, "jwt") {
		return NewVaultAPIAuthClient(authMethod, ""), nil
	}

	// Create a GCPAuth client from the CA config
	params, err := toMapStringString(authMethod.Params)
	if err != nil {
		return nil, fmt.Errorf("misconfiguration of GCP auth parameters: %w", err)
	}

	opts := []gcp.LoginOption{gcp.WithMountPath(authMethod.MountPath)}

	// Handle the login type explicitly.
	switch params["type"] {
	case "gce":
		opts = append(opts, gcp.WithGCEAuth())
	default:
		opts = append(opts, gcp.WithIAMAuth(params["service_account_email"]))
	}

	auth, err := gcp.NewGCPAuth(params["role"], opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new Vault GCP auth client: %w", err)
	}
	return auth, nil
}
