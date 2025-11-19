// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/flags/merge.go
Merge module for cli layer

## Tags
cli, user-interface

## Exports
Merge

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/flags/merge.go> a code:Module ;
    code:name "command/flags/merge.go" ;
    code:description "Merge module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:exports :Merge ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package flags

import "flag"

func Merge(dst, src *flag.FlagSet) {
	if dst == nil {
		panic("dst cannot be nil")
	}
	if src == nil {
		return
	}
	src.VisitAll(func(f *flag.Flag) {
		dst.Var(f.Value, f.Name, f.Usage)
	})
}
