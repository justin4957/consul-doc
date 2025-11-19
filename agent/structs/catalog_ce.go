// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/catalog_ce.go
Catalog Ce module for agent layer

## Tags
agent, data-model, discovery, registry, types

## Exports
IsConsulServiceID, IsSerfCheckID

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/catalog_ce.go> a code:Module ;
    code:name "agent/structs/catalog_ce.go" ;
    code:description "Catalog Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :IsConsulServiceID, :IsSerfCheckID ;
    code:tags "agent", "data-model", "discovery", "registry", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

func IsConsulServiceID(id ServiceID) bool {
	return id.ID == ConsulServiceID
}

func IsSerfCheckID(id CheckID) bool {
	return id.ID == SerfCheckID
}
