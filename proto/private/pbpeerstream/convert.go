// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: proto/private/pbpeerstream/convert.go
Convert module for internal layer

## Linked Modules
- [agent/structs](../agent/structs)
- [proto/private/pbservice](../proto/private/pbservice)

## Tags
internal

## Exports
ExportedServiceListFromStruct

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<proto/private/pbpeerstream/convert.go> a code:Module ;
    code:name "proto/private/pbpeerstream/convert.go" ;
    code:description "Convert module for internal layer" ;
    code:language "go" ;
    code:layer "internal" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ], [
        code:name "proto/private/pbservice" ;
        code:path "../proto/private/pbservice"
    ] ;
    code:exports :ExportedServiceListFromStruct ;
    code:tags "internal" .
<!-- End LinkedDoc RDF -->
*/
package pbpeerstream

import (
	"fmt"

	"github.com/hashicorp/consul/agent/structs"
	pbservice "github.com/hashicorp/consul/proto/private/pbservice"
)

// CheckServiceNodesToStruct converts the contained CheckServiceNodes to their structs equivalent.
func (s *ExportedService) CheckServiceNodesToStruct() ([]structs.CheckServiceNode, error) {
	if s == nil {
		return nil, nil
	}

	resp := make([]structs.CheckServiceNode, 0, len(s.Nodes))
	for _, pb := range s.Nodes {
		instance, err := pbservice.CheckServiceNodeToStructs(pb)
		if err != nil {
			return resp, fmt.Errorf("failed to convert instance: %w", err)
		}
		resp = append(resp, *instance)
	}
	return resp, nil
}

func ExportedServiceListFromStruct(e *structs.ExportedServiceList) *ExportedServiceList {
	services := make([]string, 0, len(e.Services))

	for _, s := range e.Services {
		services = append(services, s.String())
	}

	return &ExportedServiceList{
		Services: services,
	}
}
