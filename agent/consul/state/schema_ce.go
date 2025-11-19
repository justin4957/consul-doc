// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/schema_ce.go
Schema Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/schema_ce.go> a code:Module ;
    code:name "agent/consul/state/schema_ce.go" ;
    code:description "Schema Ce module for agent layer" ;
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

func partitionedIndexEntryName(entry string, _ string) string {
	return entry
}

func partitionedAndNamespacedIndexEntryName(entry string, _ *acl.EnterpriseMeta) string {
	return entry
}

// peeredIndexEntryName returns the peered index key for an importable entity (e.g. checks, services, or nodes).
func peeredIndexEntryName(entry, peerName string) string {
	if peerName == "" {
		peerName = structs.LocalPeerKeyword
	}
	return fmt.Sprintf("peer.%s:%s", peerName, entry)
}
