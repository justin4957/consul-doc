// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg-glue/leafcerts.go
Leafcerts module for agent layer

## Linked Modules
- [agent/leafcert](../agent/leafcert)
- [agent/proxycfg](../agent/proxycfg)

## Tags
agent, networking, service-mesh

## Exports
LocalLeafCerts

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg-glue/leafcerts.go> a code:Module ;
    code:name "agent/proxycfg-glue/leafcerts.go" ;
    code:description "Leafcerts module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/leafcert" ;
        code:path "../agent/leafcert"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../agent/proxycfg"
    ] ;
    code:exports :LocalLeafCerts ;
    code:tags "agent", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfgglue

import (
	"context"

	"github.com/hashicorp/consul/agent/leafcert"
	"github.com/hashicorp/consul/agent/proxycfg"
)

// LocalLeafCerts satisfies the proxycfg.LeafCertificate interface by sourcing data from
// the given leafcert.Manager.
func LocalLeafCerts(m *leafcert.Manager) proxycfg.LeafCertificate {
	return &localLeafCerts{m}
}

type localLeafCerts struct {
	leafCertManager *leafcert.Manager
}

func (c *localLeafCerts) Notify(ctx context.Context, req *leafcert.ConnectCALeafRequest, correlationID string, ch chan<- proxycfg.UpdateEvent) error {
	return c.leafCertManager.NotifyCallback(ctx, req, correlationID, dispatchCacheUpdate(ch))
}
