// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/computed_exported_services.go
Computed Exported Services module for internal layer

## Linked Modules
- [acl](../acl)
- [internal/resource](../internal/resource)
- [proto-public/pbmulticluster/v2](../proto-public/pbmulticluster/v2)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
data-model, internal, types

## Exports
RegisterComputedExportedServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/computed_exported_services.go> a code:Module ;
    code:name "internal/multicluster/internal/types/computed_exported_services.go" ;
    code:description "Computed Exported Services module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "acl" ;
        code:path "../acl"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbmulticluster/v2" ;
        code:path "../proto-public/pbmulticluster/v2"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :RegisterComputedExportedServices ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package types

import (
	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/internal/resource"
	pbmulticluster "github.com/hashicorp/consul/proto-public/pbmulticluster/v2"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	ComputedExportedServicesName = "global"
)

func RegisterComputedExportedServices(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbmulticluster.ComputedExportedServicesType,
		Proto:    &pbmulticluster.ComputedExportedServices{},
		Scope:    resource.ScopePartition,
		Validate: ValidateComputedExportedServices,
		ACLs: &resource.ACLHooks{
			Read:  aclReadHookComputedExportedServices,
			Write: aclWriteHookComputedExportedServices,
			List:  resource.NoOpACLListHook,
		},
	})
}

func aclReadHookComputedExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, _ *pbresource.ID, res *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshReadAllowed(authzContext)
}

func aclWriteHookComputedExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, _ *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshWriteAllowed(authzContext)
}
