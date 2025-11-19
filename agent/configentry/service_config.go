// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/configentry/service_config.go
Service Config module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, configuration

## Exports
ResolvedServiceConfigSet

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/configentry/service_config.go> a code:Module ;
    code:name "agent/configentry/service_config.go" ;
    code:description "Service Config module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :ResolvedServiceConfigSet ;
    code:tags "agent", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package configentry

import (
	"github.com/hashicorp/consul/agent/structs"
)

// ResolvedServiceConfigSet is a wrapped set of raw cross-referenced config
// entries necessary for the ConfigEntry.ResolveServiceConfig RPC process.
//
// None of these are defaulted.
type ResolvedServiceConfigSet struct {
	ServiceDefaults map[structs.ServiceID]*structs.ServiceConfigEntry
	ProxyDefaults   map[string]*structs.ProxyConfigEntry
}

func (r *ResolvedServiceConfigSet) IsEmpty() bool {
	return len(r.ServiceDefaults) == 0 && len(r.ProxyDefaults) == 0
}

func (r *ResolvedServiceConfigSet) GetServiceDefaults(sid structs.ServiceID) *structs.ServiceConfigEntry {
	if r.ServiceDefaults == nil {
		return nil
	}
	return r.ServiceDefaults[sid]
}

func (r *ResolvedServiceConfigSet) GetProxyDefaults(partition string) *structs.ProxyConfigEntry {
	if r.ProxyDefaults == nil {
		return nil
	}
	return r.ProxyDefaults[partition]
}

func (r *ResolvedServiceConfigSet) AddServiceDefaults(entry *structs.ServiceConfigEntry) {
	if entry == nil {
		return
	}

	if r.ServiceDefaults == nil {
		r.ServiceDefaults = make(map[structs.ServiceID]*structs.ServiceConfigEntry)
	}

	sid := structs.NewServiceID(entry.Name, &entry.EnterpriseMeta)
	r.ServiceDefaults[sid] = entry
}

func (r *ResolvedServiceConfigSet) AddProxyDefaults(entry *structs.ProxyConfigEntry) {
	if entry == nil {
		return
	}

	if r.ProxyDefaults == nil {
		r.ProxyDefaults = make(map[string]*structs.ProxyConfigEntry)
	}

	r.ProxyDefaults[entry.PartitionOrDefault()] = entry
}
