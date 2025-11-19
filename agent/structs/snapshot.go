// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/snapshot.go
Snapshot module for agent layer

## Tags
agent, data-model, types

## Exports
SnapshotOp, SnapshotReplyFn, SnapshotRequest, SnapshotResponse

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/snapshot.go> a code:Module ;
    code:name "agent/structs/snapshot.go" ;
    code:description "Snapshot module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :SnapshotOp, :SnapshotReplyFn, :SnapshotRequest, :SnapshotResponse ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

type SnapshotOp int

const (
	SnapshotSave SnapshotOp = iota
	SnapshotRestore
)

// SnapshotReplyFn gets a peek at the reply before the snapshot streams, which
// is useful for setting headers.
type SnapshotReplyFn func(reply *SnapshotResponse) error

// SnapshotRequest is used as a header for a snapshot RPC request. This will
// precede any streaming data that's part of the request and is JSON-encoded on
// the wire.
type SnapshotRequest struct {
	// Datacenter is the target datacenter for this request. The request
	// will be forwarded if necessary.
	Datacenter string

	// Token is the ACL token to use for the operation. If ACLs are enabled
	// then all operations require a management token.
	Token string

	// If set, any follower can service the request. Results may be
	// arbitrarily stale. Only applies to SnapshotSave.
	AllowStale bool

	// Op is the operation code for the RPC.
	Op SnapshotOp
}

// SnapshotResponse is used header for a snapshot RPC response. This will
// precede any streaming data that's part of the request and is JSON-encoded on
// the wire.
type SnapshotResponse struct {
	// Error is the overall error status of the RPC request.
	Error string

	// QueryMeta has freshness information about the server that handled the
	// request. It is only filled in for a SnapshotSave.
	QueryMeta
}
