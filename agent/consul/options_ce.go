// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/options_ce.go
Options Ce module for agent layer

## Tags
agent

## Exports
EnterpriseDeps

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/options_ce.go> a code:Module ;
    code:name "agent/consul/options_ce.go" ;
    code:description "Options Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EnterpriseDeps ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

type EnterpriseDeps struct{}
