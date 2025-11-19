// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/snapshot_endpoint.go
Snapshot Endpoint module for agent layer

## Linked Modules
- [agent/structs](../agent/structs)

## Tags
agent

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/snapshot_endpoint.go> a code:Module ;
    code:name "agent/snapshot_endpoint.go" ;
    code:description "Snapshot Endpoint module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/structs" ;
        code:path "../agent/structs"
    ] ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import (
	"bytes"
	"net/http"

	"github.com/hashicorp/consul/agent/structs"
)

// Snapshot handles requests to take and restore snapshots. This uses a special
// mechanism to make the RPC since we potentially stream large amounts of data
// as part of these requests.
func (s *HTTPHandlers) Snapshot(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	var args structs.SnapshotRequest
	s.parseDC(req, &args.Datacenter)
	s.parseToken(req, &args.Token)
	if _, ok := req.URL.Query()["stale"]; ok {
		args.AllowStale = true
	}

	switch req.Method {
	case "GET":
		args.Op = structs.SnapshotSave

		// Headers need to go out before we stream the body.
		replyFn := func(reply *structs.SnapshotResponse) error {
			setMeta(resp, &reply.QueryMeta)
			return nil
		}

		// Don't bother sending any request body through since it will
		// be ignored.
		var null bytes.Buffer
		if err := s.agent.delegate.SnapshotRPC(&args, &null, resp, replyFn); err != nil {
			return nil, err
		}
		return nil, nil

	case "PUT":
		args.Op = structs.SnapshotRestore
		if err := s.agent.delegate.SnapshotRPC(&args, req.Body, resp, nil); err != nil {
			return nil, err
		}
		return nil, nil

	default:
		return nil, MethodNotAllowedError{req.Method, []string{"GET", "PUT"}}
	}
}
