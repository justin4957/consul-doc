// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/connect/testing_spiffe.go
Testing Spiffe module for agent layer

## Linked Modules
- [sdk/testutil](../sdk/testutil)

## Tags
agent, mtls, service-mesh

## Exports
TestSpiffeIDService, TestSpiffeIDServiceWithHost, TestSpiffeIDServiceWithHostDC

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/connect/testing_spiffe.go> a code:Module ;
    code:name "agent/connect/testing_spiffe.go" ;
    code:description "Testing Spiffe module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "sdk/testutil" ;
        code:path "../sdk/testutil"
    ] ;
    code:exports :TestSpiffeIDService, :TestSpiffeIDServiceWithHost, :TestSpiffeIDServiceWithHostDC ;
    code:tags "agent", "mtls", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connect

import "github.com/hashicorp/consul/sdk/testutil"

// TestSpiffeIDService returns a SPIFFE ID representing a service.
func TestSpiffeIDService(t testutil.TestingTB, service string) *SpiffeIDService {
	return TestSpiffeIDServiceWithHost(t, service, TestClusterID+".consul")
}

// TestSpiffeIDServiceWithHost returns a SPIFFE ID representing a service with
// the specified trust domain.
func TestSpiffeIDServiceWithHost(t testutil.TestingTB, service, host string) *SpiffeIDService {
	return TestSpiffeIDServiceWithHostDC(t, service, host, "dc1")
}

// TestSpiffeIDServiceWithHostDC returns a SPIFFE ID representing a service with
// the specified trust domain for the given datacenter.
func TestSpiffeIDServiceWithHostDC(t testutil.TestingTB, service, host, datacenter string) *SpiffeIDService {
	return &SpiffeIDService{
		Host:       host,
		Namespace:  "default",
		Datacenter: datacenter,
		Service:    service,
	}
}
