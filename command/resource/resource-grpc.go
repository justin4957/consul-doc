// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/resource/resource-grpc.go
Resource Grpc module for cli layer

## Linked Modules
- [command/resource/client](../command/resource/client)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
api, cli, communication, grpc, networking, user-interface

## Exports
ResourceGRPC

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/resource/resource-grpc.go> a code:Module ;
    code:name "command/resource/resource-grpc.go" ;
    code:description "Resource Grpc module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/resource/client" ;
        code:path "../command/resource/client"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :ResourceGRPC ;
    code:tags "api", "cli", "communication", "grpc", "networking", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/hashicorp/consul/command/resource/client"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

const (
	HeaderConsulToken = "x-consul-token"
)

type ResourceGRPC struct {
	C *client.GRPCClient
}

func (resource *ResourceGRPC) Apply(parsedResource *pbresource.Resource) (*pbresource.Resource, error) {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	writeRsp, err := resource.C.Client.Write(ctx, &pbresource.WriteRequest{Resource: parsedResource})
	if err != nil {
		return nil, fmt.Errorf("error writing resource: %+v", err)
	}

	return writeRsp.Resource, err
}

func (resource *ResourceGRPC) Read(resourceType *pbresource.Type, resourceTenancy *pbresource.Tenancy, resourceName string, stale bool) (*pbresource.Resource, error) {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if !stale {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-consul-consistency-mode", "consistent")
	}
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	readRsp, err := resource.C.Client.Read(ctx, &pbresource.ReadRequest{
		Id: &pbresource.ID{
			Type:    resourceType,
			Tenancy: resourceTenancy,
			Name:    resourceName,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error reading resource: %+v", err)
	}

	return readRsp.Resource, err
}

func (resource *ResourceGRPC) List(resourceType *pbresource.Type, resourceTenancy *pbresource.Tenancy, prefix string, stale bool) ([]*pbresource.Resource, error) {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if !stale {
		ctx = metadata.AppendToOutgoingContext(ctx, "x-consul-consistency-mode", "consistent")
	}
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	listRsp, err := resource.C.Client.List(ctx, &pbresource.ListRequest{
		Type:       resourceType,
		Tenancy:    resourceTenancy,
		NamePrefix: prefix,
	})

	if err != nil {
		return nil, fmt.Errorf("error listing resource: %+v", err)
	}

	return listRsp.Resources, err
}

func (resource *ResourceGRPC) Delete(resourceType *pbresource.Type, resourceTenancy *pbresource.Tenancy, resourceName string) error {
	token, err := resource.C.Config.GetToken()
	if err != nil {
		return err
	}
	ctx := context.Background()
	if token != "" {
		ctx = metadata.AppendToOutgoingContext(context.Background(), HeaderConsulToken, token)
	}

	defer resource.C.Conn.Close()
	_, err = resource.C.Client.Delete(ctx, &pbresource.DeleteRequest{
		Id: &pbresource.ID{
			Type:    resourceType,
			Tenancy: resourceTenancy,
			Name:    resourceName,
		},
	})

	if err != nil {
		return fmt.Errorf("error deleting resource: %+v", err)
	}

	return nil
}
