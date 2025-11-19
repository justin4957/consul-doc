// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/protoc-gen-resource-types/main.go
Main module for internal layer

## Linked Modules
- [internal/resource/protoc-gen-resource-types/internal/generate](../internal/resource/protoc-gen-resource-types/internal/generate)

## Tags
data-model, internal, types

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/protoc-gen-resource-types/main.go> a code:Module ;
    code:name "internal/resource/protoc-gen-resource-types/main.go" ;
    code:description "Main module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource/protoc-gen-resource-types/internal/generate" ;
        code:path "../internal/resource/protoc-gen-resource-types/internal/generate"
    ] ;
    code:tags "data-model", "internal", "types" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"flag"

	"github.com/hashicorp/consul/internal/resource/protoc-gen-resource-types/internal/generate"
	"google.golang.org/protobuf/compiler/protogen"
	plugin "google.golang.org/protobuf/types/pluginpb"
)

var (
	file = flag.String("file", "-", "where to load data from")
)

func main() {
	flag.Parse()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {
		gp.SupportedFeatures = uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate.Generate(gp)
	})
}
