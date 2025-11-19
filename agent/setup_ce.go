// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/setup_ce.go
Setup Ce module for agent layer

## Linked Modules
- [agent/auto-config](../agent/auto-config)
- [agent/config](../agent/config)
- [agent/consul](../agent/consul)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/setup_ce.go> a code:Module ;
    code:name "agent/setup_ce.go" ;
    code:description "Setup Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/auto-config" ;
        code:path "../agent/auto-config"
    ], [
        code:name "agent/config" ;
        code:path "../agent/config"
    ], [
        code:name "agent/consul" ;
        code:path "../agent/consul"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	autoconf "github.com/hashicorp/consul/agent/auto-config"
	"github.com/hashicorp/consul/agent/config"
	"github.com/hashicorp/consul/agent/consul"
)

// initEnterpriseBaseDeps is responsible for initializing the enterprise dependencies that
// will be utilized throughout the whole Consul Agent.
func initEnterpriseBaseDeps(d BaseDeps, _ *config.RuntimeConfig) (BaseDeps, error) {
	return d, nil
}

// initEnterpriseAutoConfig is responsible for setting up auto-config for enterprise
func initEnterpriseAutoConfig(_ consul.EnterpriseDeps, _ *config.RuntimeConfig) autoconf.EnterpriseConfig {
	return autoconf.EnterpriseConfig{}
}
