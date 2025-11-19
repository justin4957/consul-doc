// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/checks/os_service.go
Os Service module for agent layer

## Tags
agent, health-checks, monitoring

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/checks/os_service.go> a code:Module ;
    code:name "agent/checks/os_service.go" ;
    code:description "Os Service module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:tags "agent", "health-checks", "monitoring" .
<!-- End LinkedDoc RDF -->
*/
package checks

import (
	"errors"
)

const (
	errOSServiceStatusCritical = "OS Service unhealthy"
)

var (
	ErrOSServiceStatusCritical = errors.New(errOSServiceStatusCritical)
)
