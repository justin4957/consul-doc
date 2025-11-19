// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/protoutil/protoutil.go
Protoutil module for internal layer

## Tags
internal

## Exports
Clone, CloneSlice

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/protoutil/protoutil.go> a code:Module ;
    code:name "internal/protoutil/protoutil.go" ;
    code:description "Protoutil module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :Clone, :CloneSlice ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package protoutil

import (
	"google.golang.org/protobuf/proto"
)

func Clone[T proto.Message](v T) T {
	return proto.Clone(v).(T)
}

func CloneSlice[T proto.Message](in []T) []T {
	if in == nil {
		return nil
	}
	out := make([]T, 0, len(in))
	for _, v := range in {
		out = append(out, Clone[T](v))
	}
	return out
}
