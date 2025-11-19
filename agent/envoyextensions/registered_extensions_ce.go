// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/envoyextensions/registered_extensions_ce.go
Registered Extensions Ce module for agent layer

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/envoyextensions/registered_extensions_ce.go> a code:Module ;
    code:name "agent/envoyextensions/registered_extensions_ce.go" ;
    code:description "Registered Extensions Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package envoyextensions

var enterpriseExtensionConstructors = map[string]extensionConstructor{}
