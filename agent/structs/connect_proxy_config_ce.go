// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/structs/connect_proxy_config_ce.go
Connect Proxy Config Ce module for agent layer

## Linked Modules
- [acl](../acl)

## Tags
agent, configuration, data-model, mtls, networking, service-mesh, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/connect_proxy_config_ce.go> a code:Module ;
    code:name "agent/structs/connect_proxy_config_ce.go" ;
    code:description "Connect Proxy Config Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ] ;
    code:tags "agent", "configuration", "data-model", "mtls", "networking", "service-mesh", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/hashicorp/consul/acl"
)

func (us *Upstream) GetEnterpriseMeta() *acl.EnterpriseMeta {
	return DefaultEnterpriseMetaInDefaultPartition()
}

func (us *Upstream) DestinationID() PeeredServiceName {
	return PeeredServiceName{
		Peer:        us.DestinationPeer,
		ServiceName: NewServiceName(us.DestinationName, DefaultEnterpriseMetaInDefaultPartition()),
	}
}

func (us *Upstream) enterpriseStringPrefix() string {
	return ""
}
