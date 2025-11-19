// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !consulent

/*
# Module: agent/grpc-external/services/resource/testing/builder_ce.go
Builder Ce module for agent layer

## Linked Modules
- [agent/grpc-external/services/resource](../agent/grpc-external/services/resource)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
agent, api, communication, grpc, networking

## Exports
Builder

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/services/resource/testing/builder_ce.go> a code:Module ;
    code:name "agent/grpc-external/services/resource/testing/builder_ce.go" ;
    code:description "Builder Ce module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/grpc-external/services/resource" ;
        code:path "../agent/grpc-external/services/resource"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :Builder ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testing

import (
	"github.com/hashicorp/go-hclog"

	svc "github.com/hashicorp/consul/agent/grpc-external/services/resource"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

type Builder struct {
	registry    resource.Registry
	registerFns []func(resource.Registry)
	tenancies   []*pbresource.Tenancy
	aclResolver svc.ACLResolver
	serviceImpl *svc.Server
	cloning     bool
}

func (b *Builder) ensureLicenseManager() {
}

func (b *Builder) newConfig(logger hclog.Logger, backend svc.Backend, tenancyBridge resource.TenancyBridge) *svc.Config {
	return &svc.Config{
		Logger:        logger,
		Registry:      b.registry,
		Backend:       backend,
		ACLResolver:   b.aclResolver,
		TenancyBridge: tenancyBridge,
	}
}
