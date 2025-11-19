// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/serf_filter.go
Serf Filter module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent

## Exports
LANMemberFilter

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/serf_filter.go> a code:Module ;
    code:name "agent/consul/serf_filter.go" ;
    code:description "Serf Filter module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:exports :LANMemberFilter ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
)

type LANMemberFilter struct {
	Partition   string
	Segment     string
	AllSegments bool
}

func (f LANMemberFilter) Validate() error {
	if f.AllSegments && f.Segment != "" {
		return fmt.Errorf("cannot specify both allSegments and segment filters")
	}
	if (f.AllSegments || f.Segment != "") && !acl.IsDefaultPartition(f.Partition) {
		return fmt.Errorf("segments do not exist outside of the default partition")
	}
	return nil
}

func (f LANMemberFilter) PartitionOrDefault() string {
	return acl.PartitionOrDefault(f.Partition)
}
