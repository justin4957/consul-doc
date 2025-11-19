// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/connect/connect.go
Connect module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, mtls, service-mesh, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/connect/connect.go> a code:Module ;
    code:name "command/connect/connect.go" ;
    code:description "Connect module for cli layer" ;
    code:language "go" ;
    code:layer "cli" ;
    code:linksTo [
        code:name "command/flags" ;
        code:path "../command/flags"
    ] ;
    code:exports :New ;
    code:tags "cli", "mtls", "service-mesh", "user-interface" .
<!-- End LinkedDoc RDF -->
*/
package connect

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

const synopsis = "Interact with Consul Connect"
const help = `
Usage: consul connect <subcommand> [options] [args]

  This command has subcommands for interacting with Consul Connect.

  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  Run the built-in Connect mTLS proxy

      $ consul connect proxy

  For more examples, ask for subcommand help or view the documentation.
`
