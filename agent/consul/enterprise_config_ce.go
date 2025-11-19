// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/enterprise_config_ce.go
Enterprise Config Ce module for agent layer

## Tags
agent, configuration

## Exports
DefaultEnterpriseConfig, EnterpriseConfig

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/enterprise_config_ce.go> a code:Module ;
    code:name "agent/consul/enterprise_config_ce.go" ;
    code:description "Enterprise Config Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :DefaultEnterpriseConfig, :EnterpriseConfig ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package consul

type EnterpriseConfig struct{}

func DefaultEnterpriseConfig() *EnterpriseConfig {
	return nil
}
