// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/catalog_events_ce.go
Catalog Events Ce module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, discovery, persistence, registry, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/catalog_events_ce.go> a code:Module ;
    code:name "agent/consul/state/catalog_events_ce.go" ;
    code:description "Catalog Events Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "discovery", "persistence", "registry", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"strings"

	"github.com/hashicorp/consul/agent/structs"
)

func (nst nodeServiceTuple) nodeTuple() nodeTuple {
	return nodeTuple{
		Node:      strings.ToLower(nst.Node),
		Partition: "",
		PeerName:  nst.PeerName,
	}
}

func newNodeTupleFromNode(node *structs.Node) nodeTuple {
	return nodeTuple{
		Node:      strings.ToLower(node.Node),
		Partition: "",
		PeerName:  node.PeerName,
	}
}

func newNodeTupleFromHealthCheck(hc *structs.HealthCheck) nodeTuple {
	return nodeTuple{
		Node:      strings.ToLower(hc.Node),
		Partition: "",
		PeerName:  hc.PeerName,
	}
}

// String satisfies the stream.Subject interface.
func (s EventSubjectService) String() string {
	key := s.Key
	if v := s.overrideKey; v != "" {
		key = v
	}
	key = strings.ToLower(key)

	if s.PeerName == "" {
		return key
	}
	return s.PeerName + "/" + key
}
