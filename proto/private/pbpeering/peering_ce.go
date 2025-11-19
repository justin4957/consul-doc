// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: proto/private/pbpeering/peering_ce.go
Peering Ce module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbpeering/peering_ce.go> a code:Module ;
    code:name "proto/private/pbpeering/peering_ce.go" ;
    code:description "Peering Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbpeering

func (r *GenerateTokenRequest) PartitionOrDefault() string {
	return ""
}

func (p *Peering) PartitionOrDefault() string {
	return ""
}

func (ptb *PeeringTrustBundle) PartitionOrDefault() string {
	return ""
}
