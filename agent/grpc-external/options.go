// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/options.go
Options module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent, api, communication, grpc, networking

## Exports
ContextWithQueryOptions, QueryOptionsFromContext

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/options.go> a code:Module ;
    code:name "agent/grpc-external/options.go" ;
    code:description "Options module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:exports :ContextWithQueryOptions, :QueryOptionsFromContext ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package external

import (
	"context"
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/hashicorp/consul/agent/structs"
	"google.golang.org/grpc/metadata"
)

// QueryOptionsFromContext returns the query options in the gRPC metadata attached to the
// given context.
func QueryOptionsFromContext(ctx context.Context) (structs.QueryOptions, error) {
	options := structs.QueryOptions{}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return options, nil
	}

	m := map[string]string{}
	for k, v := range md {
		m[k] = v[0]
	}

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &options,
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.StringToTimeDurationHookFunc(),
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return structs.QueryOptions{}, err
	}

	err = decoder.Decode(m)
	if err != nil {
		return structs.QueryOptions{}, err
	}

	return options, nil
}

// ContextWithQueryOptions returns a context with the given query options attached.
func ContextWithQueryOptions(ctx context.Context, options structs.QueryOptions) (context.Context, error) {
	md := metadata.MD{}
	m := map[string]interface{}{}
	err := mapstructure.Decode(options, &m)
	if err != nil {
		return nil, err
	}
	for k, v := range m {
		md.Set(k, fmt.Sprintf("%v", v))
	}
	return metadata.NewOutgoingContext(ctx, md), nil
}
