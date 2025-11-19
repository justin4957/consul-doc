// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/peering_ce.go
Peering Ce module for agent layer

## Linked Modules
- [proto/private/pbpeering](../proto/private/pbpeering)

## Tags
agent, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/peering_ce.go> a code:Module ;
    code:name "agent/consul/state/peering_ce.go" ;
    code:description "Peering Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "proto/private/pbpeering" ;
        code:path "../proto/private/pbpeering"
    ] ;
    code:tags "agent", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/proto/private/pbpeering"
)

func indexPeeringFromQuery(q Query) ([]byte, error) {
	var b indexBuilder
	b.String(strings.ToLower(q.Value))
	return b.Bytes(), nil
}

func indexFromPeering(p *pbpeering.Peering) ([]byte, error) {
	if p.Name == "" {
		return nil, errMissingValueForIndex
	}

	var b indexBuilder
	b.String(strings.ToLower(p.Name))
	return b.Bytes(), nil
}

func indexFromPeeringTrustBundle(ptb *pbpeering.PeeringTrustBundle) ([]byte, error) {
	if ptb.PeerName == "" {
		return nil, errMissingValueForIndex
	}

	var b indexBuilder
	b.String(strings.ToLower(ptb.PeerName))
	return b.Bytes(), nil
}

func updatePeeringTableIndexes(tx WriteTxn, idx uint64, _ string) error {
	if err := tx.Insert(tableIndex, &IndexEntry{Key: tablePeering, Value: idx}); err != nil {
		return fmt.Errorf("failed updating table index: %w", err)
	}
	return nil
}

func updatePeeringTrustBundlesTableIndexes(tx WriteTxn, idx uint64, _ string) error {
	if err := tx.Insert(tableIndex, &IndexEntry{Key: tablePeeringTrustBundles, Value: idx}); err != nil {
		return fmt.Errorf("failed updating table index: %w", err)
	}
	return nil
}
