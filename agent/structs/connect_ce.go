// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/connect_ce.go
Connect Ce module for agent layer

## Tags
agent, data-model, mtls, service-mesh, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/connect_ce.go> a code:Module ;
    code:name "agent/structs/connect_ce.go" ;
    code:description "Connect Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "data-model", "mtls", "service-mesh", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

func (req *ConnectAuthorizeRequest) TargetNamespace() string {
	return IntentionDefaultNamespace
}
