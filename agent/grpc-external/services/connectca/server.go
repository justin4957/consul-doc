// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/services/connectca/server.go
Server module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/consul/stream](../agent/consul/stream)
- [agent/structs](../agent/structs)
- [proto-public/pbconnectca](../proto-public/pbconnectca)

## Tags
agent, api, communication, grpc, mtls, networking, service-mesh

## Exports
ACLResolver, CAManager, Config, EventPublisher, NewServer, Server, StateStore

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/connectca/server.go> a code:Module ;
    code:name "agent/grpc-external/services/connectca/server.go" ;
    code:description "Server module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/consul/stream" ;
        code:path "../agent/consul/stream"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto-public/pbconnectca" ;
        code:path "../proto-public/pbconnectca"
    ] ;
    code:exports :ACLResolver, :CAManager, :Config, :EventPublisher, :NewServer, :Server, :StateStore ;
    code:tags "agent", "api", "communication", "grpc", "mtls", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package connectca

import (
	"crypto/x509"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-memdb"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto-public/pbconnectca"
)

type Server struct {
	Config
}

type Config struct {
	Publisher      EventPublisher
	GetStore       func() StateStore
	Logger         hclog.Logger
	ACLResolver    ACLResolver
	CAManager      CAManager
	ForwardRPC     func(structs.RPCInfo, func(*grpc.ClientConn) error) (bool, error)
	ConnectEnabled bool
}

type EventPublisher interface {
	Subscribe(*stream.SubscribeRequest) (*stream.Subscription, error)
}

type StateStore interface {
	CAConfig(memdb.WatchSet) (uint64, *structs.CAConfiguration, error)
	AbandonCh() <-chan struct{}
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(token string, entMeta *acl.EnterpriseMeta, authzContext *acl.AuthorizerContext) (resolver.Result, error)
}

//go:generate mockery --name CAManager --inpackage
type CAManager interface {
	AuthorizeAndSignCertificate(csr *x509.CertificateRequest, authz acl.Authorizer) (*structs.IssuedCert, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Register(registrar grpc.ServiceRegistrar) {
	pbconnectca.RegisterConnectCAServiceServer(registrar, s)
}

func (s *Server) requireConnect() error {
	if s.ConnectEnabled {
		return nil
	}
	return status.Error(codes.FailedPrecondition, "Connect must be enabled in order to use this endpoint")
}
