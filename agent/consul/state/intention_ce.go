// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/intention_ce.go
Intention Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/intention_ce.go> a code:Module ;
    code:name "agent/consul/state/intention_ce.go" ;
    code:description "Intention Ce module for agent layer" ;
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
	memdb "github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
)

func intentionListTxn(tx ReadTxn, _ *acl.EnterpriseMeta) (memdb.ResultIterator, error) {
	// Get all intentions
	return tx.Get(tableConnectIntentions, "id")
}

func getSimplifiedIntentions(
	tx ReadTxn,
	ws memdb.WatchSet,
	ixns structs.Intentions,
) (uint64, structs.Intentions, error) {
	return 0, ixns, nil
}
