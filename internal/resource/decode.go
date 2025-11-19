// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/decode.go
Decode module for internal layer

## Linked Modules
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
Decode, DecodeList, DecodedResource, GetDecodedResource, ListDecodedResource

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/decode.go> a code:Module ;
    code:name "internal/resource/decode.go" ;
    code:description "Decode module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :Decode, :DecodeList, :DecodedResource, :GetDecodedResource, :ListDecodedResource ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resource

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/proto-public/pbresource"
)

// DecodedResource is a generic holder to contain an original Resource and its
// decoded contents.
type DecodedResource[T proto.Message] struct {
	// Embedding here allows us to shadow the Resource.Data Any field to fake out
	// using a single struct with inlined data.
	*pbresource.Resource
	Data T
}

func (d *DecodedResource[T]) GetResource() *pbresource.Resource {
	if d == nil {
		return nil
	}

	return d.Resource
}

func (d *DecodedResource[T]) GetData() T {
	if d == nil {
		var zero T
		return zero
	}

	return d.Data
}

// Decode will generically decode the provided resource into a 2-field
// structure that holds onto the original Resource and the decoded contents.
//
// Returns an ErrDataParse on unmarshalling errors.
func Decode[T proto.Message](res *pbresource.Resource) (*DecodedResource[T], error) {
	var zero T
	data := zero.ProtoReflect().New().Interface().(T)
	// check that there is data to unmarshall
	if res.Data != nil {
		if err := res.Data.UnmarshalTo(data); err != nil {
			return nil, NewErrDataParse(data, err)
		}
	}
	return &DecodedResource[T]{
		Resource: res,
		Data:     data,
	}, nil
}

// DecodeList will generically decode the provided resource list into a list of 2-field
// structures that holds onto the original Resource and the decoded contents.
//
// Returns an ErrDataParse on unmarshalling errors.
func DecodeList[T proto.Message](resources []*pbresource.Resource) ([]*DecodedResource[T], error) {
	var decoded []*DecodedResource[T]
	for _, res := range resources {
		d, err := Decode[T](res)
		if err != nil {
			return nil, err
		}
		decoded = append(decoded, d)
	}
	return decoded, nil
}

// GetDecodedResource will generically read the requested resource using the
// client and either return nil on a NotFound or decode the response value.
func GetDecodedResource[T proto.Message](ctx context.Context, client pbresource.ResourceServiceClient, id *pbresource.ID) (*DecodedResource[T], error) {
	rsp, err := client.Read(ctx, &pbresource.ReadRequest{Id: id})
	switch {
	case status.Code(err) == codes.NotFound:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return Decode[T](rsp.Resource)
}

func ListDecodedResource[T proto.Message](ctx context.Context, client pbresource.ResourceServiceClient, req *pbresource.ListRequest) ([]*DecodedResource[T], error) {
	rsp, err := client.List(ctx, req)
	if err != nil {
		return nil, err
	}

	results := make([]*DecodedResource[T], len(rsp.Resources))
	for idx, rsc := range rsp.Resources {
		d, err := Decode[T](rsc)
		if err != nil {
			return nil, err
		}
		results[idx] = d
	}
	return results, nil
}
