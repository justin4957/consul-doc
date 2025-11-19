// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/config_entry_sameness_group_ce.go
Config Entry Sameness Group Ce module for agent layer

## Tags
agent, configuration, data-model, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/config_entry_sameness_group_ce.go> a code:Module ;
    code:name "agent/structs/config_entry_sameness_group_ce.go" ;
    code:description "Config Entry Sameness Group Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "configuration", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import "fmt"

// Validate assures that the sameness-groups are an enterprise only feature
func (s *SamenessGroupConfigEntry) Validate() error {
	return fmt.Errorf("sameness-groups are an enterprise-only feature")
}

// RelatedPeers is an CE placeholder noop
func (s *SamenessGroupConfigEntry) RelatedPeers() []string {
	return nil
}

// AllMembers is an CE placeholder noop
func (s *SamenessGroupConfigEntry) AllMembers() []SamenessGroupMember {
	return nil
}

// ToServiceResolverFailoverTargets is an CE placeholder noop
func (s *SamenessGroupConfigEntry) ToServiceResolverFailoverTargets() []ServiceResolverFailoverTarget {
	return nil
}

// ToQueryFailoverTargets is an CE placeholder noop
func (s *SamenessGroupConfigEntry) ToQueryFailoverTargets(namespace string) []QueryFailoverTarget {
	return nil
}
