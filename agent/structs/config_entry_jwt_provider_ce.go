// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_jwt_provider_ce.go
Config Entry Jwt Provider Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, configuration, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_jwt_provider_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_jwt_provider_ce.go" ;
    code:description "Config Entry Jwt Provider Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "configuration", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

func (e *JWTProviderConfigEntry) validatePartitionAndNamespace() error {
	if !acl.IsDefaultPartition(e.PartitionOrDefault()) {
		return fmt.Errorf("Partitions are an enterprise only feature")
	}

	if acl.DefaultNamespaceName != e.NamespaceOrDefault() {
		return fmt.Errorf("Namespaces are an enterprise only feature")
	}

	return nil
}
