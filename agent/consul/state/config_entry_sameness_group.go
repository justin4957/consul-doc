// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/state/config_entry_sameness_group.go
Config Entry Sameness Group module for agent layer

## Linked Modules
- [agent/configentry](../agent/configentry)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/config_entry_sameness_group.go> a code:Module ;
    code:name "agent/consul/state/config_entry_sameness_group.go" ;
    code:description "Config Entry Sameness Group module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/configentry" ;
        code:path "../agent/configentry"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "configuration", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/go-memdb"
)

// GetSamenessGroup returns a SamenessGroupConfigEntry from the state
// store using the provided parameters.
func (s *Store) GetSamenessGroup(ws memdb.WatchSet,
	name string,
	overrides map[configentry.KindName]structs.ConfigEntry,
	partition string) (uint64, *structs.SamenessGroupConfigEntry, error) {
	tx := s.db.ReadTxn()
	defer tx.Abort()

	return getSamenessGroupConfigEntryTxn(tx, ws, name, overrides, partition)
}
