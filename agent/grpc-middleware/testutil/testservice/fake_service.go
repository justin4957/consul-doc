// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-middleware/testutil/testservice/fake_service.go
Fake Service module for agent layer

## Tags
agent, api, communication, grpc, networking

## Exports
Simple, SimplePanic

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-middleware/testutil/testservice/fake_service.go> a code:Module ;
    code:name "agent/grpc-middleware/testutil/testservice/fake_service.go" ;
    code:description "Fake Service module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :Simple, :SimplePanic ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testservice

import (
	"context"
	"time"
)

type Simple struct {
	Name string
	DC   string
}

func (s *Simple) Flow(_ *Req, flow Simple_FlowServer) error {
	for flow.Context().Err() == nil {
		resp := &Resp{ServerName: "one", Datacenter: s.DC}
		if err := flow.Send(resp); err != nil {
			return err
		}
		time.Sleep(time.Millisecond)
	}
	return nil
}

func (s *Simple) Something(_ context.Context, _ *Req) (*Resp, error) {
	return &Resp{ServerName: s.Name, Datacenter: s.DC}, nil
}

type SimplePanic struct {
	Name, DC string
}

func (s *SimplePanic) Flow(_ *Req, flow Simple_FlowServer) error {
	for flow.Context().Err() == nil {
		time.Sleep(time.Millisecond)
		panic("panic from Flow")
	}
	return nil
}

func (s *SimplePanic) Something(_ context.Context, _ *Req) (*Resp, error) {
	time.Sleep(time.Millisecond)
	panic("panic from Something")
}
