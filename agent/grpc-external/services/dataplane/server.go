// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/services/dataplane/server.go
Dataplane gRPC service for Envoy proxy configuration and bootstrap parameters.

## Linked Modules
- [agent/grpc-external/server](../../server.go) - gRPC server infrastructure
- [agent/proxycfg](../../../proxycfg/manager.go) - Proxy configuration manager
- [agent/xds](../../../xds/server.go) - xDS server

## Tags
grpc, service, dataplane, envoy, proxy

## Exports
Server, NewServer, Config

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/dataplane/server.go> a code:Module ;
    code:language "go" ;
    code:name "agent/grpc-external/services/dataplane/server.go" ;
    code:description "Dataplane gRPC service for Envoy proxy configuration and bootstrap parameters" ;
    code:layer "grpc" ;
    code:linksTo [
        code:name "agent/grpc-external/server" ;
        code:path "../../server.go" ;
        code:relationship "gRPC server infrastructure"
    ], [
        code:name "agent/proxycfg" ;
        code:path "../../../proxycfg/manager.go" ;
        code:relationship "Proxy configuration manager"
    ], [
        code:name "agent/xds" ;
        code:path "../../../xds/server.go" ;
        code:relationship "xDS server"
    ] ;
    code:exports :Server, :NewServer, :Config ;
    code:tags "grpc", "service", "dataplane", "envoy", "proxy" .
<!-- End LinkedDoc RDF -->
*/
package dataplane

import (
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/configentry"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto-public/pbdataplane"
)

type Server struct {
	Config
}

type Config struct {
	GetStore    func() StateStore
	Logger      hclog.Logger
	ACLResolver ACLResolver
	// Datacenter of the Consul server this gRPC server is hosted on
	Datacenter string
}

type StateStore interface {
	ServiceNode(string, string, string, *acl.EnterpriseMeta, string) (uint64, *structs.ServiceNode, error)
	ReadResolvedServiceConfigEntries(memdb.WatchSet, string, *acl.EnterpriseMeta, []structs.ServiceID, structs.ProxyMode) (uint64, *configentry.ResolvedServiceConfigSet, error)
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

var _ pbdataplane.DataplaneServiceServer = (*Server)(nil)

func (s *Server) Register(registrar grpc.ServiceRegistrar) {
	pbdataplane.RegisterDataplaneServiceServer(registrar, s)
}
