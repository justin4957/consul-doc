// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/tenancy_bridge_ce.go
Tenancy Bridge Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/tenancy_bridge_ce.go> a code:Module ;
    code:name "agent/consul/tenancy_bridge_ce.go" ;
    code:description "Tenancy Bridge Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

func (b *V1TenancyBridge) PartitionExists(partition string) (bool, error) {
	if partition == "default" {
		return true, nil
	}
	return false, nil
}

func (b *V1TenancyBridge) IsPartitionMarkedForDeletion(partition string) (bool, error) {
	return false, nil
}

func (b *V1TenancyBridge) NamespaceExists(partition, namespace string) (bool, error) {
	if partition == "default" && namespace == "default" {
		return true, nil
	}
	return false, nil
}

func (b *V1TenancyBridge) IsNamespaceMarkedForDeletion(partition, namespace string) (bool, error) {
	return false, nil
}
