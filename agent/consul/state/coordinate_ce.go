// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/coordinate_ce.go
Coordinate Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/coordinate_ce.go> a code:Module ;
    code:name "agent/consul/state/coordinate_ce.go" ;
    code:description "Coordinate Ce module for agent layer" ;
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

func coordinatesMaxIndex(tx ReadTxn, entMeta *acl.EnterpriseMeta) uint64 {
	return maxIndexTxn(tx, tableCoordinates)
}

func updateCoordinatesIndexes(tx WriteTxn, idx uint64, entMeta *acl.EnterpriseMeta) error {
	// Update the index.
	if err := indexUpdateMaxTxn(tx, idx, tableCoordinates); err != nil {
		return fmt.Errorf("failed updating index: %s", err)
	}
	return nil
}

func ensureCoordinateTxn(tx WriteTxn, idx uint64, coord *structs.Coordinate) error {
	// ensure that the Partition is always empty within the state store
	coord.Partition = ""

	if err := tx.Insert(tableCoordinates, coord); err != nil {
		return fmt.Errorf("failed inserting coordinate: %s", err)
	}

	if err := updateCoordinatesIndexes(tx, idx, coord.GetEnterpriseMeta()); err != nil {
		return fmt.Errorf("failed updating coordinate index: %s", err)
	}

	return nil
}
