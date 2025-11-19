// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbservice/ids.go
Ids module for internal layer

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbservice/ids.go> a code:Module ;
    code:name "proto/private/pbservice/ids.go" ;
    code:description "Ids module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbservice

import (
	"strings"
)

// UniqueID returns a unique identifier for this CheckServiceNode, which includes
// the node name, service namespace, and service ID.
//
// The returned ID uses slashes to separate the identifiers, however the node name
// may also contain a slash, so it is not possible to parse this identifier to
// retrieve its constituent parts.
//
// This function is similar to structs.UniqueID, however at this time no guarantees
// are made that it will remain the same.
func (m *CheckServiceNode) UniqueID() string {
	if m == nil {
		return ""
	}
	builder := new(strings.Builder)

	switch {
	case m.Node != nil:
		builder.WriteString(m.Node.Partition + "/")
		builder.WriteString(m.Node.PeerName + "/")
	case m.Service != nil:
		partition := ""
		if m.Service.EnterpriseMeta != nil {
			partition = m.Service.EnterpriseMeta.Partition
		}
		builder.WriteString(partition + "/")
		builder.WriteString(m.Service.PeerName + "/")
	}

	if m.Node != nil {
		builder.WriteString(m.Node.Node + "/")
	}
	if m.Service != nil {
		namespace := ""
		if m.Service.EnterpriseMeta != nil {
			namespace = m.Service.EnterpriseMeta.Namespace
		}
		builder.WriteString(namespace + "/")
		builder.WriteString(m.Service.ID)
	}
	return builder.String()
}
