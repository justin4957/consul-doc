// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: internal/resource/protoc-gen-deepcopy/main.go
Main module for internal layer

## Linked Modules
- [internal/resource/protoc-gen-deepcopy/internal/generate](../internal/resource/protoc-gen-deepcopy/internal/generate)

## Tags
internal

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<internal/resource/protoc-gen-deepcopy/main.go> a code:Module ;
    code:name "internal/resource/protoc-gen-deepcopy/main.go" ;
    code:description "Main module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "internal/resource/protoc-gen-deepcopy/internal/generate" ;
        code:path "../internal/resource/protoc-gen-deepcopy/internal/generate"
    ] ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	plugin "google.golang.org/protobuf/types/pluginpb"

	"github.com/hashicorp/consul/internal/resource/protoc-gen-deepcopy/internal/generate"
)

var (
	file = flag.String("file", "-", "where to load data from")
)

// This file is responsible for generating a DeepCopy and DeepCopyInto overwrite for proto
// structs which allows Kubernetes CRDs to get created directly from the proto-types.
func main() {
	flag.Parse()

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gp *protogen.Plugin) error {
		gp.SupportedFeatures = uint64(plugin.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate.Generate(gp)
	})
}
