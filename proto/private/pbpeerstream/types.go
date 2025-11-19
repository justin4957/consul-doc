// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbpeerstream/types.go
Types module for internal layer

## Tags
data-model, internal, types

## Exports
KnownTypeURL

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbpeerstream/types.go> a code:Module ;
    code:name "proto/private/pbpeerstream/types.go" ;
    code:description "Types module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :KnownTypeURL ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package pbpeerstream

const (
	apiTypePrefix = "type.googleapis.com/"

	TypeURLExportedService        = apiTypePrefix + "hashicorp.consul.internal.peerstream.ExportedService"
	TypeURLExportedServiceList    = apiTypePrefix + "hashicorp.consul.internal.peerstream.ExportedServiceList"
	TypeURLPeeringTrustBundle     = apiTypePrefix + "hashicorp.consul.internal.peering.PeeringTrustBundle"
	TypeURLPeeringServerAddresses = apiTypePrefix + "hashicorp.consul.internal.peering.PeeringServerAddresses"
)

func KnownTypeURL(s string) bool {
	switch s {
	case TypeURLExportedService, TypeURLExportedServiceList, TypeURLPeeringTrustBundle, TypeURLPeeringServerAddresses:
		return true
	}
	return false
}
