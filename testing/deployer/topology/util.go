// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/topology/util.go
Util module for internal layer

## Tags
internal

## Exports
MergeSlices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/topology/util.go> a code:Module ;
    code:name "testing/deployer/topology/util.go" ;
    code:description "Util module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :MergeSlices ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package topology

func MergeSlices[V any](x, y []V) []V {
	switch {
	case len(x) == 0 && len(y) == 0:
		return nil
	case len(x) == 0:
		return y
	case len(y) == 0:
		return x
	}

	out := make([]V, 0, len(x)+len(y))
	out = append(out, x...)
	out = append(out, y...)
	return out
}
