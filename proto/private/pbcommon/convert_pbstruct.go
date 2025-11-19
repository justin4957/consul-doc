// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbcommon/convert_pbstruct.go
Convert Pbstruct module for internal layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
internal

## Exports
EnvoyExtensionsFromStructs, EnvoyExtensionsToStructs, MapStringInterfaceToProtobufTypesStruct, ProtobufTypesStructToMapStringInterface, SliceToPBListValue

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbcommon/convert_pbstruct.go> a code:Module ;
    code:name "proto/private/pbcommon/convert_pbstruct.go" ;
    code:description "Convert Pbstruct module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :EnvoyExtensionsFromStructs, :EnvoyExtensionsToStructs, :MapStringInterfaceToProtobufTypesStruct, :ProtobufTypesStructToMapStringInterface, :SliceToPBListValue ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbcommon

import (
	"github.com/hashicorp/consul/agent/structs"
	"google.golang.org/protobuf/types/known/structpb"
)

// ProtobufTypesStructToMapStringInterface converts a protobuf/structpb.Struct into a
// map[string]interface{}.
func ProtobufTypesStructToMapStringInterface(s *structpb.Struct) map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.AsMap()
}

// MapStringInterfaceToProtobufTypesStruct converts a map[string]interface{} into a proto.Struct
func MapStringInterfaceToProtobufTypesStruct(m map[string]interface{}) *structpb.Struct {
	if len(m) < 1 {
		return nil
	}
	// TODO - handle the error better. It probably requires mog to be able to use alternative method signatures though
	s, _ := structpb.NewStruct(m)
	return s
}

// SliceToPBListValue converts a []interface{} into a proto.ListValue. It's used
// internally by MapStringInterfaceToProtobufTypesStruct when it encouters slices.
// TODO (remove usage of this struct in favor of structpb.NewListValue)
func SliceToPBListValue(s []interface{}) *structpb.ListValue {
	if len(s) < 1 {
		return nil
	}
	// TODO - handle the error better. It probably requires mog to use alt method signatures though
	val, _ := structpb.NewList(s)
	return val
}

// EnvoyExtensionsToStructs takes a protobuf EnvoyExtension argument and converts it to the
// structs EnvoyExtension
func EnvoyExtensionsToStructs(args []*EnvoyExtension) []structs.EnvoyExtension {
	o := make([]structs.EnvoyExtension, len(args))
	for i := range args {
		var e structs.EnvoyExtension
		if args[i] != nil {
			e = structs.EnvoyExtension{
				Name:          args[i].Name,
				Required:      args[i].Required,
				ConsulVersion: args[i].ConsulVersion,
				EnvoyVersion:  args[i].EnvoyVersion,
				Arguments:     ProtobufTypesStructToMapStringInterface(args[i].Arguments),
			}
		}

		o[i] = e
	}

	return o
}

// EnvoyExtensionsFromStructs takes a structs EnvoyExtension argument and converts it to the
// protobuf EnvoyExtension
func EnvoyExtensionsFromStructs(args []structs.EnvoyExtension) []*EnvoyExtension {
	o := make([]*EnvoyExtension, len(args))
	for i, e := range args {
		o[i] = &EnvoyExtension{
			Name:          e.Name,
			Required:      e.Required,
			ConsulVersion: e.ConsulVersion,
			EnvoyVersion:  e.EnvoyVersion,
			Arguments:     MapStringInterfaceToProtobufTypesStruct(e.Arguments),
		}
	}

	return o
}
