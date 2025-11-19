// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/metrics.go
Metrics module for agent layer

## Linked Modules
- [agent/consul](../agent/consul)
- [tlsutil](../tlsutil)

## Tags
agent

## Exports
CertExpirationGauges

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/metrics.go> a code:Module ;
    code:name "agent/metrics.go" ;
    code:description "Metrics module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul" ;
        code:path "../agent/consul"
    ], [
        code:name "tlsutil" ;
        code:path "../tlsutil"
    ] ;
    code:exports :CertExpirationGauges ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"crypto/x509"
	"fmt"
	"time"

	"github.com/armon/go-metrics/prometheus"
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/agent/consul"
	"github.com/hashicorp/consul/tlsutil"
)

var CertExpirationGauges = []prometheus.GaugeDefinition{
	{
		Name: metricsKeyAgentTLSCertExpiry,
		Help: "Seconds until the agent tls certificate expires. Updated every hour",
	},
}

var metricsKeyAgentTLSCertExpiry = []string{"agent", "tls", "cert", "expiry"}

// tlsCertExpirationMonitor returns a CertExpirationMonitor which will
// monitor the expiration of the certificate used for agent TLS.
func tlsCertExpirationMonitor(c *tlsutil.Configurator, logger hclog.Logger) consul.CertExpirationMonitor {
	return consul.CertExpirationMonitor{
		Key:    metricsKeyAgentTLSCertExpiry,
		Logger: logger,
		Query: func() (time.Duration, time.Duration, error) {
			raw := c.Cert()
			if raw == nil {
				return 0, 0, fmt.Errorf("tls not enabled")
			}

			cert, err := x509.ParseCertificate(raw.Certificate[0])
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse agent tls cert: %w", err)
			}

			lifetime := time.Since(cert.NotBefore) + time.Until(cert.NotAfter)
			return lifetime, time.Until(cert.NotAfter), nil
		},
	}
}
