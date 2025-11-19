// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/peering_backend_ce.go
Peering Backend Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/peering_backend_ce.go> a code:Module ;
    code:name "agent/consul/peering_backend_ce.go" ;
    code:description "Peering Backend Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"
	"strings"
)

func (b *PeeringBackend) enterpriseCheckPartitions(partition string) error {
	if partition == "" || strings.EqualFold(partition, "default") {
		return nil
	}
	return fmt.Errorf("Partitions are a Consul Enterprise feature")
}

func (b *PeeringBackend) enterpriseCheckNamespaces(namespace string) error {
	if namespace == "" || strings.EqualFold(namespace, "default") {
		return nil
	}
	return fmt.Errorf("Namespaces are a Consul Enterprise feature")
}
