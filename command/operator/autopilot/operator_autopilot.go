// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/operator/autopilot/operator_autopilot.go
Operator Autopilot module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/operator/autopilot/operator_autopilot.go> a code:Module ;
    code:name "command/operator/autopilot/operator_autopilot.go" ;
    code:description "Operator Autopilot module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package autopilot

import (
	"github.com/hashicorp/consul/command/flags"
	"github.com/mitchellh/cli"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return flags.Usage(help, nil)
}

const synopsis = "Provides tools for modifying Autopilot configuration"
const help = `
Usage: consul operator autopilot <subcommand> [options]

  The Autopilot operator command is used to interact with Consul's Autopilot
  subsystem. The command can be used to view or modify the current configuration.
`
