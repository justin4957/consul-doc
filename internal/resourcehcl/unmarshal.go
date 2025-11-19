// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resourcehcl/unmarshal.go
Unmarshal module for internal layer

## Linked Modules
- [internal/protohcl](../internal/protohcl)
- [internal/resource](../internal/resource)
- [proto-public/pbresource](../proto-public/pbresource)

## Tags
internal

## Exports
Unmarshal, UnmarshalOptions

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resourcehcl/unmarshal.go> a code:Module ;
    code:name "internal/resourcehcl/unmarshal.go" ;
    code:description "Unmarshal module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/protohcl" ;
        code:path "../internal/protohcl"
    ], [
        code:name "internal/resource" ;
        code:path "../internal/resource"
    ], [
        code:name "proto-public/pbresource" ;
        code:path "../proto-public/pbresource"
    ] ;
    code:exports :Unmarshal, :UnmarshalOptions ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package resourcehcl

import (
	"reflect"

	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"

	"github.com/hashicorp/consul/internal/protohcl"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

// Unmarshal the given HCL source into a resource.
func Unmarshal(src []byte, reg resource.Registry) (*pbresource.Resource, error) {
	return UnmarshalOptions{}.Unmarshal(src, reg)
}

type UnmarshalOptions struct{ SourceFileName string }

// Unmarshal the given HCL source into a resource.
func (u UnmarshalOptions) Unmarshal(src []byte, reg resource.Registry) (*pbresource.Resource, error) {
	var out pbresource.Resource
	err := (protohcl.UnmarshalOptions{
		SourceFileName: u.SourceFileName,
		AnyTypeProvider: anyProvider{
			base: &protohcl.AnyTypeURLProvider{TypeURLFieldName: "Type"},
			reg:  reg,
		},
		FieldNamer: fieldNamer{acroynms: []string{"ID", "TCP", "UDP", "HTTP"}},
		Functions:  map[string]function.Function{"gvk": gvk},
	}).Unmarshal(src, &out)
	return &out, err
}

var (
	typeType = cty.Capsule("type", reflect.TypeOf(pbresource.Type{}))

	gvk = function.New(&function.Spec{
		Params: []function.Parameter{
			{Name: "GVK String", Type: cty.String},
		},
		Type: function.StaticReturnType(typeType),
		Impl: func(args []cty.Value, _ cty.Type) (cty.Value, error) {
			t, err := resource.ParseGVK(args[0].AsString())
			if err != nil {
				return cty.NilVal, err
			}
			return cty.CapsuleVal(typeType, t), nil
		},
	})
)
