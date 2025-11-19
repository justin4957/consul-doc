// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/hcp/scada/capabilities.go
Capabilities module for agent layer

## Tags
agent

## Exports
CAPCoreAPI

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/hcp/scada/capabilities.go> a code:Module ;
    code:name "agent/hcp/scada/capabilities.go" ;
    code:description "Capabilities module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :CAPCoreAPI ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package scada

import "github.com/hashicorp/hcp-scada-provider/capability"

// CAPCoreAPI is the capability used to securely expose the Consul HTTP API to HCP
var CAPCoreAPI = capability.NewAddr("core_api")
