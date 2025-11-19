// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/auto-config/auto_config_ce.go
Auto Config Ce module for agent layer

## Tags
agent, configuration

## Exports
AutoConfigEnterprise

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/auto-config/auto_config_ce.go> a code:Module ;
    code:name "agent/auto-config/auto_config_ce.go" ;
    code:description "Auto Config Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :AutoConfigEnterprise ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package autoconf

// AutoConfigEnterprise has no fields in CE
type AutoConfigEnterprise struct{}

// newAutoConfigEnterprise initializes the enterprise AutoConfig struct
func newAutoConfigEnterprise(config Config) AutoConfigEnterprise {
	return AutoConfigEnterprise{}
}
