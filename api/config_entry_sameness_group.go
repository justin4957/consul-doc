// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

/*
# Module: api/config_entry_sameness_group.go
Config Entry Sameness Group module for api layer

## Tags
api, client, configuration

## Exports
SamenessGroupConfigEntry, SamenessGroupMember

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<api/config_entry_sameness_group.go> a code:Module ;
    code:name "api/config_entry_sameness_group.go" ;
    code:description "Config Entry Sameness Group module for api layer" ;
    code:language "go" ;
    code:layer "api" ;
    code:exports :SamenessGroupConfigEntry, :SamenessGroupMember ;
    code:tags "api", "client", "configuration" .
<!-- End LinkedDoc RDF -->
*/
package api

type SamenessGroupConfigEntry struct {
	Kind               string
	Name               string
	Partition          string `json:",omitempty"`
	DefaultForFailover bool   `json:",omitempty" alias:"default_for_failover"`
	IncludeLocal       bool   `json:",omitempty" alias:"include_local"`
	Members            []SamenessGroupMember
	Meta               map[string]string `json:",omitempty"`
	CreateIndex        uint64
	ModifyIndex        uint64
}

type SamenessGroupMember struct {
	Partition string `json:",omitempty"`
	Peer      string `json:",omitempty"`
}

func (s *SamenessGroupConfigEntry) GetKind() string            { return s.Kind }
func (s *SamenessGroupConfigEntry) GetName() string            { return s.Name }
func (s *SamenessGroupConfigEntry) GetPartition() string       { return s.Partition }
func (s *SamenessGroupConfigEntry) GetNamespace() string       { return "" }
func (s *SamenessGroupConfigEntry) GetCreateIndex() uint64     { return s.CreateIndex }
func (s *SamenessGroupConfigEntry) GetModifyIndex() uint64     { return s.ModifyIndex }
func (s *SamenessGroupConfigEntry) GetMeta() map[string]string { return s.Meta }
