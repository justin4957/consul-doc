// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/usage_ce.go
Usage Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, persistence, storage

## Exports
EnterpriseConfigEntryUsage, EnterpriseKVUsage, EnterpriseNodeUsage, EnterprisePeeringUsage, EnterpriseServiceUsage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/usage_ce.go> a code:Module ;
    code:name "agent/consul/state/usage_ce.go" ;
    code:description "Usage Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :EnterpriseConfigEntryUsage, :EnterpriseKVUsage, :EnterpriseNodeUsage, :EnterprisePeeringUsage, :EnterpriseServiceUsage ;
    code:tags "agent", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/structs"
)

type EnterpriseServiceUsage struct{}
type EnterpriseNodeUsage struct{}
type EnterprisePeeringUsage struct{}
type EnterpriseKVUsage struct{}
type EnterpriseConfigEntryUsage struct{}

func addEnterpriseNodeUsage(map[string]int, memdb.Change) {}

func addEnterprisePeeringUsage(map[string]int, memdb.Change) {}

func addEnterpriseServiceInstanceUsage(map[string]int, memdb.Change) {}

func addEnterpriseServiceUsage(map[string]int, map[structs.ServiceName]uniqueServiceState) {}

func addEnterpriseConnectServiceInstanceUsage(map[string]int, *structs.ServiceNode, int) {}

func addEnterpriseBillableServiceInstanceUsage(map[string]int, *structs.ServiceNode, int) {}

func addEnterpriseKVUsage(map[string]int, memdb.Change) {}

func addEnterpriseConfigEntryUsage(map[string]int, memdb.Change) {}

func compileEnterpriseServiceUsage(ws memdb.WatchSet, tx ReadTxn, usage structs.ServiceUsage) (structs.ServiceUsage, error) {
	return usage, nil
}

func compileEnterpriseNodeUsage(tx ReadTxn, usage NodeUsage) (NodeUsage, error) {
	return usage, nil
}

func compileEnterprisePeeringUsage(tx ReadTxn, usage PeeringUsage) (PeeringUsage, error) {
	return usage, nil
}

func compileEnterpriseKVUsage(tx ReadTxn, usage KVUsage) (KVUsage, error) {
	return usage, nil
}

func compileEnterpriseConfigEntryUsage(tx ReadTxn, usage ConfigEntryUsage) (ConfigEntryUsage, error) {
	return usage, nil
}
