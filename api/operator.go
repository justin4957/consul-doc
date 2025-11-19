// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/operator.go
Operator module for api layer

## Tags
api, client

## Exports
Operator

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/operator.go> a code:Module ;
    code:name "api/operator.go" ;
    code:description "Operator module for api layer" ;
    code:language "go" ;
    code:layer "api" ;
    code:exports :Operator ;
    code:tags "api", "client" .
<!-- End LinkedDoc RDF -->
*/
package api

// Operator can be used to perform low-level operator tasks for Consul.
type Operator struct {
	c *Client
}

// Operator returns a handle to the operator endpoints.
func (c *Client) Operator() *Operator {
	return &Operator{c}
}
