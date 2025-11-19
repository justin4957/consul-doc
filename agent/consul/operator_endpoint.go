// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/operator_endpoint.go
Operator Endpoint module for agent layer

## Tags
agent

## Exports
Operator

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/operator_endpoint.go> a code:Module ;
    code:name "agent/consul/operator_endpoint.go" ;
    code:description "Operator Endpoint module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :Operator ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import "github.com/hashicorp/go-hclog"

// Operator endpoint is used to perform low-level operator tasks for Consul.
type Operator struct {
	srv    *Server
	logger hclog.Logger
}
