// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Package configentry contains structs and logic related to the Configuration
// Entry subsystem. Currently this is restricted to structs used during
// runtime, but which are not serialized to the network or disk.

/*
# Module: agent/configentry/doc.go
Doc module for agent layer

## Tags
agent, configuration

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/configentry/doc.go> a code:Module ;
    code:name "agent/configentry/doc.go" ;
    code:description "Doc module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package configentry
