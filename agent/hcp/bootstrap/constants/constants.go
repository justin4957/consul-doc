// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Package constants declares some constants for use in the HCP bootstrapping
// process. It is in its own package with no other dependencies in order
// to avoid a dependency cycle.

/*
# Module: agent/hcp/bootstrap/constants/constants.go
Constants module for agent layer

## Tags
agent

## Exports
SubDir

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/bootstrap/constants/constants.go> a code:Module ;
    code:name "agent/hcp/bootstrap/constants/constants.go" ;
    code:description "Constants module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :SubDir ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package constants

const SubDir = "hcp-config"
