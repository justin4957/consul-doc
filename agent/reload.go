// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: agent/reload.go
Reload module for agent layer

## Linked Modules
- [agent/config](../agent/config)

## Tags
agent

## Exports
ConfigReloader

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<agent/reload.go> a code:Module ;
    code:name "agent/reload.go" ;
    code:description "Reload module for agent layer" ;
    code:language "go" ;
    code:layer "agent" ;
    code:linksTo [
        code:name "agent/config" ;
        code:path "../agent/config"
    ] ;
    code:exports :ConfigReloader ;
    code:tags "agent" .
<!-- End LinkedDoc RDF -->
*/
package agent

import "github.com/hashicorp/consul/agent/config"

// ConfigReloader is a function type which may be implemented to support reloading
// of configuration.
type ConfigReloader func(rtConfig *config.RuntimeConfig) error
