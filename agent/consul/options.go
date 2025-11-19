// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/consul/options.go
Options module for agent layer

## Linked Modules
- [agent/consul/stream](../agent/consul/stream)
- [agent/grpc-external/limiter](../agent/grpc-external/limiter)
- [agent/hcp](../agent/hcp)
- [agent/leafcert](../agent/leafcert)
- [agent/pool](../agent/pool)
- [agent/router](../agent/router)
- [agent/rpc/middleware](../agent/rpc/middleware)
- [agent/token](../agent/token)
- [internal/resource](../internal/resource)
- [tlsutil](../tlsutil)

## Tags
agent

## Exports
Deps, GRPCClientConner, LeaderForwarder

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/consul/options.go> a code:Module ;
    code:name "agent/consul/options.go" ;
    code:description "Options module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ], [
        code:name "agent/grpc-external/limiter" ;
        code:path "../agent/grpc-external/limiter"
    ], [
        code:name "agent/hcp" ;
        code:path "../agent/hcp"
    ], [
        code:name "agent/leafcert" ;
        code:path "../agent/leafcert"
    ], [
        code:name "agent/pool" ;
        code:path "../agent/pool"
    ], [
        code:name "agent/router" ;
        code:path "../agent/router"
    ], [
        code:name "agent/rpc/middleware" ;
        code:path "../agent/rpc/middleware"
    ], [
        code:name "agent/token" ;
        code:path "../agent/token"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "tlsutil" ;
        code:path "../tlsutil"
    ] ;
    code:exports :Deps, :GRPCClientConner, :LeaderForwarder ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package consul

import (
	"google.golang.org/grpc"

	"github.com/hashicorp/consul-net-rpc/net/rpc"
	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/agent/grpc-external/limiter"
	"github.com/hashicorp/consul/agent/hcp"
	"github.com/hashicorp/consul/agent/leafcert"
	"github.com/hashicorp/consul/agent/pool"
	"github.com/hashicorp/consul/agent/router"
	"github.com/hashicorp/consul/agent/rpc/middleware"
	"github.com/hashicorp/consul/agent/token"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/tlsutil"
)

type Deps struct {
	LeafCertManager  *leafcert.Manager
	EventPublisher   *stream.EventPublisher
	Logger           hclog.InterceptLogger
	TLSConfigurator  *tlsutil.Configurator
	Tokens           *token.Store
	Router           *router.Router
	ConnPool         *pool.ConnPool
	GRPCConnPool     GRPCClientConner
	LeaderForwarder  LeaderForwarder
	XDSStreamLimiter *limiter.SessionLimiter
	Registry         resource.Registry
	// GetNetRPCInterceptorFunc, if not nil, sets the net/rpc rpc.ServerServiceCallInterceptor on
	// the server side to record metrics around the RPC requests. If nil, no interceptor is added to
	// the rpc server.
	GetNetRPCInterceptorFunc func(recorder *middleware.RequestRecorder) rpc.ServerServiceCallInterceptor
	// NewRequestRecorderFunc provides a middleware.RequestRecorder for the server to use; it cannot be nil
	NewRequestRecorderFunc func(logger hclog.Logger, isLeader func() bool, localDC string) *middleware.RequestRecorder

	// HCP contains the dependencies required when integrating with the HashiCorp Cloud Platform
	HCP hcp.Deps

	Experiments []string

	EnterpriseDeps
}

type GRPCClientConner interface {
	ClientConn(datacenter string) (*grpc.ClientConn, error)
	ClientConnLeader() (*grpc.ClientConn, error)
	SetGatewayResolver(func(string) string)
}

type LeaderForwarder interface {
	// UpdateLeaderAddr updates the leader address in the local DC's resolver.
	UpdateLeaderAddr(datacenter, addr string)
}
