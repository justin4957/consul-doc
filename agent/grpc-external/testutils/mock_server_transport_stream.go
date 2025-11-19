// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/grpc-external/testutils/mock_server_transport_stream.go
Mock Server Transport Stream module for agent layer

## Tags
agent, api, communication, grpc, networking

## Exports
MockServerTransportStream

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/grpc-external/testutils/mock_server_transport_stream.go> a code:Module ;
    code:name "agent/grpc-external/testutils/mock_server_transport_stream.go" ;
    code:description "Mock Server Transport Stream module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :MockServerTransportStream ;
    code:tags "agent", "api", "communication", "grpc", "networking" .
<!-- End LinkedDoc RDF -->
*/
package testutils

import "google.golang.org/grpc/metadata"

type MockServerTransportStream struct {
	MD metadata.MD
}

func (m *MockServerTransportStream) Method() string {
	return ""
}

func (m *MockServerTransportStream) SetHeader(md metadata.MD) error {
	return nil
}

func (m *MockServerTransportStream) SendHeader(md metadata.MD) error {
	m.MD = metadata.Join(m.MD, md)
	return nil
}

func (m *MockServerTransportStream) SetTrailer(md metadata.MD) error {
	return nil
}
