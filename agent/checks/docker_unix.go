// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

//go:build !windows

/*
# Module: agent/checks/docker_unix.go
Docker Unix module for agent layer

## Tags
agent, health-checks, monitoring

## Exports
DefaultDockerHost

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/checks/docker_unix.go> a code:Module ;
    code:name "agent/checks/docker_unix.go" ;
    code:description "Docker Unix module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:exports :DefaultDockerHost ;
    code:tags "agent", "health-checks", "monitoring" .
<!-- End LinkedDoc RDF -->
*/
package checks

const DefaultDockerHost = "unix:///var/run/docker.sock"
