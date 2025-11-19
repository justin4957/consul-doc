// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/system_metadata.go
System Metadata module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/system_metadata.go> a code:Module ;
    code:name "agent/consul/system_metadata.go" ;
    code:description "System Metadata module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/consul/agent/structs"
)

func (s *Server) GetSystemMetadata(key string) (string, error) {
	_, entry, err := s.fsm.State().SystemMetadataGet(nil, key)
	if err != nil {
		return "", err
	}
	if entry == nil {
		return "", nil
	}

	return entry.Value, nil
}

func (s *Server) SetSystemMetadataKey(key, val string) error {
	args := &structs.SystemMetadataRequest{
		Op:    structs.SystemMetadataUpsert,
		Entry: &structs.SystemMetadataEntry{Key: key, Value: val},
	}

	// TODO(rpc-metrics-improv): Double check request name here
	_, err := s.leaderRaftApply("SystemMetadata.Upsert", structs.SystemMetadataRequestType, args)

	return err
}

func (s *Server) ApplyCensusRequest(req *structs.CensusRequest) error {
	// TODO(rpc-metrics-improv): dedicated metrics label for reporting/manual snapshot operations
	_, err := s.leaderRaftApply("Reporting.Census", structs.CensusRequestType, req)
	return err
}

func (s *Server) deleteSystemMetadataKey(key string) error {
	args := &structs.SystemMetadataRequest{
		Op:    structs.SystemMetadataDelete,
		Entry: &structs.SystemMetadataEntry{Key: key},
	}

	// TODO(rpc-metrics-improv): Double check request name here
	_, err := s.leaderRaftApply("SystemMetadata.Delete", structs.SystemMetadataRequestType, args)

	return err
}
