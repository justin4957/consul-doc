// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/server_connect.go
Server Connect module for agent layer

## Linked Modules
- [agent/connect](../agent/connect)
- [agent/consul/state](../agent/consul/state)
- [agent/structs](../agent/structs)
- [lib](../lib)

## Tags
agent, mtls, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/server_connect.go> a code:Module ;
    code:name "agent/consul/server_connect.go" ;
    code:description "Server Connect module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/connect" ;
        code:path "../agent/connect"
    ], [
        code:name "agent/consul/state" ;
        code:path "../agent/consul/state"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "lib" ;
        code:path "../lib"
    ] ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"fmt"

	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/agent/connect"
	"github.com/hashicorp/consul/agent/consul/state"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/lib"
)

func (s *Server) getCARoots(ws memdb.WatchSet, state *state.Store) (*structs.IndexedCARoots, error) {
	index, roots, config, err := state.CARootsAndConfig(ws)
	if err != nil {
		return nil, err
	}

	trustDomain, err := s.getTrustDomain(config)
	if err != nil {
		return nil, err
	}

	indexedRoots := &structs.IndexedCARoots{}

	indexedRoots.TrustDomain = trustDomain

	indexedRoots.Index, indexedRoots.Roots = index, roots
	if indexedRoots.Roots == nil {
		indexedRoots.Roots = make(structs.CARoots, 0)
	}

	// The response should not contain all fields as there are sensitive
	// data such as key material stored within the struct. So here we
	// pull out some of the fields and copy them into
	for i, r := range indexedRoots.Roots {
		var intermediates []string
		if r.IntermediateCerts != nil {
			intermediates = make([]string, len(r.IntermediateCerts))
			copy(intermediates, r.IntermediateCerts)
		}
		// IMPORTANT: r must NEVER be modified, since it is a pointer
		// directly to the structure in the memdb store.

		indexedRoots.Roots[i] = &structs.CARoot{
			ID:                  r.ID,
			Name:                r.Name,
			SerialNumber:        r.SerialNumber,
			SigningKeyID:        r.SigningKeyID,
			ExternalTrustDomain: r.ExternalTrustDomain,
			NotBefore:           r.NotBefore,
			NotAfter:            r.NotAfter,
			RootCert:            lib.EnsureTrailingNewline(r.RootCert),
			IntermediateCerts:   intermediates,
			RaftIndex:           r.RaftIndex,
			Active:              r.Active,
			PrivateKeyType:      r.PrivateKeyType,
			PrivateKeyBits:      r.PrivateKeyBits,
		}

		if r.Active {
			indexedRoots.ActiveRootID = r.ID
		}
	}

	return indexedRoots, nil
}

func (s *Server) getTrustDomain(config *structs.CAConfiguration) (string, error) {
	if config == nil || config.ClusterID == "" {
		return "", fmt.Errorf("CA has not finished initializing")
	}

	// Build TrustDomain based on the ClusterID stored.
	signingID := connect.SpiffeIDSigningForCluster(config.ClusterID)
	if signingID == nil {
		// If CA is bootstrapped at all then this should never happen but be
		// defensive.
		return "", fmt.Errorf("no cluster trust domain setup")
	}

	return signingID.Host(), nil
}
