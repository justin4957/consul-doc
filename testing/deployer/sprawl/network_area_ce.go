// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: testing/deployer/sprawl/network_area_ce.go
Network Area Ce module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/network_area_ce.go> a code:Module ;
    code:name "testing/deployer/sprawl/network_area_ce.go" ;
    code:description "Network Area Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sprawl

func (s *Sprawl) initNetworkAreas() error {
	return nil
}

func (s *Sprawl) waitForNetworkAreaEstablishment() error {
	return nil
}
