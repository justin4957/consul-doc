// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build !linux && !darwin

/*
# Module: sdk/freeport/ephemeral_fallback.go
Ephemeral Fallback module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/freeport/ephemeral_fallback.go> a code:Module ;
    code:name "sdk/freeport/ephemeral_fallback.go" ;
    code:description "Ephemeral Fallback module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package freeport

func getEphemeralPortRange() (int, int, error) {
	return 0, 0, nil
}
