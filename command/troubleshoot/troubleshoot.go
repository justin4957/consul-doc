// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

/*
# Module: command/troubleshoot/troubleshoot.go
Troubleshoot module for cli layer

## Linked Modules
- [command/flags](../command/flags)

## Tags
cli, user-interface

## Exports
New

<!-- LinkedDoc RDF -->
@prefix code: <https://schema.codedoc.org/> .
<command/troubleshoot/troubleshoot.go> a code:Module ;
    code:name "command/troubleshoot/troubleshoot.go" ;
    code:description "Troubleshoot module for cli layer" ;
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
package troubleshoot

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

const synopsis = `CLI tools for troubleshooting Consul service mesh`
const help = `
Usage: consul troubleshoot <subcommand> [options]

  This command has subcommands for troubleshooting the service mesh.

  Here are some simple examples, and more detailed examples are available
  in the subcommands or the documentation.

  Troubleshoot Get Upstreams

    $ consul troubleshoot upstreams

  Troubleshoot Proxy

    $ consul troubleshoot proxy -upstream [options]

  For more examples, ask for subcommand help or view the documentation.
`
