// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: testing/deployer/sprawl/grpc.go
Grpc module for internal layer

## Linked Modules
- [testing/deployer/sprawl/internal/secrets](../testing/deployer/sprawl/internal/secrets)
- [testing/deployer/topology](../testing/deployer/topology)
- [testing/deployer/util](../testing/deployer/util)

## Tags
api, communication, grpc, internal, networking

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<testing/deployer/sprawl/grpc.go> a code:Module ;
    code:name "testing/deployer/sprawl/grpc.go" ;
    code:description "Grpc module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "testing/deployer/sprawl/internal/secrets" ;
        code:path "../testing/deployer/sprawl/internal/secrets"
    ], [
        code:name "testing/deployer/topology" ;
        code:path "../testing/deployer/topology"
    ], [
        code:name "testing/deployer/util" ;
        code:path "../testing/deployer/util"
    ] ;
    code:tags "api", "communication", "grpc", "internal", "networking" .
<!-- End LinkedDoc RDF -->
*/
package sprawl

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/hashicorp/go-rootcerts"
	"google.golang.org/grpc"

	"github.com/hashicorp/consul/testing/deployer/sprawl/internal/secrets"
	"github.com/hashicorp/consul/testing/deployer/topology"
	"github.com/hashicorp/consul/testing/deployer/util"
)

func (s *Sprawl) dialServerGRPC(cluster *topology.Cluster, node *topology.Node, token string) (*grpc.ClientConn, func(), error) {
	var (
		logger = s.logger.With("cluster", cluster.Name)
	)

	tls := &tls.Config{
		ServerName: fmt.Sprintf("server.%s.consul", cluster.Datacenter),
	}

	rootConfig := &rootcerts.Config{
		CACertificate: []byte(s.secrets.ReadGeneric(cluster.Name, secrets.CAPEM)),
	}
	if err := rootcerts.ConfigureTLS(tls, rootConfig); err != nil {
		return nil, nil, err
	}

	return util.DialExposedGRPCConn(
		context.Background(),
		logger,
		node.ExposedPort(8503),
		token,
		tls,
	)
}
