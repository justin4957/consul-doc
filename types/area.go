// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: types/area.go
Area module for internal layer

## Tags
data-model, internal, types

## Exports
AreaID, AreaLAN, AreaWAN

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<types/area.go> a code:Module ;
    code:name "types/area.go" ;
    code:description "Area module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :AreaID, :AreaLAN, :AreaWAN ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package types

// AreaID is a strongly-typed string used to uniquely represent a network area,
// which is a relationship between Consul servers.
type AreaID string

// This represents the existing WAN area that's built in to Consul. Consul
// Enterprise generalizes areas, which are represented with UUIDs.
const AreaWAN AreaID = "wan"

// This represents the existing LAN area that's built in to Consul. Consul
// Enterprise generalizes areas, which are represented with UUIDs.
const AreaLAN AreaID = "lan"
