// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: test/integration/consul-container/test/gateways/tenancy_ce.go
Tenancy Ce module for internal layer

## Linked Modules
- [api](../api)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<test/integration/consul-container/test/gateways/tenancy_ce.go> a code:Module ;
    code:name "test/integration/consul-container/test/gateways/tenancy_ce.go" ;
    code:description "Tenancy Ce module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package gateways

import (
	"testing"

	"github.com/hashicorp/consul/api"
)

func getOrCreateNamespace(_ *testing.T, _ *api.Client) string {
	return ""
}
