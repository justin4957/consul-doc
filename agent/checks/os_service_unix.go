// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: agent/checks/os_service_unix.go
Os Service Unix module for agent layer

## Tags
agent, health-checks, monitoring

## Exports
NewOSServiceClient, OSServiceClient

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/checks/os_service_unix.go> a code:Module ;
    code:name "agent/checks/os_service_unix.go" ;
    code:description "Os Service Unix module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :NewOSServiceClient, :OSServiceClient ;
    code:tags "agent", "health-checks", "monitoring" .
<!-- End LinkedDoc RDF -->
*/
package checks

import "fmt"

type OSServiceClient struct {
}

func NewOSServiceClient() (*OSServiceClient, error) {
	return nil, fmt.Errorf("not implemented")
}

func (client *OSServiceClient) Check(serviceName string) error {
	return fmt.Errorf("not implemented")
}
