// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/consul/enterprise_client_ce.go
Enterprise Client Ce module for agent layer

## Tags
agent

## Exports
EnterpriseClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/enterprise_client_ce.go> a code:Module ;
    code:name "agent/consul/enterprise_client_ce.go" ;
    code:description "Enterprise Client Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :EnterpriseClient ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"github.com/hashicorp/serf/serf"
)

type EnterpriseClient struct{}

func (c *Client) initEnterprise(_ Deps) error {
	return nil
}

func enterpriseModifyClientSerfConfigLAN(_ *Config, _ *serf.Config) {
	// nothing
}

func (c *Client) startEnterprise() error {
	return nil
}

func (c *Client) handleEnterpriseUserEvents(event serf.UserEvent) bool {
	return false
}

func (c *Client) enterpriseStats() map[string]map[string]string {
	return nil
}
