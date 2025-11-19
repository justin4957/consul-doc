// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/system_metadata.go
System Metadata module for agent layer

## Tags
agent, data-model, types

## Exports
SystemMetadataEntry, SystemMetadataOp, SystemMetadataRequest

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/system_metadata.go> a code:Module ;
    code:name "agent/structs/system_metadata.go" ;
    code:description "System Metadata module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :SystemMetadataEntry, :SystemMetadataOp, :SystemMetadataRequest ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

// SystemMetadataOp is the operation for a request related to system metadata.
type SystemMetadataOp string

const (
	SystemMetadataUpsert SystemMetadataOp = "upsert"
	SystemMetadataDelete SystemMetadataOp = "delete"
)

// SystemMetadataRequest is used to upsert and delete system metadata.
type SystemMetadataRequest struct {
	// Datacenter is the target for this request.
	Datacenter string

	// Op is the type of operation being requested.
	Op SystemMetadataOp

	// Entry is the key to modify.
	Entry *SystemMetadataEntry

	// WriteRequest is a common struct containing ACL tokens and other
	// write-related common elements for requests.
	WriteRequest
}

const (
	SystemMetadataIntentionFormatKey           = "intention-format"
	SystemMetadataIntentionFormatConfigValue   = "config-entry"
	SystemMetadataIntentionFormatLegacyValue   = "legacy"
	SystemMetadataVirtualIPsEnabled            = "virtual-ips"
	SystemMetadataTermGatewayVirtualIPsEnabled = "virtual-ips-term-gateway"
)

type SystemMetadataEntry struct {
	Key   string
	Value string `json:",omitempty"`
	RaftIndex
}

// RequestDatacenter returns the datacenter for a given request.
func (c *SystemMetadataRequest) RequestDatacenter() string {
	return c.Datacenter
}
