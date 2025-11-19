// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resourcehcl/naming.go
Naming module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resourcehcl/naming.go> a code:Module ;
    code:name "internal/resourcehcl/naming.go" ;
    code:description "Naming module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcehcl

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// fieldNamer implements protohcl.FieldNamer to name fields using PascalCase
// with support for acroynms (e.g. ID, TCP).
type fieldNamer struct{ acroynms []string }

func (n fieldNamer) NameField(fd protoreflect.FieldDescriptor) string {
	camel := fd.JSONName()
	upper := strings.ToUpper(camel)

	for _, a := range n.acroynms {
		if upper == a {
			return a
		}
	}

	return strings.ToUpper(camel[:1]) + camel[1:]
}

func (n fieldNamer) GetField(fds protoreflect.FieldDescriptors, name string) protoreflect.FieldDescriptor {
	for _, a := range n.acroynms {
		if name == a {
			return fds.ByJSONName(strings.ToLower(a))
		}
	}

	if len(name) <= 1 {
		return fds.ByJSONName(name)
	}
	camel := strings.ToLower(name[:1]) + name[1:]
	return fds.ByJSONName(camel)
}
