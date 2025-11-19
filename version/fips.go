// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !fips

/*
# Module: version/fips.go
Fips module for internal layer

## Tags
internal

## Exports
GetFIPSInfo, IsFIPS

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<version/fips.go> a code:Module ;
    code:name "version/fips.go" ;
    code:description "Fips module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :GetFIPSInfo, :IsFIPS ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package version

func IsFIPS() bool {
	return false
}

func GetFIPSInfo() string {
	return ""
}
