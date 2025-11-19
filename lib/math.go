// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/math.go
Math module for internal layer

## Tags
internal

## Exports
AbsInt, MaxInt, MaxUint64, MinInt

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/math.go> a code:Module ;
    code:name "lib/math.go" ;
    code:description "Math module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :AbsInt, :MaxInt, :MaxUint64, :MinInt ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package lib

func AbsInt(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}
