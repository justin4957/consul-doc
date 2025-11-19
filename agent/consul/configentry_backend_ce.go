// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/configentry_backend_ce.go
Configentry Backend Ce module for agent layer

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/configentry_backend_ce.go> a code:Module ;
    code:name "agent/consul/configentry_backend_ce.go" ;
    code:description "Configentry Backend Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"
	"strings"
)

func (b *ConfigEntryBackend) enterpriseCheckPartitions(partition string) error {
	if partition == "" || strings.EqualFold(partition, "default") {
		return nil
	}
	return fmt.Errorf("Partitions are a Consul Enterprise feature")
}
