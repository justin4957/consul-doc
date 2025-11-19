// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/resources.go
Resources module for internal layer

## Linked Modules
- [testing/deployer/topology](../testing/deployer/topology)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/resources.go> a code:Module ;
    code:name "testing/deployer/sprawl/resources.go" ;
    code:description "Resources module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "testing/deployer/topology" ;
        code:path "../testing/deployer/topology"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package sprawl

import (
	"github.com/hashicorp/consul/testing/deployer/topology"
)

func (s *Sprawl) populateInitialResources(cluster *topology.Cluster) error {
	if len(cluster.InitialResources) == 0 {
		return nil
	}

	for _, res := range cluster.InitialResources {
		if _, err := s.writeResource(cluster, res); err != nil {
			return err
		}
	}

	return nil
}
