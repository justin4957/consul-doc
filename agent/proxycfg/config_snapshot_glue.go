// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/proxycfg/config_snapshot_glue.go
Config Snapshot Glue module for agent layer

## Linked Modules
- [acl](../acl)
- [agent/structs](../agent/structs)
- [logging](../logging)

## Tags
agent, configuration, networking, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/proxycfg/config_snapshot_glue.go> a code:Module ;
    code:name "agent/proxycfg/config_snapshot_glue.go" ;
    code:description "Config Snapshot Glue module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "logging" ;
        code:path "../logging"
    ] ;
    code:tags "agent", "configuration", "networking", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package proxycfg

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The below functions are added to ConfigSnapshot to allow it to conform to
// the ProxySnapshot interface.
func (s *ConfigSnapshot) AllowEmptyListeners() bool {
	// Ingress and API gateways are allowed to inform LDS of no listeners.
	return s.Kind == structs.ServiceKindIngressGateway ||
		s.Kind == structs.ServiceKindAPIGateway
}

func (s *ConfigSnapshot) AllowEmptyRoutes() bool {
	// Ingress and API gateways are allowed to inform RDS of no routes.
	return s.Kind == structs.ServiceKindIngressGateway ||
		s.Kind == structs.ServiceKindAPIGateway
}

func (s *ConfigSnapshot) AllowEmptyClusters() bool {
	// Mesh, Ingress, API and Terminating gateways are allowed to inform CDS of no clusters.
	return s.Kind == structs.ServiceKindMeshGateway ||
		s.Kind == structs.ServiceKindTerminatingGateway ||
		s.Kind == structs.ServiceKindIngressGateway ||
		s.Kind == structs.ServiceKindAPIGateway
}

func (s *ConfigSnapshot) Authorize(authz acl.Authorizer) error {
	var authzContext acl.AuthorizerContext
	switch s.Kind {
	case structs.ServiceKindConnectProxy:
		s.ProxyID.FillAuthzContext(&authzContext)
		if err := authz.ToAllowAuthorizer().ServiceWriteAllowed(s.Proxy.DestinationServiceName, &authzContext); err != nil {
			return status.Error(codes.PermissionDenied, err.Error())
		}
	case structs.ServiceKindMeshGateway, structs.ServiceKindTerminatingGateway, structs.ServiceKindIngressGateway, structs.ServiceKindAPIGateway:
		s.ProxyID.FillAuthzContext(&authzContext)
		if err := authz.ToAllowAuthorizer().ServiceWriteAllowed(s.Service, &authzContext); err != nil {
			return status.Error(codes.PermissionDenied, err.Error())
		}
	default:
		return status.Error(codes.Internal, "Invalid service kind")
	}

	// Authed OK!
	return nil
}

func (s *ConfigSnapshot) LoggerName() string {
	switch s.Kind {
	case structs.ServiceKindConnectProxy:
	case structs.ServiceKindTerminatingGateway:
		return logging.TerminatingGateway
	case structs.ServiceKindMeshGateway:
		return logging.MeshGateway
	case structs.ServiceKindIngressGateway:
		return logging.IngressGateway
	}

	return ""
}
