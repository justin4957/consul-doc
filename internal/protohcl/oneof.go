// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/protohcl/oneof.go
Oneof module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/protohcl/oneof.go> a code:Module ;
    code:name "internal/protohcl/oneof.go" ;
    code:description "Oneof module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package protohcl

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

type oneOfTracker struct {
	namer FieldNamer
	set   map[protoreflect.FullName]string
}

func newOneOfTracker(namer FieldNamer) *oneOfTracker {
	return &oneOfTracker{
		namer: namer,
		set:   make(map[protoreflect.FullName]string),
	}
}

func (o *oneOfTracker) markFieldAsSet(desc protoreflect.FieldDescriptor) error {
	oneof := desc.ContainingOneof()
	if oneof == nil {
		return nil
	}

	oneOfName := oneof.FullName()

	if otherFieldName, ok := o.set[oneOfName]; ok {
		oneOfFields := oneof.Fields()
		var builder strings.Builder

		for i := 0; i < oneOfFields.Len(); i++ {
			name := o.namer.NameField(oneOfFields.Get(i))

			if i == oneOfFields.Len()-1 {
				builder.WriteString(fmt.Sprintf("%q", name))
			} else if i == oneOfFields.Len()-2 {
				builder.WriteString(fmt.Sprintf("%q and ", name))
			} else {
				builder.WriteString(fmt.Sprintf("%q, ", name))
			}
		}

		return fmt.Errorf("Cannot set %q because %q was previously set. Only one of %s may be set.", o.namer.NameField(desc), otherFieldName, builder.String())
	}

	o.set[oneOfName] = o.namer.NameField(desc)
	return nil
}
