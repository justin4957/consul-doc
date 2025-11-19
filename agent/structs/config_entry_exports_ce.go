// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_exports_ce.go
Config Entry Exports Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, configuration, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_exports_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_exports_ce.go" ;
    code:description "Config Entry Exports Ce module for agent layer" ;
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

func (e *ExportedServicesConfigEntry) validateServicesEnterprise() error {
	for i, svc := range e.Services {
		for j, consumer := range svc.Consumers {
			if !acl.IsDefaultPartition(consumer.Partition) {
				return fmt.Errorf("Services[%d].Consumers[%d]: partitions are an enterprise-only feature", i, j)
			}
			if consumer.SamenessGroup != "" {
				return fmt.Errorf("Services[%d].Consumers[%d]: sameness-groups are an enterprise-only feature", i, j)
			}
		}
	}
	return nil
}
