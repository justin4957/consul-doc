// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/operator/usage/usage.go
Usage module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/operator/usage/usage.go> a code:Module ;
    code:name "command/operator/usage/usage.go" ;
    code:description "Usage module for cli layer" ;
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
package usage

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

const synopsis = "Provides cluster-level usage information"
const help = `
Usage: consul operator usage <subcommand> [options] [args]

  This command has subcommands for displaying usage information. The subcommands
  default to working with services registered with the local datacenter.

  For more examples, ask for subcommand help or view the documentation.
`
