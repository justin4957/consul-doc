// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/xds/secrets.go
Secrets module for agent layer

## Linked Modules
- [agent/proxycfg](../agent/proxycfg)
- [agent/structs](../agent/structs)

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/secrets.go> a code:Module ;
    code:name "agent/xds/secrets.go" ;
    code:description "Secrets module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	"errors"
	"fmt"

	envoy_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_tls_v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/agent/proxycfg"
	"github.com/hashicorp/consul/agent/structs"
)

// secretsFromSnapshot returns the xDS API representation of the "secrets"
// in the snapshot
func (s *ResourceGenerator) secretsFromSnapshot(cfgSnap *proxycfg.ConfigSnapshot) ([]proto.Message, error) {
	if cfgSnap == nil {
		return nil, errors.New("nil config given")
	}

	switch cfgSnap.Kind {
	case structs.ServiceKindAPIGateway:
		return s.secretsFromSnapshotAPIGateway(cfgSnap), nil // return any attached certs
	case structs.ServiceKindConnectProxy,
		structs.ServiceKindTerminatingGateway,
		structs.ServiceKindMeshGateway,
		structs.ServiceKindIngressGateway:
		return nil, nil
	default:
		return nil, fmt.Errorf("Invalid service kind: %v", cfgSnap.Kind)
	}
}

// secretsFromSnapshotAPIGateway returns the "secrets" for an api-gateway service
func (s *ResourceGenerator) secretsFromSnapshotAPIGateway(cfgSnap *proxycfg.ConfigSnapshot) []proto.Message {
	var resources []proto.Message

	cfgSnap.APIGateway.FileSystemCertificates.ForEachKey(func(ref structs.ResourceReference) bool {
		cert, ok := cfgSnap.APIGateway.FileSystemCertificates.Get(ref)
		if !ok || cert == nil {
			return true
		}
		resources = append(resources, &envoy_tls_v3.Secret{
			Name: ref.Name,
			Type: &envoy_tls_v3.Secret_TlsCertificate{
				TlsCertificate: &envoy_tls_v3.TlsCertificate{
					CertificateChain: &envoy_core_v3.DataSource{
						Specifier: &envoy_core_v3.DataSource_Filename{
							Filename: cert.Certificate,
						}},
					PrivateKey: &envoy_core_v3.DataSource{
						Specifier: &envoy_core_v3.DataSource_Filename{
							Filename: cert.PrivateKey,
						},
					},
				},
			},
		})
		return true
	})

	return resources
}
