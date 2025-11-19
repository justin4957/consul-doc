// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_intentions_ce.go
Config Entry Intentions Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, configuration, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_intentions_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_intentions_ce.go" ;
    code:description "Config Entry Intentions Ce module for agent layer" ;
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

func validateSourceIntentionEnterpriseMeta(_, _ *acl.EnterpriseMeta) error {
	return nil
}

func (s *SourceIntention) validateSamenessGroup() error {
	if s.SamenessGroup != "" {
		return fmt.Errorf("Sameness groups are a Consul Enterprise feature.")
	}

	return nil
}
