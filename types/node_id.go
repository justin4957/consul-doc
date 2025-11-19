// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: types/node_id.go
Node identifier type definition for unique node identification across space and time.

## Linked Modules
- [agent/structs](../agent/structs/structs.go) - Node structures that use NodeID
- [agent/consul/state](../agent/consul/state/state_store.go) - State store node indexing
- [api](../api/api.go) - API node types

## Tags
types, data-model, nodes, identifiers

## Exports
NodeID

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<types/node_id.go> a code:Module ;
    code:language "go" ;
    code:name "types/node_id.go" ;
    code:description "Node identifier type definition for unique node identification across space and time" ;
    code:layer "api" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs/structs.go" ;
        code:relationship "Node structures that use NodeID"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state/state_store.go" ;
        code:relationship "State store node indexing"
    ], [
        code:name "api" ;
        code:path "../api/api.go" ;
        code:relationship "API node types"
    ] ;
    code:exports :NodeID ;
    code:tags "types", "data-model", "nodes", "identifiers" .
<!-- End LinkedDoc RDF -->
*/
package types

// NodeID is a unique identifier for a node across space and time.
type NodeID string
