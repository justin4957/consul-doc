// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_mesh_ce.go
Config Entry Mesh Ce module for agent layer

## Tags
agent, configuration, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_mesh_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_mesh_ce.go" ;
    code:description "Config Entry Mesh Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

func (e *MeshConfigEntry) validateEnterpriseMeta() error {
	return nil
}
