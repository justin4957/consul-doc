// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/xds/protocol_trace.go
Protocol Trace module for agent layer

## Tags
agent, envoy, service-mesh

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/xds/protocol_trace.go> a code:Module ;
    code:name "agent/xds/protocol_trace.go" ;
    code:description "Protocol Trace module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "envoy", "service-mesh" .
<!-- End LinkedDoc RDF -->
*/
package xds

import (
	envoy_discovery_v3 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/hashicorp/go-hclog"

	"github.com/mitchellh/copystructure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func logTraceRequest(logger hclog.Logger, msg string, pb proto.Message) {
	logTraceProto(logger, msg, pb, false)
}

func logTraceResponse(logger hclog.Logger, msg string, pb proto.Message) {
	logTraceProto(logger, msg, pb, true)
}

func logTraceProto(logger hclog.Logger, msg string, pb proto.Message, response bool) {
	if !logger.IsTrace() {
		return
	}

	dir := "request"
	if response {
		dir = "response"
	}

	// Duplicate the request so we can scrub the huge Node field for logging.
	// If the cloning fails, then log anyway but don't scrub the node field.
	if dup, err := copystructure.Copy(pb); err == nil {
		pb = dup.(proto.Message)

		// strip the node field
		switch x := pb.(type) {
		case *envoy_discovery_v3.DiscoveryRequest:
			x.Node = nil
		case *envoy_discovery_v3.DeltaDiscoveryRequest:
			x.Node = nil
		}
	}

	m := protojson.MarshalOptions{
		Indent: "  ",
	}

	out := ""
	outBytes, err := m.Marshal(pb)
	if err != nil {
		out = "<ERROR: " + err.Error() + ">"
	} else {
		out = string(outBytes)
	}

	logger.Trace(msg, "direction", dir, "protobuf", out)
}
