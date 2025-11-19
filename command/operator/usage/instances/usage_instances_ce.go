// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: command/operator/usage/instances/usage_instances_ce.go
Usage Instances Ce module for cli layer

## Linked Modules
- [acl](../acl)
- [api](../api)

## Tags
cli, user-interface

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/operator/usage/instances/usage_instances_ce.go> a code:Module ;
    code:name "command/operator/usage/instances/usage_instances_ce.go" ;
    code:description "Usage Instances Ce module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package instances

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/api"
)

const showPartitionNamespace = false

func getBillableInstanceCounts(usage api.ServiceUsage, datacenter string) []serviceCount {
	return []serviceCount{
		{
			datacenter:    datacenter,
			partition:     acl.DefaultPartitionName,
			namespace:     acl.DefaultNamespaceName,
			instanceCount: usage.BillableServiceInstances,
			services:      usage.Services,
		},
	}
}

func getConnectInstanceCounts(usage api.ServiceUsage, datacenter string) []serviceCount {
	var counts []serviceCount

	for serviceType, instanceCount := range usage.ConnectServiceInstances {
		counts = append(counts, serviceCount{
			datacenter:    datacenter,
			partition:     acl.DefaultPartitionName,
			namespace:     acl.DefaultNamespaceName,
			serviceType:   serviceType,
			instanceCount: instanceCount,
		})
	}

	return counts
}
