// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-sources/local/config_source.go
Config Source module for agent layer

## Linked Modules
- [agent/grpc-external/limiter](../agent/grpc-external/limiter)
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)

## Tags
agent, configuration, networking, service-mesh

## Exports
ConfigSource, NewConfigSource

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-sources/local/config_source.go> a code:Module ;
    code:name "agent/proxycfg-sources/local/config_source.go" ;
    code:description "Config Source module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/grpc-external/limiter" ;
        code:path "../agent/grpc-external/limiter"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :ConfigSource, :NewConfigSource ;
    code:tags "agent", "configuration", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package local

import (
	"context"

	"github.com/hashicorp/consul/agent/grpc-external/limiter"
	"github.com/hashicorp/consul/agent/proxycfg"
	structs "github.com/hashicorp/consul/agent/structs"
)

// ConfigSource wraps a proxycfg.Manager to create watches on services
// local to the agent (pre-registered by Sync).
type ConfigSource struct {
	manager ConfigManager
}

// NewConfigSource builds a ConfigSource with the given proxycfg.Manager.
func NewConfigSource(cfgMgr ConfigManager) *ConfigSource {
	return &ConfigSource{cfgMgr}
}

func (m *ConfigSource) Watch(serviceID structs.ServiceID, nodeName string, _ string) (
	<-chan *proxycfg.ConfigSnapshot,
	limiter.SessionTerminatedChan,
	proxycfg.SrcTerminatedChan,
	context.CancelFunc,
	error,
) {
	watchCh, cancelWatch := m.manager.Watch(proxycfg.ProxyID{
		ServiceID: serviceID,
		NodeName:  nodeName,

		// Note: we *intentionally* don't set Token here. All watches on local
		// services use the same ACL token, regardless of whatever token is
		// presented in the xDS stream (the token presented to the xDS server
		// is checked before the watch is created).
		Token: "",
	})
	return watchCh, nil, nil, cancelWatch, nil
}
