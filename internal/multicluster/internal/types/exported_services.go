// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/multicluster/internal/types/exported_services.go
Exported Services module for internal layer

## Linked Modules
- [acl](../acl)
- [internal/resource](../internal/resource)
- [proto-public/pbmulticluster/v2](../proto-public/pbmulticluster/v2)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
data-model, internal, types

## Exports
RegisterExportedServices

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/multicluster/internal/types/exported_services.go> a code:Module ;
    code:name "internal/multicluster/internal/types/exported_services.go" ;
    code:description "Exported Services module for internal layer" ;
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
    code:exports :RegisterExportedServices ;
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

func RegisterExportedServices(r resource.Registry) {
	r.Register(resource.Registration{
		Type:     pbmulticluster.ExportedServicesType,
		Proto:    &pbmulticluster.ExportedServices{},
		Scope:    resource.ScopeNamespace,
		Validate: ValidateExportedServices,
		ACLs: &resource.ACLHooks{
			Read:  aclReadHookExportedServices,
			Write: aclWriteHookExportedServices,
			List:  resource.NoOpACLListHook,
		},
	})
}

func aclReadHookExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, _ *pbresource.ID, res *pbresource.Resource) error {
	if res == nil {
		return resource.ErrNeedResource
	}

	var exportedService pbmulticluster.ExportedServices

	if err := res.Data.UnmarshalTo(&exportedService); err != nil {
		return resource.NewErrDataParse(&exportedService, err)
	}

	for _, serviceName := range exportedService.Services {
		if err := authorizer.ToAllowAuthorizer().ServiceReadAllowed(serviceName, authzContext); err != nil {
			return err
		}
	}
	return nil
}

func aclWriteHookExportedServices(authorizer acl.Authorizer, authzContext *acl.AuthorizerContext, res *pbresource.Resource) error {
	var exportedService pbmulticluster.ExportedServices

	if err := res.Data.UnmarshalTo(&exportedService); err != nil {
		return resource.NewErrDataParse(&exportedService, err)
	}

	for _, serviceName := range exportedService.Services {
		if err := authorizer.ToAllowAuthorizer().ServiceWriteAllowed(serviceName, authzContext); err != nil {
			return err
		}
	}
	return nil
}
