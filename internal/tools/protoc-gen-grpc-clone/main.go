// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/tools/protoc-gen-grpc-clone/main.go
Main module for internal layer

## Linked Modules
- [internal/tools/protoc-gen-grpc-clone/internal/generate](../internal/tools/protoc-gen-grpc-clone/internal/generate)

## Tags
api, communication, grpc, internal, networking

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/tools/protoc-gen-grpc-clone/main.go> a code:Module ;
    code:name "internal/tools/protoc-gen-grpc-clone/main.go" ;
    code:description "Main module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/tools/protoc-gen-grpc-clone/internal/generate" ;
        code:path "../internal/tools/protoc-gen-grpc-clone/internal/generate"
    ] ;
    code:tags "api", "communication", "grpc", "internal", "networking" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"flag"

	"github.com/hashicorp/consul/internal/tools/protoc-gen-grpc-clone/internal/generate"
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
