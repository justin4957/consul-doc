// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/flags/flag_slice_value.go
Flag Slice Value module for cli layer

## Tags
cli, user-interface

## Exports
AppendSliceValue

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/flags/flag_slice_value.go> a code:Module ;
    code:name "command/flags/flag_slice_value.go" ;
    code:description "Flag Slice Value module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:exports :AppendSliceValue ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package flags

import "strings"

// AppendSliceValue implements the flag.Value interface and allows multiple
// calls to the same variable to append a list.
type AppendSliceValue []string

func (s *AppendSliceValue) String() string {
	return strings.Join(*s, ",")
}

func (s *AppendSliceValue) Set(value string) error {
	if *s == nil {
		*s = make([]string, 0, 1)
	}

	*s = append(*s, value)
	return nil
}
