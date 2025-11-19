// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: test-integ/topoutil/naming_shim.go
Naming Shim module for internal layer

## Linked Modules
- [testing/deployer/topology](../testing/deployer/topology)

## Tags
internal

## Exports
NewBlankspaceServiceWithDefaults, NewFortioServiceWithDefaults

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test-integ/topoutil/naming_shim.go> a code:Module ;
    code:name "test-integ/topoutil/naming_shim.go" ;
    code:description "Naming Shim module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "testing/deployer/topology" ;
        code:path "../testing/deployer/topology"
    ] ;
    code:exports :NewBlankspaceServiceWithDefaults, :NewFortioServiceWithDefaults ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package topoutil

import (
	"github.com/hashicorp/consul/testing/deployer/topology"
)

// Deprecated: NewFortioWorkloadWithDefaults
func NewFortioServiceWithDefaults(
	cluster string,
	sid topology.ID,
	mut func(*topology.Workload),
) *topology.Workload {
	return NewFortioWorkloadWithDefaults(cluster, sid, mut)
}

// Deprecated: NewBlankspaceWorkloadWithDefaults
func NewBlankspaceServiceWithDefaults(
	cluster string,
	sid topology.ID,
	mut func(*topology.Workload),
) *topology.Workload {
	return NewBlankspaceWorkloadWithDefaults(cluster, sid, mut)
}
