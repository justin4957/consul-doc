// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/testutils/server.go
Server module for agent layer

## Tags
agent, api, communication, grpc, networking

## Exports
GRPCService, RunTestServer

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/testutils/server.go> a code:Module ;
    code:name "agent/grpc-external/testutils/server.go" ;
    code:description "Server module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :GRPCService, :RunTestServer ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testutils

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type GRPCService interface {
	Register(grpc.ServiceRegistrar)
}

func RunTestServer(t *testing.T, services ...GRPCService) net.Addr {
	t.Helper()

	lis, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	grpcServer := grpc.NewServer()
	for _, svc := range services {
		svc.Register(grpcServer)
	}

	go grpcServer.Serve(lis)
	t.Cleanup(grpcServer.Stop)

	return lis.Addr()
}
