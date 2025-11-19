// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/config/default_ce.go
Default Ce module for agent layer

## Tags
agent, configuration

## Exports
DefaultEnterpriseSource, OverrideEnterpriseSource

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/config/default_ce.go> a code:Module ;
    code:name "agent/config/default_ce.go" ;
    code:description "Default Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :DefaultEnterpriseSource, :OverrideEnterpriseSource ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package config

// DefaultEnterpriseSource returns the consul agent configuration for enterprise mode.
// These can be overridden by the user and therefore this source should be merged in the
// head and processed before user configuration.
func DefaultEnterpriseSource() Source {
	return LiteralSource{Name: "enterprise-defaults"}
}

// OverrideEnterpriseSource returns the consul agent configuration for the enterprise mode.
// This should be merged in the tail after the DefaultConsulSource.
func OverrideEnterpriseSource() Source {
	return LiteralSource{Name: "enterprise-overrides"}
}
