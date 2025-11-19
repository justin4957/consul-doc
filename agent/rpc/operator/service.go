// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/rpc/operator/service.go
Service module for agent layer

## Linked Modules
- [acl](../acl)
- [acl/resolver](../acl/resolver)
- [agent/grpc-external](../agent/grpc-external)
- [agent/structs](../agent/structs)
- [proto/private/pboperator](../proto/private/pboperator)

## Tags
agent, communication, networking

## Exports
Backend, Config, NewServer, Server

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/rpc/operator/service.go> a code:Module ;
    code:name "agent/rpc/operator/service.go" ;
    code:description "Service module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "acl/resolver" ;
        code:path "../acl/resolver"
    ], [
        code:name "agent/grpc-external" ;
        code:path "../agent/grpc-external"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pboperator" ;
        code:path "../proto/private/pboperator"
    ] ;
    code:exports :Backend, :Config, :NewServer, :Server ;
    code:tags "agent", "communication", "networking" .
<!-- End LinkedDoc RDF -->
*/
package operator

import (
	"context"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	external "github.com/hashicorp/consul/agent/grpc-external"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/proto/private/pboperator"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

// For private/internal gRPC handlers, protoc-gen-rpc-glue generates the
// requisite methods to satisfy the structs.RPCInfo interface using fields
// from the pbcommon package. This service is public, so we can't use those
// fields in our proto definition. Instead, we construct our RPCInfo manually.
var writeRequest struct {
	structs.WriteRequest
	structs.DCSpecificRequest
}

var readRequest struct {
	structs.QueryOptions
	structs.DCSpecificRequest
}

// Server implements pboperator.OperatorService to provide RPC operations for
// managing operator operation.
type Server struct {
	Config
}

func (s *Server) TransferLeader(ctx context.Context, request *pboperator.TransferLeaderRequest) (*pboperator.TransferLeaderResponse, error) {
	resp := &pboperator.TransferLeaderResponse{Success: false}
	handled, err := s.ForwardRPC(&writeRequest, func(conn *grpc.ClientConn) error {
		ctx := external.ForwardMetadataContext(ctx)
		var err error
		resp, err = pboperator.NewOperatorServiceClient(conn).TransferLeader(ctx, request)
		return err
	})
	if handled || err != nil {
		return resp, err
	}

	var authzCtx acl.AuthorizerContext
	entMeta := structs.DefaultEnterpriseMetaInDefaultPartition()

	options, err := external.QueryOptionsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	authz, err := s.Backend.ResolveTokenAndDefaultMeta(options.Token, entMeta, &authzCtx)
	if err != nil {
		return resp, err
	}

	if err := authz.ToAllowAuthorizer().OperatorWriteAllowed(&authzCtx); err != nil {
		return resp, err
	}

	return s.Backend.TransferLeader(ctx, request)
}

type Config struct {
	Backend    Backend
	Logger     hclog.Logger
	ForwardRPC func(structs.RPCInfo, func(*grpc.ClientConn) error) (bool, error)
	Datacenter string
}

func NewServer(cfg Config) *Server {
	requireNotNil(cfg.Backend, "Backend")
	requireNotNil(cfg.Logger, "Logger")
	requireNotNil(cfg.ForwardRPC, "ForwardRPC")
	if cfg.Datacenter == "" {
		panic("Datacenter is required")
	}
	return &Server{
		Config: cfg,
	}
}

func requireNotNil(v interface{}, name string) {
	if v == nil {
		panic(name + " is required")
	}
}

var _ pboperator.OperatorServiceServer = (*Server)(nil)

func (s *Server) Register(registrar grpc.ServiceRegistrar) {
	pboperator.RegisterOperatorServiceServer(registrar, s)
}

// Backend defines the core integrations the Operator endpoint depends on. A
// functional implementation will integrate with various operator operation such as
// raft, autopilot operation. The only currently implemented operation is raft leader transfer
type Backend interface {
	TransferLeader(ctx context.Context, request *pboperator.TransferLeaderRequest) (*pboperator.TransferLeaderResponse, error)
	ResolveTokenAndDefaultMeta(token string, entMeta *acl.EnterpriseMeta, authzCtx *acl.AuthorizerContext) (resolver.Result, error)
}
