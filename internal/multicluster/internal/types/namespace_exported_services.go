// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/namespace_exported_services.go
Namespace Exported Services module for internal layer

## Linked Modules
- [acl](../acl)
- [internal/resource](../internal/resource)
- [proto-public/pbmulticluster/v2](../proto-public/pbmulticluster/v2)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
data-model, internal, types

## Exports
RegisterNamespaceExportedServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/namespace_exported_services.go> a code:Module ;
    code:name "internal/multicluster/internal/types/namespace_exported_services.go" ;
    code:description "Namespace Exported Services module for internal layer" ;
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
    code:exports :RegisterNamespaceExportedServices ;
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

func RegisterNamespaceExportedServices(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbmulticluster.NamespaceExportedServicesType,
		Proto:    &pbmulticluster.NamespaceExportedServices{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateNamespaceExportedServices,
		ACLs: &resource.ACLHooks{
			Read:  aclReadHookNamespaceExportedServices,
			Write: aclWriteHookNamespaceExportedServices,
			List:  resource.NoOpACLListHook,
		},
	})
}

func aclReadHookNamespaceExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, id *pbresource.ID, res *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshReadAllowed(authzContext)
}

func aclWriteHookNamespaceExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, res *pbresource.Resource) error {
	return authorizer.ToAllowAuthorizer().MeshWriteAllowed(authzContext)
}
