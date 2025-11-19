// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/structs/envoy_extension.go
Envoy Extension module for agent layer

## Linked Modules
- [api](../api)

## Tags
agent, data-model, types

## Exports
EnvoyExtension, EnvoyExtensions

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/structs/envoy_extension.go> a code:Module ;
    code:name "agent/structs/envoy_extension.go" ;
    code:description "Envoy Extension module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "api" ;
        code:path "../api"
    ] ;
    code:exports :EnvoyExtension, :EnvoyExtensions ;
    code:tags "agent", "data-model", "types" .
<!-- End LinkedDoc RDF -->
*/
package structs

import (
	"github.com/hashicorp/consul/api"
)

// EnvoyExtension has configuration for an extension that patches Envoy resources.
type EnvoyExtension struct {
	Name          string
	Required      bool
	Arguments     map[string]interface{} `bexpr:"-"`
	ConsulVersion string
	EnvoyVersion  string
}

type EnvoyExtensions []EnvoyExtension

func (es EnvoyExtensions) ToAPI() []api.EnvoyExtension {
	extensions := make([]api.EnvoyExtension, len(es))
	for i, e := range es {
		extensions[i] = api.EnvoyExtension{
			Name:          e.Name,
			Required:      e.Required,
			Arguments:     e.Arguments,
			EnvoyVersion:  e.EnvoyVersion,
			ConsulVersion: e.ConsulVersion,
		}
	}
	return extensions
}
