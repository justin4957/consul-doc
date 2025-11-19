// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/agent_ce.go
Agent Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/config](../agent/config)
- [agent/consul](../agent/consul)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)
- [api](../api)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/agent_ce.go> a code:Module ;
    code:name "agent/agent_ce.go" ;
    code:description "Agent Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/config" ;
        code:path "../agent/config"
    ], [
        code:name "agent/consul" ;
        code:path "../agent/consul"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"context"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/config"
	"github.com/hashicorp/consul/agent/consul"
	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/api"
)

// enterpriseAgent embeds fields that we only access in consul-enterprise builds
type enterpriseAgent struct{}

// fillAgentServiceEnterpriseMeta is a noop stub for the func defined agent_ent.go
func fillAgentServiceEnterpriseMeta(_ *api.AgentService, _ *acl.EnterpriseMeta) {}

// fillHealthCheckEnterpriseMeta is a noop stub for the func defined agent_ent.go
func fillHealthCheckEnterpriseMeta(_ *api.HealthCheck, _ *acl.EnterpriseMeta) {}

// initEnterprise is a noop stub for the func defined agent_ent.go
func (a *Agent) initEnterprise(consulCfg *consul.Config) error {
	return nil
}

// reloadEnterprise is a noop stub for the func defined agent_ent.go
func (a *Agent) reloadEnterprise(conf *config.RuntimeConfig) error {
	return nil
}

// enterpriseConsulConfig is a noop stub for the func defined in agent_ent.go
func enterpriseConsulConfig(_ *consul.Config, _ *config.RuntimeConfig) {
}

// validateFIPSConfig is a noop stub for the func defined in agent_ent.go
func validateFIPSConfig(_ *config.RuntimeConfig) error {
	return nil
}

// WriteEvent is a noop stub for the func defined agent_ent.go
func (a *Agent) WriteEvent(eventType string, payload interface{}) {
}

// startLicenseManager is used to start the license management process
func (a *Agent) startLicenseManager(_ context.Context) error {
	return nil
}

// stopLicenseManager is used to stop the license management go routines
func (a *Agent) stopLicenseManager() {}

// enterpriseStats outputs all the Agent stats specific to Consul Enterprise
func (a *Agent) enterpriseStats() map[string]map[string]string {
	return nil
}

func (a *Agent) AgentEnterpriseMeta() *acl.EnterpriseMeta {
	return structs.NodeEnterpriseMetaInDefaultPartition()
}

func (a *Agent) registerEntCache() {}

func (*Agent) fillEnterpriseProxyDataSources(*proxycfg.DataSources) {}

func (a *Agent) writeAuditRPCEvent(_ string, _ string) interface{} {
	return nil
}
