// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/config/runtime_ce.go
Runtime Ce module for agent layer

## Tags
agent, configuration

## Exports
EnterpriseRuntimeConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/runtime_ce.go> a code:Module ;
    code:name "agent/config/runtime_ce.go" ;
    code:description "Runtime Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EnterpriseRuntimeConfig ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

type EnterpriseRuntimeConfig struct{}

func (c *RuntimeConfig) PartitionOrEmpty() string   { return "" }
func (c *RuntimeConfig) PartitionOrDefault() string { return "" }
