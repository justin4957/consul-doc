// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/config_entry_sameness_group_ce.go
Config Entry Sameness Group Ce module for agent layer

## Linked Modules
- [agent/configentry](../agent/configentry)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, persistence, storage

## Exports
SamenessGroupDefaultIndex

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/config_entry_sameness_group_ce.go> a code:Module ;
    code:name "agent/consul/state/config_entry_sameness_group_ce.go" ;
    code:description "Config Entry Sameness Group Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/configentry" ;
        code:path "../agent/configentry"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :SamenessGroupDefaultIndex ;
    code:tags "agent", "configuration", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"fmt"

	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/go-memdb"
)

// SamenessGroupDefaultIndex is a placeholder for CE. Sameness-groups are enterprise only.
type SamenessGroupDefaultIndex struct{}

var _ memdb.Indexer = (*SamenessGroupDefaultIndex)(nil)
var _ memdb.MultiIndexer = (*SamenessGroupDefaultIndex)(nil)

func (*SamenessGroupDefaultIndex) FromObject(obj interface{}) (bool, [][]byte, error) {
	return false, nil, nil
}

func (*SamenessGroupDefaultIndex) FromArgs(args ...interface{}) ([]byte, error) {
	return nil, nil
}

func checkSamenessGroup(tx ReadTxn, newConfig structs.ConfigEntry) error {
	return fmt.Errorf("sameness-groups are an enterprise-only feature")
}

// getExportedServicesConfigEntryTxn is a convenience method for fetching a
// sameness-group config entries.
//
// If an override KEY is present for the requested config entry, the index
// returned will be 0. Any override VALUE (nil or otherwise) will be returned
// if there is a KEY match.
func getSamenessGroupConfigEntryTxn(
	tx ReadTxn,
	ws memdb.WatchSet,
	name string,
	overrides map[configentry.KindName]structs.ConfigEntry,
	partition string,
) (uint64, *structs.SamenessGroupConfigEntry, error) {
	return 0, nil, nil
}

func getDefaultSamenessGroup(tx ReadTxn, ws memdb.WatchSet, partition string) (uint64, *structs.SamenessGroupConfigEntry, error) {
	return 0, nil, nil
}
