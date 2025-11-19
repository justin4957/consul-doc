// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/partition_exported_services.go
Partition Exported Services module for internal layer

## Linked Modules
- [acl](../acl)
- [internal/resource](../internal/resource)
- [proto-public/pbmulticluster/v2](../proto-public/pbmulticluster/v2)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
data-model, internal, types

## Exports
RegisterPartitionExportedServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/partition_exported_services.go> a code:Module ;
    code:name "internal/multicluster/internal/types/partition_exported_services.go" ;
    code:description "Partition Exported Services module for internal layer" ;
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
    code:exports :RegisterPartitionExportedServices ;
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

func RegisterPartitionExportedServices(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbmulticluster.PartitionExportedServicesType,
		Proto:    &pbmulticluster.PartitionExportedServices{},
		Scope:    resource.ScopePartition,
		Validate: ValidatePartitionExportedServices,
		ACLs: &resource.ACLHooks{
			Read:  aclReadHookPartitionExportedServices,
			Write: aclWriteHookPartitionExportedServices,
			List:  resource.NoOpACLListHook,
		},
	})
}

func aclReadHookPartitionExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, id *pbresource.ID, res *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshReadAllowed(authzContext)
}

func aclWriteHookPartitionExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, res *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshWriteAllowed(authzContext)
}
