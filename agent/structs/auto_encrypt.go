// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/auto_encrypt.go
Auto Encrypt module for agent layer

## Tags
agent, data-model, types

## Exports
SignedResponse

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/auto_encrypt.go> a code:Module ;
    code:name "agent/structs/auto_encrypt.go" ;
    code:description "Auto Encrypt module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :SignedResponse ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

type SignedResponse struct {
	IssuedCert           IssuedCert     `json:",omitempty"`
	ConnectCARoots       IndexedCARoots `json:",omitempty"`
	ManualCARoots        []string       `json:",omitempty"`
	GossipKey            string         `json:",omitempty"`
	VerifyServerHostname bool           `json:",omitempty"`
}
