// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/auto-config/config_ce.go
Config Ce module for agent layer

## Tags
agent, configuration

## Exports
EnterpriseConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/auto-config/config_ce.go> a code:Module ;
    code:name "agent/auto-config/config_ce.go" ;
    code:description "Config Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EnterpriseConfig ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package autoconf

// EnterpriseConfig stub - only populated in Consul Enterprise
type EnterpriseConfig struct{}

// finalize is a noop for CE
func (*EnterpriseConfig) validateAndFinalize() error {
	return nil
}
