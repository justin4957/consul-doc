// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/resourcetest/testing.go
Testing module for internal layer

## Linked Modules
- [sdk/testutil](../sdk/testutil)

## Tags
internal

## Exports
T

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/resourcetest/testing.go> a code:Module ;
    code:name "internal/resource/resourcetest/testing.go" ;
    code:description "Testing module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "sdk/testutil" ;
        code:path "../sdk/testutil"
    ] ;
    code:exports :T ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcetest

import "github.com/hashicorp/consul/sdk/testutil"

// T represents the subset of testing.T methods that will be used
// by the various functionality in this package
type T interface {
	testutil.TestingTB
}
