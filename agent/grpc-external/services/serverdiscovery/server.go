// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/services/serverdiscovery/server.go
Server module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/consul/stream](../agent/consul/stream)
- [proto-public/pbserverdiscovery](../proto-public/pbserverdiscovery)

## Tags
agent, api, communication, discovery, dns, grpc, networking

## Exports
ACLResolver, Config, EventPublisher, NewServer, Server

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/serverdiscovery/server.go> a code:Module ;
    code:name "agent/grpc-external/services/serverdiscovery/server.go" ;
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
        code:name "proto-public/pbserverdiscovery" ;
        code:path "../proto-public/pbserverdiscovery"
    ] ;
    code:exports :ACLResolver, :Config, :EventPublisher, :NewServer, :Server ;
    code:tags "agent", "api", "communication", "discovery", "dns", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package serverdiscovery

import (
	"google.golang.org/grpc"

	"github.com/hashicorp/go-hclog"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/agent/consul/stream"
	"github.com/hashicorp/consul/proto-public/pbserverdiscovery"
)

type Server struct {
	Config
}

type Config struct {
	Publisher   EventPublisher
	Logger      hclog.Logger
	ACLResolver ACLResolver
}

type EventPublisher interface {
	Subscribe(*stream.SubscribeRequest) (*stream.Subscription, error)
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Register(registrar grpc.ServiceRegistrar) {
	pbserverdiscovery.RegisterServerDiscoveryServiceServer(registrar, s)
}
