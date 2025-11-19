// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/state/config_entry_exported_services_ce.go
Config Entry Exported Services Ce module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/configentry](../agent/configentry)
- [agent/structs](../agent/structs)
- [proto/private/pbconfigentry](../proto/private/pbconfigentry)

## Tags
agent, configuration, persistence, storage

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/state/config_entry_exported_services_ce.go> a code:Module ;
    code:name "agent/consul/state/config_entry_exported_services_ce.go" ;
    code:description "Config Entry Exported Services Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/configentry" ;
        code:path "../agent/configentry"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pbconfigentry" ;
        code:path "../proto/private/pbconfigentry"
    ] ;
    code:tags "agent", "configuration", "persistence", "storage" .
<!-- End LinkedDoc RDF -->
*/
package state

import (
	"sort"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pbconfigentry"
	"github.com/hashicorp/go-memdb"
)

func getSimplifiedExportedServices(
	tx ReadTxn,
	ws memdb.WatchSet,
	overrides map[configentry.KindName]structs.ConfigEntry,
	entMeta acl.EnterpriseMeta,
) (uint64, *SimplifiedExportedServices, error) {
	idx, exports, err := getExportedServicesConfigEntryTxn(tx, ws, overrides, &entMeta)
	if exports == nil {
		return idx, nil, err
	}
	simple := SimplifiedExportedServices(*exports)
	return idx, &simple, err
}

func (s *Store) GetSimplifiedExportedServices(ws memdb.WatchSet, entMeta acl.EnterpriseMeta) (uint64, *SimplifiedExportedServices, error) {
	tx := s.db.Txn(false)
	defer tx.Abort()
	return getSimplifiedExportedServices(tx, ws, nil, entMeta)
}

func prepareExportedServicesResponse(exportedServices []structs.ExportedService, entMeta *acl.EnterpriseMeta) []*pbconfigentry.ResolvedExportedService {

	resp := make([]*pbconfigentry.ResolvedExportedService, len(exportedServices))

	for idx, exportedService := range exportedServices {
		consumerPeers := []string{}

		for _, consumer := range exportedService.Consumers {
			if consumer.Peer != "" {
				consumerPeers = append(consumerPeers, consumer.Peer)
			}
		}

		sort.Strings(consumerPeers)

		resp[idx] = &pbconfigentry.ResolvedExportedService{
			Service: exportedService.Name,
			Consumers: &pbconfigentry.Consumers{
				Peers: consumerPeers,
			},
		}
	}

	return resp
}
