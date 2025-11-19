// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: command/catalog/helpers_ce.go
Helpers Ce module for cli layer

## Linked Modules
- [api](../api)

## Tags
cli, discovery, registry, user-interface

## Exports
NodeRow, NodesHeader

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/catalog/helpers_ce.go> a code:Module ;
    code:name "command/catalog/helpers_ce.go" ;
    code:description "Helpers Ce module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :NodeRow, :NodesHeader ;
    code:tags "cli", "discovery", "registry", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package catalog

import (
	"fmt"
	"strings"

	"github.com/hashicorp/consul/api"
)

func NodesHeader(isDetailed bool) string {
	if isDetailed {
		return "Node\x1fID\x1fAddress\x1fDC\x1fTaggedAddresses\x1fMeta"
	} else {
		return "Node\x1fID\x1fAddress\x1fDC"
	}
}

func NodeRow(node *api.Node, isDetailed bool) string {
	if isDetailed {
		return fmt.Sprintf("%s\x1f%s\x1f%s\x1f%s\x1f%s\x1f%s",
			node.Node, node.ID, node.Address, node.Datacenter,
			mapToKV(node.TaggedAddresses, ", "), mapToKV(node.Meta, ", "))
	} else {
		// Shorten the ID in non-detailed mode to just the first octet.
		id := node.ID
		idx := strings.Index(id, "-")
		if idx > 0 {
			id = id[0:idx]
		}
		return fmt.Sprintf("%s\x1f%s\x1f%s\x1f%s",
			node.Node, id, node.Address, node.Datacenter)
	}
}
