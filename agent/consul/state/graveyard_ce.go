// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/graveyard_ce.go
Graveyard Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/graveyard_ce.go> a code:Module ;
    code:name "agent/consul/state/graveyard_ce.go" ;
    code:description "Graveyard Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"fmt"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func (g *Graveyard) insertTombstoneWithTxn(tx WriteTxn, _ string, stone *Tombstone, updateMax bool) error {
	if err := tx.Insert("tombstones", stone); err != nil {
		return err
	}

	if updateMax {
		if err := indexUpdateMaxTxn(tx, stone.Index, "tombstones"); err != nil {
			return fmt.Errorf("failed updating tombstone index: %v", err)
		}
	} else {
		if err := tx.Insert(tableIndex, &IndexEntry{"tombstones", stone.Index}); err != nil {
			return fmt.Errorf("failed updating tombstone index: %s", err)
		}
	}
	return nil
}

// GetMaxIndexTxn returns the highest index tombstone whose key matches the
// given context, using a prefix match.
func (g *Graveyard) GetMaxIndexTxn(tx ReadTxn, prefix string, _ *acl.EnterpriseMeta) (uint64, error) {
	var lindex uint64
	q := Query{Value: prefix, EnterpriseMeta: *structs.DefaultEnterpriseMetaInDefaultPartition()}
	stones, err := tx.Get(tableTombstones, indexID+"_prefix", q)
	if err != nil {
		return 0, fmt.Errorf("failed querying tombstones: %s", err)
	}
	for stone := stones.Next(); stone != nil; stone = stones.Next() {
		s := stone.(*Tombstone)
		if s.Index > lindex {
			lindex = s.Index
		}
	}
	return lindex, nil
}
