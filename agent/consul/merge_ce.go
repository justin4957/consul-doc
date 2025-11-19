// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/merge_ce.go
Merge Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/merge_ce.go> a code:Module ;
    code:name "agent/consul/merge_ce.go" ;
    code:description "Merge Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"

	"github.com/hashicorp/serf/serf"
)

func (md *lanMergeDelegate) enterpriseNotifyMergeMember(m *serf.Member) error {
	if memberFIPS := m.Tags["fips"]; memberFIPS != "" {
		return fmt.Errorf("Member '%s' is FIPS Consul; FIPS Consul is only available in Consul Enterprise",
			m.Name)
	}
	if memberPartition := m.Tags["ap"]; memberPartition != "" {
		return fmt.Errorf("Member '%s' part of partition '%s'; Partitions are a Consul Enterprise feature",
			m.Name, memberPartition)
	}
	if segment := m.Tags["segment"]; segment != "" {
		return fmt.Errorf("Member '%s' part of segment '%s'; Network Segments are a Consul Enterprise feature",
			m.Name, segment)
	}
	return nil
}
