// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: lib/maps/maps.go
Maps module for internal layer

## Tags
internal

## Exports
SliceOfKeys, SliceOfValues

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<lib/maps/maps.go> a code:Module ;
    code:name "lib/maps/maps.go" ;
    code:description "Maps module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :SliceOfKeys, :SliceOfValues ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package maps

func SliceOfKeys[K comparable, V any](m map[K]V) []K {
	if len(m) == 0 {
		return nil
	}
	res := make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func SliceOfValues[K comparable, V any](m map[K]V) []V {
	if len(m) == 0 {
		return nil
	}
	res := make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
