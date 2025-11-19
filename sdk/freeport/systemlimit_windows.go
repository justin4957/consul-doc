// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build windows

/*
# Module: sdk/freeport/systemlimit_windows.go
Systemlimit Windows module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<sdk/freeport/systemlimit_windows.go> a code:Module ;
    code:name "sdk/freeport/systemlimit_windows.go" ;
    code:description "Systemlimit Windows module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package freeport

func systemLimit() (int, error) {
	return 0, nil
}
