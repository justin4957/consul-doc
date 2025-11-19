// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/testhelpers/testhelpers.go
Testhelpers module for internal layer

## Tags
internal

## Exports
SkipFlake

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/testhelpers/testhelpers.go> a code:Module ;
    code:name "lib/testhelpers/testhelpers.go" ;
    code:description "Testhelpers module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :SkipFlake ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package testhelpers

import (
	"os"
	"testing"
)

func SkipFlake(t *testing.T) {
	if os.Getenv("RUN_FLAKEY_TESTS") != "true" {
		t.Skip("Skipped because marked as flake.")
	}
}
