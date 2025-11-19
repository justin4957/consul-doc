// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/connect/ca/ca.go
Ca module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, mtls, service-mesh, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/connect/ca/ca.go> a code:Module ;
    code:name "command/connect/ca/ca.go" ;
    code:description "Ca module for cli layer" ;
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
package ca

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

const synopsis = "Interact with the Consul Connect Certificate Authority (CA)"
const help = `
Usage: consul connect ca <subcommand> [options] [args]

  This command has subcommands for interacting with Consul Connect's 
  Certificate Authority (CA).

  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  Get the configuration:

      $ consul connect ca get-config

  Update the configuration:

      $ consul connect ca set-config -config-file ca.json

  For more examples, ask for subcommand help or view the documentation.
`
