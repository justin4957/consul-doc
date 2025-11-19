// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/authz.go
Authz module for internal layer

## Tags
authentication, internal, security

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/authz.go> a code:Module ;
    code:name "internal/resource/authz.go" ;
    code:description "Authz module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "authentication", "internal", "security" .
<!-- End LinkedDoc RDF -->
*/
package resource

func peerNameV2ToV1(peer string) string {
	// The name of the local/default peer is different between v1 and v2.
	if peer == DefaultPeerName {
		return ""
	}
	return peer
}

func peerNameV1ToV2(peer string) string {
	// The name of the local/default peer is different between v1 and v2.
	if peer == "" {
		return DefaultPeerName
	}
	return peer
}
