// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/protohcl/naming.go
Naming module for internal layer

## Tags
internal

## Exports
FieldNamer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/protohcl/naming.go> a code:Module ;
    code:name "internal/protohcl/naming.go" ;
    code:description "Naming module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:exports :FieldNamer ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package protohcl

import "google.golang.org/protobuf/reflect/protoreflect"

type FieldNamer interface {
	NameField(protoreflect.FieldDescriptor) string
	GetField(protoreflect.FieldDescriptors, string) protoreflect.FieldDescriptor
}

type textFieldNamer struct{}

func (textFieldNamer) NameField(fd protoreflect.FieldDescriptor) string {
	return fd.TextName()
}

func (textFieldNamer) GetField(fds protoreflect.FieldDescriptors, name string) protoreflect.FieldDescriptor {
	return fds.ByTextName(name)
}
