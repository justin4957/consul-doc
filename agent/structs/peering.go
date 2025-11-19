// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/peering.go
Cluster peering structures for multi-cluster service mesh federation.

## Linked Modules
- [agent/consul/state](../consul/state/peering.go) - Peering state store operations
- [agent/grpc-external/services/peerstream](../../agent/grpc-external/services/peerstream/server.go) - Peer streaming gRPC service
- [api/peering](../../api/peering.go) - API peering types
- [agent/structs/discovery_chain](./discovery_chain.go) - Discovery chain for peered services

## Tags
structures, data-model, api, peering, multi-cluster, federation

## Exports
PeeringToken, PeeringTokenRemote, ExportedServiceList, ExportedDiscoveryChainInfo

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/peering.go> a code:Module ;
    code:language "go" ;
    code:name "agent/structs/peering.go" ;
    code:description "Cluster peering structures for multi-cluster service mesh federation" ;
    code:layer "api" ;
    code:linksTo [
        code:name "agent/consul/state" ;
        code:path "../consul/state/peering.go" ;
        code:relationship "Peering state store operations"
    ], [
        code:name "agent/grpc-external/services/peerstream" ;
        code:path "../../agent/grpc-external/services/peerstream/server.go" ;
        code:relationship "Peer streaming gRPC service"
    ], [
        code:name "api/peering" ;
        code:path "../../api/peering.go" ;
        code:relationship "API peering types"
    ], [
        code:name "agent/structs/discovery_chain" ;
        code:path "./discovery_chain.go" ;
        code:relationship "Discovery chain for peered services"
    ] ;
    code:exports :PeeringToken, :PeeringTokenRemote, :ExportedServiceList, :ExportedDiscoveryChainInfo ;
    code:tags "structures", "data-model", "api", "peering", "multi-cluster", "federation" .
<!-- End LinkedDoc RDF -->
*/
package structs

// PeeringToken identifies a peer in order for a connection to be established.
type PeeringToken struct {
	CA                    []string
	ManualServerAddresses []string
	ServerAddresses       []string
	ServerName            string
	PeerID                string
	EstablishmentSecret   string
	Remote                PeeringTokenRemote
}

type PeeringTokenRemote struct {
	Partition  string
	Datacenter string
	Locality   *Locality
}

type IndexedExportedServiceList struct {
	Services map[string]ServiceList
	QueryMeta
}

// NOTE: this is not serialized via msgpack so it can be changed without concern.
type ExportedServiceList struct {
	// Services is a list of exported services that apply to both standard
	// service discovery and service mesh.
	Services []ServiceName

	// DiscoChains is a map of service names to their exported discovery chains
	// for service mesh purposes as defined in the exported-services
	// configuration entry.
	DiscoChains map[ServiceName]ExportedDiscoveryChainInfo
}

// NOTE: this is not serialized via msgpack so it can be changed without concern.
type ExportedDiscoveryChainInfo struct {
	// Protocol is the overall protocol associated with this discovery chain.
	Protocol string

	// TCPTargets is the list of discovery chain targets that are reachable by
	// this discovery chain.
	//
	// NOTE: this is only populated if Protocol=tcp.
	TCPTargets []*DiscoveryTarget
}

func (i ExportedDiscoveryChainInfo) Equal(o ExportedDiscoveryChainInfo) bool {
	switch {
	case i.Protocol != o.Protocol:
		return false
	case len(i.TCPTargets) != len(o.TCPTargets):
		return false
	}

	for j := 0; j < len(i.TCPTargets); j++ {
		if i.TCPTargets[j].ID != o.TCPTargets[j].ID {
			return false
		}
	}

	return true
}

// ListAllDiscoveryChains returns all discovery chains (union of Services and DiscoChains).
func (list *ExportedServiceList) ListAllDiscoveryChains() map[ServiceName]ExportedDiscoveryChainInfo {
	chainsByName := make(map[ServiceName]ExportedDiscoveryChainInfo)
	if list == nil {
		return chainsByName
	}

	for _, svc := range list.Services {
		chainsByName[svc] = list.DiscoChains[svc]
	}
	for chainName, info := range list.DiscoChains {
		chainsByName[chainName] = info
	}
	return chainsByName
}
